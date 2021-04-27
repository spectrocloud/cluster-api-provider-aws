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

package scope

import (
	"context"
	"fmt"

	awsclient "github.com/aws/aws-sdk-go/aws/client"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2/klogr"
	"sigs.k8s.io/cluster-api-provider-aws/pkg/cloud/awserrors"
	"sigs.k8s.io/cluster-api-provider-aws/pkg/cloud/throttle"

	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
	clusterv1exp "sigs.k8s.io/cluster-api/exp/api/v1alpha3"
	"sigs.k8s.io/cluster-api/util/conditions"
	"sigs.k8s.io/cluster-api/util/patch"
	"sigs.k8s.io/controller-runtime/pkg/client"

	infrav1 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha3"
	controlplanev1exp "sigs.k8s.io/cluster-api-provider-aws/controlplane/eks/api/v1alpha3"
	infrav1exp "sigs.k8s.io/cluster-api-provider-aws/exp/api/v1alpha3"
	"sigs.k8s.io/cluster-api-provider-aws/pkg/cloud"
)

// ManagedMachinePoolScopeParams defines the input parameters used to create a new Scope.
type ManagedMachinePoolScopeParams struct {
	Client             client.Client
	Logger             logr.Logger
	Cluster            *clusterv1.Cluster
	ControlPlane       *controlplanev1exp.AWSManagedControlPlane
	ManagedMachinePool *infrav1exp.AWSManagedMachinePool
	MachinePool        *clusterv1exp.MachinePool
	ControllerName     string
	Endpoints          []ServiceEndpoint
	Session            awsclient.ConfigProvider

	EnableIAM bool
}

// NewManagedMachinePoolScope creates a new Scope from the supplied parameters.
// This is meant to be called for each reconcile iteration.
func NewManagedMachinePoolScope(params ManagedMachinePoolScopeParams) (*ManagedMachinePoolScope, error) {
	if params.ControlPlane == nil {
		return nil, errors.New("failed to generate new scope from nil AWSManagedMachinePool")
	}
	if params.MachinePool == nil {
		return nil, errors.New("failed to generate new scope from nil MachinePool")
	}
	if params.Logger == nil {
		params.Logger = klogr.New()
	}

	managedScope := &ManagedControlPlaneScope{
		Logger:         params.Logger,
		Client:         params.Client,
		Cluster:        params.Cluster,
		ControlPlane:   params.ControlPlane,
		controllerName: params.ControllerName,
	}
	session, serviceLimiters, err := sessionForClusterWithRegion(params.Client, managedScope, params.ControlPlane.Spec.Region, params.Endpoints, params.Logger)
	if err != nil {
		return nil, errors.Errorf("failed to create aws session: %v", err)
	}

	helper, err := patch.NewHelper(params.ManagedMachinePool, params.Client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init patch helper")
	}

	return &ManagedMachinePoolScope{
		Logger:             params.Logger,
		client:             params.Client,
		Cluster:            params.Cluster,
		ControlPlane:       params.ControlPlane,
		ManagedMachinePool: params.ManagedMachinePool,
		MachinePool:        params.MachinePool,
		patchHelper:        helper,
		session:            session,
		serviceLimiters:    serviceLimiters,
		controllerName:     params.ControllerName,
		enableIAM:          params.EnableIAM,
	}, nil
}

// ManagedMachinePoolScope defines the basic context for an actuator to operate upon.
type ManagedMachinePoolScope struct {
	logr.Logger
	client      client.Client
	patchHelper *patch.Helper

	Cluster            *clusterv1.Cluster
	ControlPlane       *controlplanev1exp.AWSManagedControlPlane
	ManagedMachinePool *infrav1exp.AWSManagedMachinePool
	MachinePool        *clusterv1exp.MachinePool

	session         awsclient.ConfigProvider
	serviceLimiters throttle.ServiceLimiters
	controllerName  string

	enableIAM bool
}

// Name returns the managed machine pool name.
func (s *ManagedMachinePoolScope) Name() string {
	return s.ManagedMachinePool.Name
}

