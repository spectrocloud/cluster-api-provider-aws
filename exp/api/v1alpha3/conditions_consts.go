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

package v1alpha3

import clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"

const (
	// ASGReadyCondition reports on current status of the autoscaling group. Ready indicates the group is provisioned
	ASGReadyCondition clusterv1.ConditionType = "ASGReady"
	// ASGNotFoundReason used when the autoscaling group couldn't be retrieved.
	ASGNotFoundReason = "ASGNotFound"
	// ASGProvisionFailedReason used for failures during autoscaling group provisioning.
	ASGProvisionFailedReason = "ASGProvisionFailed"
	// ASGDeletionInProgress ASG is in a deletion in progress state.
	ASGDeletionInProgress = "ASGDeletionInProgress"

	// LaunchTemplateReadyCondition represents the status of an AWSMachinePool's associated Launch Template
	LaunchTemplateReadyCondition clusterv1.ConditionType = "LaunchTemplateReady"
	// LaunchTemplateNotFoundReason is used when an associated Launch Template can't be found
	LaunchTemplateNotFoundReason = "LaunchTemplateNotFound"
	// LaunchTemplateCreateFailedReason used for failures during Launch Template creation
	LaunchTemplateCreateFailedReason = "LaunchTemplateCreateFailed"

	// InstanceRefreshStartedCondition reports on successfully starting instance refresh.
	InstanceRefreshStartedCondition clusterv1.ConditionType = "InstanceRefreshStarted"
	// InstanceRefreshNotReadyReason used to report instance refresh is not initiated.
	// If there are instance refreshes that are in progress, then a new instance refresh request will fail.
	InstanceRefreshNotReadyReason = "InstanceRefreshNotReady"
	// InstanceRefreshFailedReason used to report when there instance refresh is not initiated.
	InstanceRefreshFailedReason = "InstanceRefreshFailed"
)

const (
	// EKSNodegroupReadyCondition condition reports on the successful reconciliation of eks control plane.
	EKSNodegroupReadyCondition clusterv1.ConditionType = "EKSNodegroupReady"
	// EKSNodegroupReconciliationFailedReason used to report failures while reconciling EKS control plane
	EKSNodegroupReconciliationFailedReason = "EKSNodegroupReconciliationFailed"
	// WaitingForEKSControlPlaneReason used when the machine pool is waiting for
	// EKS control plane infrastructure to be ready before proceeding.
	WaitingForEKSControlPlaneReason = "WaitingForEKSControlPlane"
)

const (
	// IAMNodegroupRolesReadyCondition condition reports on the successful
	// reconciliation of EKS nodegroup iam roles.
	IAMNodegroupRolesReadyCondition clusterv1.ConditionType = "IAMNodegroupRolesReady"
	// IAMNodegroupRolesReconciliationFailedReason used to report failures while
	// reconciling EKS nodegroup iam roles
	IAMNodegroupRolesReconciliationFailedReason = "IAMNodegroupRolesReconciliationFailed"
)
