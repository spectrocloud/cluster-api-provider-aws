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
	"encoding/base64"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"

	ekscontrolplanev1 "sigs.k8s.io/cluster-api-provider-aws/v2/controlplane/eks/api/v1beta2"
	"sigs.k8s.io/cluster-api-provider-aws/v2/pkg/record"
	"sigs.k8s.io/cluster-api/util/kubeconfig"
	"sigs.k8s.io/cluster-api/util/secret"
)

const (
	tokenPrefix       = "k8s-aws-v1." //nolint:gosec
	clusterNameHeader = "x-k8s-aws-id"
	tokenAgeMins      = 15
)

func (s *Service) reconcileKubeconfig(ctx context.Context, cluster *eks.Cluster) error {
	s.scope.Debug("Reconciling EKS kubeconfigs for cluster", "cluster-name", s.scope.KubernetesClusterName())

	clusterRef := types.NamespacedName{
		Name:      s.scope.Cluster.Name,
		Namespace: s.scope.Cluster.Namespace,
	}

	// Create the kubeconfig used by CAPI
	configSecret, err := secret.GetFromNamespacedName(ctx, s.scope.Client, clusterRef, secret.Kubeconfig)
	if err != nil {
		if !apierrors.IsNotFound(err) {
			return errors.Wrap(err, "failed to get kubeconfig secret")
		}

		if createErr := s.createCAPIKubeconfigSecret(
			ctx,
			cluster,
			&clusterRef,
		); createErr != nil {
			return fmt.Errorf("creating kubeconfig secret: %w", err)
		}
	} else if updateErr := s.updateCAPIKubeconfigSecret(ctx, configSecret, cluster); updateErr != nil {
		return fmt.Errorf("updating kubeconfig secret: %w", err)
	}

	// Set initialized to true to indicate the kubconfig has been created
	s.scope.ControlPlane.Status.Initialized = true

	return nil
}

func (s *Service) reconcileAdditionalKubeconfigs(ctx context.Context, cluster *eks.Cluster) error {
	s.scope.Debug("Reconciling additional EKS kubeconfigs for cluster", "cluster-name", s.scope.KubernetesClusterName())

	clusterRef := types.NamespacedName{
		Name:      s.scope.Cluster.Name + "-user",
		Namespace: s.scope.Cluster.Namespace,
	}

	// Create the additional kubeconfig for users. This doesn't need updating on every sync
	configSecret, err := secret.GetFromNamespacedName(ctx, s.scope.Client, clusterRef, secret.Kubeconfig)
	if err != nil {
		if !apierrors.IsNotFound(err) {
			return errors.Wrap(err, "failed to get kubeconfig (user) secret")
		}

		createErr := s.createUserKubeconfigSecret(
			ctx,
			cluster,
			&clusterRef,
		)
		if createErr != nil {
			return err
		}
	} else if updateErr := s.updateUserKubeconfigSecret(ctx, configSecret, cluster); updateErr != nil {
		return fmt.Errorf("updating kubeconfig secret: %w", err)
	}

	return nil
}

func (s *Service) createCAPIKubeconfigSecret(ctx context.Context, cluster *eks.Cluster, clusterRef *types.NamespacedName) error {
	controllerOwnerRef := *metav1.NewControllerRef(s.scope.ControlPlane, ekscontrolplanev1.GroupVersion.WithKind("AWSManagedControlPlane"))

	clusterName := s.scope.KubernetesClusterName()
	userName := s.getKubeConfigUserName(clusterName, false)

	cfg, err := s.createBaseKubeConfig(cluster, userName)
	if err != nil {
		return fmt.Errorf("creating base kubeconfig: %w", err)
	}

	token, err := s.generateToken()
	if err != nil {
		return fmt.Errorf("generating presigned token: %w", err)
	}

	cfg.AuthInfos = map[string]*api.AuthInfo{
		userName: {
			Token: token,
		},
	}

	out, err := clientcmd.Write(*cfg)
	if err != nil {
		return errors.Wrap(err, "failed to serialize config to yaml")
	}

	kubeconfigSecret := kubeconfig.GenerateSecretWithOwner(*clusterRef, out, controllerOwnerRef)
	if err := s.scope.Client.Create(ctx, kubeconfigSecret); err != nil {
		return errors.Wrap(err, "failed to create kubeconfig secret")
	}

	record.Eventf(s.scope.ControlPlane, "SucessfulCreateKubeconfig", "Created kubeconfig for cluster %q", s.scope.Name())
	return nil
}

