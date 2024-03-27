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
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/service/eks"
	awsiam "github.com/aws/aws-sdk-go/service/iam"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"sigs.k8s.io/cluster-api-provider-aws/cmd/clusterawsadm/converters"
	iamv1 "sigs.k8s.io/cluster-api-provider-aws/iam/api/v1beta1"
	tagConverter "sigs.k8s.io/cluster-api-provider-aws/pkg/cloud/converters"
	"sigs.k8s.io/cluster-api-provider-aws/pkg/cloud/services/iam"
	"sigs.k8s.io/cluster-api/controllers/remote"
)

func (s *Service) reconcileOIDCProvider(cluster *eks.Cluster) error {
	if !s.scope.ControlPlane.Spec.AssociateOIDCProvider {
		return nil
	}

	if !s.scope.EnableIAM() {
		return errors.New("'AssociateOIDCProvider' provided without enabling the 'EKSEnableIAM' feature flag")
	}

	if s.scope.ControlPlane.Status.OIDCProvider.ARN == "" {
		s.scope.Info("Reconciling EKS OIDC Provider", "cluster-name", cluster.Name)
		oidcProvider, err := s.CreateOIDCProvider(cluster)
		if err != nil {
			return errors.Wrap(err, "failed to create OIDC provider")
		}
		s.scope.ControlPlane.Status.OIDCProvider.ARN = oidcProvider
		anno := s.scope.ControlPlane.GetAnnotations()
		if anno == nil {
			anno = make(map[string]string)
		}
		anno["aws.spectrocloud.com/oidcProviderArn"] = oidcProvider
		s.scope.ControlPlane.SetAnnotations(anno)
		if err := s.scope.PatchObject(); err != nil {
			return errors.Wrap(err, "failed to update control plane with OIDC provider ARN")
		}
	}

	// tagging the OIDC provider with the same tags of cluster
	inputForTags := awsiam.TagOpenIDConnectProviderInput{
		OpenIDConnectProviderArn: &s.scope.ControlPlane.Status.OIDCProvider.ARN,
		Tags:                     tagConverter.MapToIAMTags(tagConverter.MapPtrToMap(cluster.Tags)),
	}
	if _, err := s.IAMClient.TagOpenIDConnectProvider(&inputForTags); err != nil {
		return errors.Wrap(err, "failed to tag OIDC provider")
	}

	if s.scope.ControlPlane.Status.OIDCProvider.TrustPolicy == "" {
		policy, err := converters.IAMPolicyDocumentToJSON(s.buildOIDCTrustPolicy())
		if err != nil {
			return errors.Wrap(err, "failed to parse IAM policy")
		}
		if err := s.reconcileTrustPolicy(); err != nil {
			return errors.Wrap(err, "failed to reconcile trust policy in workload cluster")
		}
		s.scope.ControlPlane.Status.OIDCProvider.TrustPolicy = iam.WhitespaceRegex.ReplaceAllString(policy, "")
		if err := s.scope.PatchObject(); err != nil {
			return errors.Wrap(err, "failed to update control plane with OIDC provider trustPolicy")
		}
	}

	return nil
}

func (s *Service) reconcileTrustPolicy() error {
	ctx := context.Background()

	clusterKey := client.ObjectKey{
		Name:      s.scope.Name(),
		Namespace: s.scope.Namespace(),
	}

	restConfig, err := remote.RESTConfig(ctx, s.scope.ControlPlane.Name, s.scope.Client, clusterKey)
	if err != nil {
		return fmt.Errorf("getting remote client for %s/%s: %w", s.scope.Namespace(), s.scope.Name(), err)
	}

	remoteClient, err := client.New(restConfig, client.Options{})
	if err != nil {
		return fmt.Errorf("getting client for remote cluster: %w", err)
	}

	configMapRef := types.NamespacedName{
		Name:      iam.TrustPolicyConfigMapName,
		Namespace: iam.TrustPolicyConfigMapNamespace,
	}

	trustPolicyConfigMap := &corev1.ConfigMap{}

	err = remoteClient.Get(ctx, configMapRef, trustPolicyConfigMap)
	if err != nil && !apierrors.IsNotFound(err) {
		return fmt.Errorf("getting %s/%s config map: %w", iam.TrustPolicyConfigMapNamespace, iam.TrustPolicyConfigMapName, err)
	}

	policy, err := converters.IAMPolicyDocumentToJSON(s.buildOIDCTrustPolicy())
	if err != nil {
		return errors.Wrap(err, "failed to parse IAM policy")
	}

	trustPolicyConfigMap.Data = map[string]string{
		"trust-policy.json": policy,
	}

	if trustPolicyConfigMap.UID == "" {
		trustPolicyConfigMap.Name = iam.TrustPolicyConfigMapName
		trustPolicyConfigMap.Namespace = iam.TrustPolicyConfigMapNamespace
		s.V(2).Info("Creating new Trust Policy ConfigMap", "cluster", s.scope.Name(), "configmap", iam.TrustPolicyConfigMapName)
		return remoteClient.Create(ctx, trustPolicyConfigMap)
	}

	s.V(2).Info("Updating existing Trust Policy ConfigMap", "cluster", s.scope.Name(), "configmap", iam.TrustPolicyConfigMapName)
	return remoteClient.Update(ctx, trustPolicyConfigMap)
}

func (s *Service) deleteOIDCProvider() error {
	anno := s.scope.ControlPlane.GetAnnotations()
	arn := anno["aws.spectrocloud.com/oidcProviderArn"]

	if arn == "" {
		arn = s.scope.ControlPlane.Status.OIDCProvider.ARN
	}

	if !s.scope.ControlPlane.Spec.AssociateOIDCProvider || arn == "" {
		return nil
	}

	providerARN := arn
	if err := s.DeleteOIDCProvider(&providerARN); err != nil {
		return errors.Wrap(err, "failed to delete OIDC provider")
	}

	s.scope.ControlPlane.Status.OIDCProvider.ARN = ""
	if err := s.scope.PatchObject(); err != nil {
		return errors.Wrap(err, "failed to update control plane with OIDC provider ARN")
	}

	anno["aws.spectrocloud.com/oidcProviderArn"] = ""
	s.scope.ControlPlane.SetAnnotations(anno)

	return nil
}

func (s *Service) buildOIDCTrustPolicy() iamv1.PolicyDocument {
	providerARN := s.scope.ControlPlane.Status.OIDCProvider.ARN
	conditionValue := providerARN[strings.Index(providerARN, "/")+1:] + ":sub"

	return iamv1.PolicyDocument{
		Version: "2012-10-17",
		Statement: iamv1.Statements{
			iamv1.StatementEntry{
				Sid:    "",
				Effect: "Allow",
				Principal: iamv1.Principals{
					iamv1.PrincipalFederated: iamv1.PrincipalID{providerARN},
				},
				Action: iamv1.Actions{"sts:AssumeRoleWithWebIdentity"},
				Condition: iamv1.Conditions{
					"ForAnyValue:StringLike": map[string][]string{
						conditionValue: {"system:serviceaccount:${SERVICE_ACCOUNT_NAMESPACE}:${SERVICE_ACCOUNT_NAME}"},
					},
				},
			},
		},
	}
}
