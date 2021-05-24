package eks

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/pkg/errors"
	"sigs.k8s.io/cluster-api-provider-aws/controlplane/eks/api/v1alpha3"
	"sigs.k8s.io/cluster-api-provider-aws/pkg/cloud/converters"
	"sigs.k8s.io/cluster-api-provider-aws/pkg/eks/identity_provider"
)

func (s *Service) reconcileIdentityProvider(ctx context.Context) error {
	s.scope.Info("begin identity provider reconcile")
	if s.scope.ControlPlane.Spec.OIDCIdentityProviderConfig == nil {
		return nil
	}

	clusterName := s.scope.KubernetesClusterName()	
	current, err := s.getAssociatedIdentityProvider(ctx, clusterName)
	if err != nil {
		return errors.Wrap(err, "unable to list associated identity providers")
	}
	
	desired := s.convertSDKToIdentityProvider(s.scope.OIDCIdentityProviderConfig())
	
	if desired == nil && current == nil {
		s.scope.Info("no identity provider required or installed, no action needed")
		return nil
	}

	identityProviderPlan := identity_provider.NewPlan(clusterName, current, desired, s.EKSClient)

	procedures, err := identityProviderPlan.Create(ctx)
	if err != nil {
		s.scope.Error(err, "failed creating eks identity provider plane")
		return fmt.Errorf("creating eks identity provider plan: %w", err)
	}
	s.scope.V(2).Info("computed EKS identity provider plan", "numprocs", len(procedures))

	// Perform required operations
	for _, procedure := range procedures {
		s.scope.V(2).Info("Executing identity provider procedure", "name", procedure.Name())
		if err := procedure.Do(ctx); err != nil {
			s.scope.Error(err, "failed executing identity provider procedure", "name", procedure.Name())
			return fmt.Errorf("%s: %w", procedure.Name(), err)
		}
	}

	return nil
}

func (s *Service) getAssociatedIdentityProvider(ctx context.Context, clusterName string) (*identity_provider.OidcIdentityProviderConfig, error) {
	list, err := s.EKSClient.ListIdentityProviderConfigsWithContext(ctx, &eks.ListIdentityProviderConfigsInput{
		ClusterName: aws.String(clusterName),
	})
	if err != nil {
		return nil, err
	}

	// these is only one identity provider

	if len(list.IdentityProviderConfigs) == 0 {
		return nil, nil
	}

	providerconfig, err := s.EKSClient.DescribeIdentityProviderConfigWithContext(ctx, &eks.DescribeIdentityProviderConfigInput{
		ClusterName:            aws.String(clusterName),
		IdentityProviderConfig: list.IdentityProviderConfigs[0],
	})

	if err != nil {
		return nil, err
	}

	config := providerconfig.IdentityProviderConfig.Oidc
	
	return &identity_provider.OidcIdentityProviderConfig{
		ClientId:                   config.ClientId,
		GroupsClaim:                config.GroupsClaim,
		GroupsPrefix:               config.GroupsPrefix,
		IdentityProviderConfigArn:  config.IdentityProviderConfigArn,
		IdentityProviderConfigName: config.IdentityProviderConfigName,
		IssuerUrl:                  config.IssuerUrl,
		RequiredClaims:             config.RequiredClaims,
		Status:                     config.Status,
		Tags:                       converters.MapPtrToMap(config.Tags),
		UsernameClaim:              config.UsernameClaim,
		UsernamePrefix:             config.UsernamePrefix,
	}, nil
}

func (s *Service)convertSDKToIdentityProvider(in *v1alpha3.OIDCIdentityProviderConfig) *identity_provider.OidcIdentityProviderConfig {
	if in != nil {
		return &identity_provider.OidcIdentityProviderConfig{
			ClientId:                   in.ClientId,
			GroupsClaim:                in.GroupsClaim,
			GroupsPrefix:               in.GroupsPrefix,
			IdentityProviderConfigName: in.IdentityProviderConfigName,
			IssuerUrl:                  in.IssuerUrl,
			RequiredClaims:             converters.MapToMapPtr(in.RequiredClaims),
			Tags:                       in.Tags,
			UsernameClaim:              in.UsernameClaim,
			UsernamePrefix:             in.UsernamePrefix,
		}
	}

	return nil
}