func (s *Service) updateCAPIKubeconfigSecret(ctx context.Context, configSecret *corev1.Secret, cluster *eks.Cluster) error {
	s.scope.Debug("Updating EKS kubeconfigs for cluster", "cluster-name", s.scope.KubernetesClusterName())

	data, ok := configSecret.Data[secret.KubeconfigDataName]
	if !ok {
		return errors.Errorf("missing key %q in secret data", secret.KubeconfigDataName)
	}

	config, err := clientcmd.Load(data)
	if err != nil {
		return errors.Wrap(err, "failed to convert kubeconfig Secret into a clientcmdapi.Config")
	}

	token, err := s.generateToken()
	if err != nil {
		return fmt.Errorf("generating presigned token: %w", err)
	}

	userName := s.getKubeConfigUserName(*cluster.Name, false)
	config.AuthInfos[userName].Token = token

	out, err := clientcmd.Write(*config)
	if err != nil {
		return errors.Wrap(err, "failed to serialize config to yaml")
	}

	configSecret.Data[secret.KubeconfigDataName] = out

	err = s.scope.Client.Update(ctx, configSecret)
	if err != nil {
		return fmt.Errorf("updating kubeconfig secret: %w", err)
	}

	return nil
}

func (s *Service) createUserKubeconfigSecret(ctx context.Context, cluster *eks.Cluster, clusterRef *types.NamespacedName) error {
	controllerOwnerRef := *metav1.NewControllerRef(s.scope.ControlPlane, ekscontrolplanev1.GroupVersion.WithKind("AWSManagedControlPlane"))

	clusterName := s.scope.KubernetesClusterName()
	userName := s.getKubeConfigUserName(clusterName, true)

	cfg, err := s.createBaseKubeConfig(cluster, userName)
	if err != nil {
		return fmt.Errorf("creating base kubeconfig: %w", err)
	}

	// Version v1alpha1 was removed in Kubernetes v1.23.
	// Version v1 was released in Kubernetes v1.23.
	// Version v1beta1 was selected as it has the widest range of support
	// This should be changed to v1 once EKS no longer supports Kubernetes <v1.23
	execConfig := &api.ExecConfig{APIVersion: "client.authentication.k8s.io/v1beta1"}
	switch s.scope.TokenMethod() {
	case ekscontrolplanev1.EKSTokenMethodIAMAuthenticator:
		execConfig.Command = "aws-iam-authenticator"
		execConfig.Args = []string{
			"token",
			"-i",
			clusterName,
		}
	case ekscontrolplanev1.EKSTokenMethodAWSCli:
		execConfig.Command = "aws"
		execConfig.Args = []string{
			"eks",
			"get-token",
			"--cluster-name",
			clusterName,
		}
	default:
		return fmt.Errorf("using token method %s: %w", s.scope.TokenMethod(), ErrUnknownTokenMethod)
	}
	cfg.AuthInfos = map[string]*api.AuthInfo{
		userName: {
			Exec: execConfig,
		},
	}

	out, err := clientcmd.Write(*cfg)
	if err != nil {
		return errors.Wrap(err, "failed to serialize config to yaml")
	}

	kubeconfigSecret := kubeconfig.GenerateSecretWithOwner(*clusterRef, out, controllerOwnerRef)
	if err := s.scope.Client.Create(ctx, kubeconfigSecret); err != nil {
		return errors.Wrap(err, "failed to create kubeconfig secret")
	}

	record.Eventf(s.scope.ControlPlane, "SucessfulCreateUserKubeconfig", "Created user kubeconfig for cluster %q", s.scope.Name())
	return nil
}

