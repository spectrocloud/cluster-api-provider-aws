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
	"github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/eks/eksiface"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"

	"sigs.k8s.io/cluster-api-provider-aws/pkg/cloud/scope"
	"sigs.k8s.io/cluster-api-provider-aws/pkg/cloud/services/eks/iam"
)

// Service holds a collection of interfaces.
// The interfaces are broken down like this to group functions together.
// One alternative is to have a large list of functions from the ec2 client.
type Service struct {
	scope     *scope.ManagedControlPlaneScope
	EC2Client ec2iface.EC2API
	EKSClient eksiface.EKSAPI
	iam.IAMService
	STSClient stsiface.STSAPI
}

// NewService returns a new service given the api clients.
func NewService(controlPlaneScope *scope.ManagedControlPlaneScope) *Service {
	return &Service{
		scope:     controlPlaneScope,
		EC2Client: scope.NewEC2Client(controlPlaneScope, controlPlaneScope, controlPlaneScope.ControlPlane),
		EKSClient: scope.NewEKSClient(controlPlaneScope, controlPlaneScope, controlPlaneScope.ControlPlane),
		IAMService: iam.IAMService{
			Logger:    controlPlaneScope.Logger,
			IAMClient: scope.NewIAMClient(controlPlaneScope, controlPlaneScope, controlPlaneScope.ControlPlane),
		},
		STSClient: scope.NewSTSClient(controlPlaneScope, controlPlaneScope, controlPlaneScope.ControlPlane),
	}
}

// NodegroupService holds a collection of interfaces.
// The interfaces are broken down like this to group functions together.
// One alternative is to have a large list of functions from the ec2 client.
type NodegroupService struct {
	scope             *scope.ManagedMachinePoolScope
	AutoscalingClient autoscalingiface.AutoScalingAPI
	EKSClient         eksiface.EKSAPI
	iam.IAMService
	STSClient stsiface.STSAPI
}

// NewNodegroupService returns a new service given the api clients.
func NewNodegroupService(machinePoolScope *scope.ManagedMachinePoolScope) *NodegroupService {
	return &NodegroupService{
		scope:             machinePoolScope,
		AutoscalingClient: scope.NewASGClient(machinePoolScope, machinePoolScope, machinePoolScope.ManagedMachinePool),
		EKSClient:         scope.NewEKSClient(machinePoolScope, machinePoolScope, machinePoolScope.ManagedMachinePool),
		IAMService: iam.IAMService{
			Logger:    machinePoolScope.Logger,
			IAMClient: scope.NewIAMClient(machinePoolScope, machinePoolScope, machinePoolScope.ManagedMachinePool),
		},
		STSClient: scope.NewSTSClient(machinePoolScope, machinePoolScope, machinePoolScope.ManagedMachinePool),
	}
}