// Name returns the managed machine pool namespace.
func (s *ManagedMachinePoolScope) Namespace() string {
	return s.ManagedMachinePool.Namespace
}

// ManagedPoolName returns the managed machine pool name.
func (s *ManagedMachinePoolScope) ManagedPoolName() string {
	return s.ManagedMachinePool.Name
}

// ServiceLimiter returns the AWS SDK session. Used for creating clients
func (s *ManagedMachinePoolScope) ServiceLimiter(service string) *throttle.ServiceLimiter {
	if sl, ok := s.serviceLimiters[service]; ok {
		return sl
	}
	return nil
}

// ClusterName returns the cluster name.
func (s *ManagedMachinePoolScope) ClusterName() string {
	return s.Cluster.Name
}

// EnableIAM indicates that reconciliation should create IAM roles
func (s *ManagedMachinePoolScope) EnableIAM() bool {
	return s.enableIAM
}

// IdentityRef returns the cluster identityRef.
func (s *ManagedMachinePoolScope) IdentityRef() *infrav1.AWSIdentityReference {
	return s.ControlPlane.Spec.IdentityRef
}

func (m *ManagedMachinePoolScope) CoreSecurityGroups(scope EC2Scope) ([]string, error) {
	// These are common across both controlplane and node machines
	sgRoles := []infrav1.SecurityGroupRole{
		controlplanev1exp.SecurityGroupCluster,
		infrav1.SecurityGroupEKSNodeAdditional,
	}

	ids := make([]string, 0, len(sgRoles))
	for _, sg := range sgRoles {
		if _, ok := scope.SecurityGroups()[sg]; !ok {
			return nil, awserrors.NewFailedDependency(
				fmt.Sprintf("%s security group not available", sg),
			)
		}
		ids = append(ids, scope.SecurityGroups()[sg].ID)
	}
	return ids, nil
}

func (m *ManagedMachinePoolScope) LaunchTemplateID() string {
	if lt := m.ManagedMachinePool.Spec.LaunchTemplate; lt != nil {
		return lt.ID
	}
	return ""
}

func (m *ManagedMachinePoolScope) OwnerObject() conditions.Setter {
	return m.ManagedMachinePool
}

// AdditionalTags returns AdditionalTags from the scope's ManagedMachinePool
// The returned value will never be nil.
func (s *ManagedMachinePoolScope) AdditionalTags() infrav1.Tags {
	if s.ManagedMachinePool.Spec.AdditionalTags == nil {
		s.ManagedMachinePool.Spec.AdditionalTags = infrav1.Tags{}
	}
	//name, _ := eks.GenerateEKSName(s.Cluster.Name, s.Cluster.GetNamespace())
	s.ManagedMachinePool.Spec.AdditionalTags["eks:cluster-name"] = s.ClusterName()
	s.ManagedMachinePool.Spec.AdditionalTags["eks:nodegroup-name"] = s.Name()
	return s.ManagedMachinePool.Spec.AdditionalTags.DeepCopy()
}

// RoleName returns the node group role name
func (s *ManagedMachinePoolScope) RoleName() string {
	return s.ManagedMachinePool.Spec.RoleName
}

// Version returns the nodegroup Kubernetes version
func (s *ManagedMachinePoolScope) Version() *string {
	return s.MachinePool.Spec.Template.Spec.Version
}

// ControlPlaneSubnets returns the control plane subnets.
func (s *ManagedMachinePoolScope) ControlPlaneSubnets() infrav1.Subnets {
	return s.ControlPlane.Spec.NetworkSpec.Subnets
}

// SubnetIDs returns the machine pool subnet IDs.
func (s *ManagedMachinePoolScope) SubnetIDs() ([]string, error) {
	strategy, err := newDefaultSubnetPlacementStrategy(s.Logger)
	if err != nil {
		return []string{}, fmt.Errorf("getting subnet placement strategy: %w", err)
	}

	return strategy.Place(&placementInput{
		SpecSubnetIDs:           s.ManagedMachinePool.Spec.SubnetIDs,
		SpecAvailabilityZones:   s.ManagedMachinePool.Spec.AvailabilityZones,
		ParentAvailabilityZones: s.MachinePool.Spec.FailureDomains,
		ControlplaneSubnets:     s.ControlPlaneSubnets(),
	})
}