func (s *Service) updateUserKubeconfigSecret(ctx context.Context, configSecret *corev1.Secret, cluster *eks.Cluster) error {
	s.scope.V(2).Info("Updating EKS kubeconfigs for cluster", "cluster-name", s.scope.KubernetesClusterName())

	data, ok := configSecret.Data[secret.KubeconfigDataName]
	if !ok {
		return errors.Errorf("missing key %q in secret data", secret.KubeconfigDataName)
	}

	config, err := clientcmd.Load(data)
	if err != nil {
		return errors.Wrap(err, "failed to convert kubeconfig Secret into a clientcmdapi.Config")
	}

	userName := s.getKubeConfigUserName(*cluster.Name, false)

	clusterName := s.scope.KubernetesClusterName()

	if config.AuthInfos[userName].Exec != nil && config.AuthInfos[userName].Exec.APIVersion == "client.authentication.k8s.io/v1beta1" {
		return nil
	}

	execConfig := &api.ExecConfig{APIVersion: "client.authentication.k8s.io/v1beta1"}
	switch s.scope.TokenMethod() {
	case ekscontrolplanev1.EKSTokenMethodIAMAuthenticator:
		execConfig.Command = "aws-iam-authenticator"
		execConfig.Args = []string{
			"token",
			"-i",
			clusterName,
		}
	case ekscontrolplanev1.EKSTokenMethodAWSCli:
		execConfig.Command = "aws"
		execConfig.Args = []string{
			"eks",
			"get-token",
			"--cluster-name",
			clusterName,
		}
	default:
		return fmt.Errorf("using token method %s: %w", s.scope.TokenMethod(), ErrUnknownTokenMethod)
	}

	config.AuthInfos = map[string]*api.AuthInfo{
		userName: {
			Exec: execConfig,
		},
	}

	out, err := clientcmd.Write(*config)
	if err != nil {
		return errors.Wrap(err, "failed to serialize config to yaml")
	}

	configSecret.Data[secret.KubeconfigDataName] = out

	err = s.scope.Client.Update(ctx, configSecret)
	if err != nil {
		return fmt.Errorf("updating kubeconfig secret: %w", err)
	}

	return nil
}

func (s *Service) createBaseKubeConfig(cluster *eks.Cluster, userName string) (*api.Config, error) {
	clusterName := s.scope.KubernetesClusterName()
	contextName := fmt.Sprintf("%s@%s", userName, clusterName)

	certData, err := base64.StdEncoding.DecodeString(*cluster.CertificateAuthority.Data)
	if err != nil {
		return nil, fmt.Errorf("decoding cluster CA cert: %w", err)
	}

	cfg := &api.Config{
		APIVersion: api.SchemeGroupVersion.Version,
		Clusters: map[string]*api.Cluster{
			clusterName: {
				Server:                   *cluster.Endpoint,
				CertificateAuthorityData: certData,
			},
		},
		Contexts: map[string]*api.Context{
			contextName: {
				Cluster:  clusterName,
				AuthInfo: userName,
			},
		},
		CurrentContext: contextName,
	}

	return cfg, nil
}

func (s *Service) generateToken() (string, error) {
	eksClusterName := s.scope.KubernetesClusterName()

	req, output := s.STSClient.GetCallerIdentityRequest(&sts.GetCallerIdentityInput{})
	req.HTTPRequest.Header.Add(clusterNameHeader, eksClusterName)
	s.Trace("generating token for AWS identity", "user", output.UserId, "account", output.Account, "arn", output.Arn)

	presignedURL, err := req.Presign(tokenAgeMins * time.Minute)
	if err != nil {
		return "", fmt.Errorf("presigning AWS get caller identity: %w", err)
	}

	encodedURL := base64.RawURLEncoding.EncodeToString([]byte(presignedURL))
	return fmt.Sprintf("%s%s", tokenPrefix, encodedURL), nil
}

func (s *Service) getKubeConfigUserName(clusterName string, isUser bool) string {
	if isUser {
		return fmt.Sprintf("%s-user", clusterName)
	}

	return fmt.Sprintf("%s-capi-admin", clusterName)
}
