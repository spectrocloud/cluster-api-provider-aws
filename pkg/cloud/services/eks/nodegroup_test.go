/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package eks

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	ekstypes "github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/go-logr/logr"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	capierrors "sigs.k8s.io/cluster-api/errors"
	clusterv1 "sigs.k8s.io/cluster-api/api/core/v1beta2"

	ekscontrolplanev1 "sigs.k8s.io/cluster-api-provider-aws/v2/controlplane/eks/api/v1beta2"
	expinfrav1 "sigs.k8s.io/cluster-api-provider-aws/v2/exp/api/v1beta2"
	"sigs.k8s.io/cluster-api-provider-aws/v2/pkg/cloud/scope"
	"sigs.k8s.io/cluster-api-provider-aws/v2/pkg/cloud/services/eks/iam"
	"sigs.k8s.io/cluster-api-provider-aws/v2/pkg/cloud/services/eks/mock_eksiface"
	"sigs.k8s.io/cluster-api-provider-aws/v2/pkg/logger"
)

func TestSetStatus(t *testing.T) {
	g := NewWithT(t)
	message := "VcpuLimitExceeded"

	s := &NodegroupService{
		scope: &scope.ManagedMachinePoolScope{
			ManagedMachinePool: &expinfrav1.AWSManagedMachinePool{
				Status: expinfrav1.AWSManagedMachinePoolStatus{
					Ready: false,
				},
			},
		},
	}

	issue := ekstypes.Issue{
		Code:        ekstypes.NodegroupIssueCodeAsgInstanceLaunchFailures,
		Message:     &message,
		ResourceIds: []string{"my-worker-nodes"},
	}
	ng := &ekstypes.Nodegroup{
		Status: ekstypes.NodegroupStatusDegraded,
		Health: &ekstypes.NodegroupHealth{
			Issues: []ekstypes.Issue{issue},
		},
	}

	err := s.setStatus(context.TODO(), ng)
	g.Expect(err).ToNot(BeNil())
	// ensure machine pool status values are set as expected
	g.Expect(*s.scope.ManagedMachinePool.Status.FailureMessage).To(ContainSubstring(message))
	g.Expect(s.scope.ManagedMachinePool.Status.Ready).To(Equal(false))
	g.Expect(*s.scope.ManagedMachinePool.Status.FailureReason).To(Equal(string(capierrors.InsufficientResourcesMachineError)))
}

func TestReconcileNodegroupConfig_NodeRepairConfig(t *testing.T) {
	tests := []struct {
		name        string
		specRepair  *expinfrav1.NodeRepairConfig
		liveRepair  *ekstypes.NodeRepairConfig
		expectCall  bool
		expectValue *ekstypes.NodeRepairConfig
	}{
		{
			name:       "nil spec, nil live: no call (adoption / never-set nodegroup)",
			specRepair: nil,
			liveRepair: nil,
			expectCall: false,
		},
		{
			name:       "nil spec, live enabled: no call (does not overwrite unmanaged field)",
			specRepair: nil,
			liveRepair: &ekstypes.NodeRepairConfig{Enabled: aws.Bool(true)},
			expectCall: false,
		},
		{
			name:        "explicit spec differs from live: call fires with spec value",
			specRepair:  &expinfrav1.NodeRepairConfig{Enabled: aws.Bool(true)},
			liveRepair:  &ekstypes.NodeRepairConfig{Enabled: aws.Bool(false)},
			expectCall:  true,
			expectValue: &ekstypes.NodeRepairConfig{Enabled: aws.Bool(true)},
		},
		{
			name:       "explicit spec matches live: no call",
			specRepair: &expinfrav1.NodeRepairConfig{Enabled: aws.Bool(false)},
			liveRepair: &ekstypes.NodeRepairConfig{Enabled: aws.Bool(false)},
			expectCall: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			mockControl := gomock.NewController(t)
			defer mockControl.Finish()
			eksMock := mock_eksiface.NewMockEKSAPI(mockControl)

			if tc.expectCall {
				expected := tc.expectValue
				eksMock.EXPECT().
					UpdateNodegroupConfig(gomock.Eq(context.TODO()), gomock.AssignableToTypeOf(&eks.UpdateNodegroupConfigInput{})).
					DoAndReturn(func(_ context.Context, in *eks.UpdateNodegroupConfigInput, _ ...func(*eks.Options)) (*eks.UpdateNodegroupConfigOutput, error) {
						g.Expect(in.NodeRepairConfig).To(Equal(expected))
						return &eks.UpdateNodegroupConfigOutput{}, nil
					})
			}

			log := logger.NewLogger(logr.Discard())
			s := &NodegroupService{
				scope: &scope.ManagedMachinePoolScope{
					Logger: *log,
					ControlPlane: &ekscontrolplanev1.AWSManagedControlPlane{
						Spec: ekscontrolplanev1.AWSManagedControlPlaneSpec{
							EKSClusterName: "test-cluster",
						},
					},
					ManagedMachinePool: &expinfrav1.AWSManagedMachinePool{
						Spec: expinfrav1.AWSManagedMachinePoolSpec{
							EKSNodegroupName: "test-ng",
							NodeRepairConfig: tc.specRepair,
						},
					},
					MachinePool: &clusterv1.MachinePool{},
				},
				EKSClient:  eksMock,
				IAMService: iam.IAMService{Wrapper: log},
			}

			ng := &ekstypes.Nodegroup{
				NodegroupName: aws.String("test-ng"),
				ScalingConfig: &ekstypes.NodegroupScalingConfig{
					DesiredSize: aws.Int32(1),
					MinSize:     aws.Int32(1),
					MaxSize:     aws.Int32(1),
				},
				NodeRepairConfig: tc.liveRepair,
			}

			err := s.reconcileNodegroupConfig(context.TODO(), ng)
			g.Expect(err).NotTo(HaveOccurred())
		})
	}
}

func TestIsSymbolicLaunchTemplateVersion(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    bool
	}{
		{name: "$Latest is symbolic", version: "$Latest", want: true},
		{name: "$Default is symbolic", version: "$Default", want: true},
		{name: "concrete version 1", version: "1", want: false},
		{name: "concrete version 42", version: "42", want: false},
		{name: "empty string", version: "", want: false},
		{name: "lowercase $latest is not symbolic", version: "$latest", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymbolicLaunchTemplateVersion(tt.version); got != tt.want {
				t.Errorf("isSymbolicLaunchTemplateVersion(%q) = %v, want %v", tt.version, got, tt.want)
			}
		})
	}
}