// NodegroupReadyFalse marks the ready condition false using warning if error isn't
// empty
func (s *ManagedMachinePoolScope) NodegroupReadyFalse(reason string, err string) error {
	severity := clusterv1.ConditionSeverityWarning
	if err == "" {
		severity = clusterv1.ConditionSeverityInfo
	}
	conditions.MarkFalse(
		s.ManagedMachinePool,
		infrav1exp.EKSNodegroupReadyCondition,
		reason,
		severity,
		err,
	)
	if err := s.PatchObject(); err != nil {
		return errors.Wrap(err, "failed to mark nodegroup not ready")
	}
	return nil
}

// IAMReadyFalse marks the ready condition false using warning if error isn't
// empty
func (s *ManagedMachinePoolScope) IAMReadyFalse(reason string, err string) error {
	severity := clusterv1.ConditionSeverityWarning
	if err == "" {
		severity = clusterv1.ConditionSeverityInfo
	}
	conditions.MarkFalse(
		s.ManagedMachinePool,
		infrav1exp.IAMNodegroupRolesReadyCondition,
		reason,
		severity,
		err,
	)
	if err := s.PatchObject(); err != nil {
		return errors.Wrap(err, "failed to mark nodegroup role not ready")
	}
	return nil
}

// PatchObject persists the control plane configuration and status.
func (s *ManagedMachinePoolScope) PatchObject() error {
	return s.patchHelper.Patch(
		context.TODO(),
		s.ManagedMachinePool,
		patch.WithOwnedConditions{Conditions: []clusterv1.ConditionType{
			infrav1exp.EKSNodegroupReadyCondition,
			infrav1exp.IAMNodegroupRolesReadyCondition,
		}})
}

// Close closes the current scope persisting the control plane configuration and status.
func (s *ManagedMachinePoolScope) Close() error {
	return s.PatchObject()
}

// InfraCluster returns the AWS infrastructure cluster or control plane object.
func (s *ManagedMachinePoolScope) InfraCluster() cloud.ClusterObject {
	return s.ControlPlane
}

// Session returns the AWS SDK session. Used for creating clients
func (s *ManagedMachinePoolScope) Session() awsclient.ConfigProvider {
	return s.session
}

// ControllerName returns the name of the controller that
// created the ManagedMachinePool.
func (s *ManagedMachinePoolScope) ControllerName() string {
	return s.controllerName
}

// KubernetesClusterName is the name of the EKS cluster name.
func (s *ManagedMachinePoolScope) KubernetesClusterName() string {
	return s.ControlPlane.Spec.EKSClusterName
}

// NodegroupName is the name of the EKS nodegroup
func (s *ManagedMachinePoolScope) NodegroupName() string {
	return s.ManagedMachinePool.Spec.EKSNodegroupName
}

// GetRawBootstrapData returns the bootstrap data from the secret in the Machine's bootstrap.dataSecretName.
// TODO: stolen from AWSMachinePool, it's not clear if this will remain identical to
// that code
func (m *ManagedMachinePoolScope) GetRawBootstrapData() ([]byte, error) {
	if m.MachinePool.Spec.Template.Spec.Bootstrap.DataSecretName == nil {
		return nil, errors.New("error retrieving bootstrap data: linked Machine's bootstrap.dataSecretName is nil")
	}

	secret := &corev1.Secret{}
	key := types.NamespacedName{Namespace: m.Namespace(), Name: *m.MachinePool.Spec.Template.Spec.Bootstrap.DataSecretName}

	if err := m.client.Get(context.TODO(), key, secret); err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve bootstrap data secret for AWSMachine %s/%s", m.Namespace(), m.Name())
	}

	value, ok := secret.Data["value"]
	if !ok {
		return nil, errors.New("error retrieving bootstrap data: secret value key is missing")
	}

	return value, nil
}
