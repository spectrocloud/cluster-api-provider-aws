package eks

import (
	"context"
	"testing"

	ekstypes "github.com/aws/aws-sdk-go-v2/service/eks/types"
	. "github.com/onsi/gomega"
	capierrors "sigs.k8s.io/cluster-api/errors"

	expinfrav1 "sigs.k8s.io/cluster-api-provider-aws/v2/exp/api/v1beta2"
	"sigs.k8s.io/cluster-api-provider-aws/v2/pkg/cloud/scope"
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
