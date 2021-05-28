package identity_provider

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/pkg/errors"
	"sigs.k8s.io/cluster-api-provider-aws/pkg/cloud/services/wait"
)

var oidcType = aws.String("oidc")

type WaitIdentityProviderAssociatedProcudure struct {
	plan                 *plan
	identityProviderName string
}

func (w *WaitIdentityProviderAssociatedProcudure) Name() string {
	return "wait_identity_provider_association"
}

func (w *WaitIdentityProviderAssociatedProcudure) Do(ctx context.Context) error {
	if err := wait.WaitForWithRetryable(wait.NewBackoff(), func() (bool, error) {
		out, err := w.plan.eksClient.DescribeIdentityProviderConfigWithContext(ctx, &eks.DescribeIdentityProviderConfigInput{
			ClusterName: aws.String(w.plan.clusterName),
			IdentityProviderConfig: &eks.IdentityProviderConfig{
				Name: w.plan.currentIdentityProvider.IdentityProviderConfigName,
				Type: oidcType,
			},
		})

		if err != nil {
			return false, err
		}

		if aws.StringValue(out.IdentityProviderConfig.Oidc.Status) == eks.ConfigStatusActive {
			return true, nil
		}

		return false, nil
	}); err != nil {
		return errors.Wrap(err, "failed waiting for identity provider association to be ready")
	}

	return nil
}

type DisassociateIdentityProviderConfig struct {
	plan *plan
}

func (d *DisassociateIdentityProviderConfig) Name() string {
	return "dissociate_identity_provider"
}

func (d *DisassociateIdentityProviderConfig) Do(ctx context.Context) error {
	if err := wait.WaitForWithRetryable(wait.NewBackoff(), func() (bool, error) {
		_, err := d.plan.eksClient.DisassociateIdentityProviderConfigWithContext(ctx, &eks.DisassociateIdentityProviderConfigInput{
			ClusterName: aws.String(d.plan.clusterName),
			IdentityProviderConfig: &eks.IdentityProviderConfig{
				Name: d.plan.currentIdentityProvider.IdentityProviderConfigName,
				Type: oidcType,
			},
		})

		if err != nil {
			return false, err
		}

		return true, nil
	}); err != nil {
		return errors.Wrap(err, "failing disassociating identity provider config")
	}

	return nil
}

type AssociateIdentityProviderProcedure struct {
	plan *plan
}

func (a *AssociateIdentityProviderProcedure) Name() string {
	return "associate_identity_provider"
}

func (a *AssociateIdentityProviderProcedure) Do(ctx context.Context) error {
	oidc := a.plan.desiredIdentityProvider
	input := &eks.AssociateIdentityProviderConfigInput{
		ClusterName: aws.String(a.plan.clusterName),
		Oidc: &eks.OidcIdentityProviderConfigRequest{
			ClientId:                   oidc.ClientId,
			GroupsClaim:                oidc.GroupsClaim,
			GroupsPrefix:               oidc.GroupsPrefix,
			IdentityProviderConfigName: oidc.IdentityProviderConfigName,
			IssuerUrl:                  oidc.IssuerUrl,
			RequiredClaims:             oidc.RequiredClaims,
			UsernameClaim:              oidc.UsernameClaim,
			UsernamePrefix:             oidc.UsernamePrefix,
		},
	}

	if len(oidc.Tags) > 0 {
		input.Tags = aws.StringMap(oidc.Tags)
	}

	_, err := a.plan.eksClient.AssociateIdentityProviderConfigWithContext(ctx, input)
	if err != nil {
		return errors.Wrap(err, "failed associating identity provider")
	}

	return nil
}

type UpdatedIdentityProviderTagsProcedure struct {
	plan *plan
}

func (u *UpdatedIdentityProviderTagsProcedure) Name() string {
	return "update_identity_provider_tags"
}

func (u *UpdatedIdentityProviderTagsProcedure) Do(ctx context.Context) error {
	arn := u.plan.currentIdentityProvider.IdentityProviderConfigArn
	_, err := u.plan.eksClient.TagResource(&eks.TagResourceInput{
		ResourceArn: arn,
		Tags:        aws.StringMap(u.plan.desiredIdentityProvider.Tags),
	})

	if err != nil {
		return errors.Wrap(err, "updating identity provider tags")
	}

	return nil
}

type RemoveIdentityProviderTagsProcedure struct {
	plan *plan
}

func (r *RemoveIdentityProviderTagsProcedure) Name() string {
	return "remove_identity_provider_tags"
}

func (r *RemoveIdentityProviderTagsProcedure) Do(ctx context.Context) error {
	keys := make([]*string, 0, len(r.plan.currentIdentityProvider.Tags))

	for key := range r.plan.currentIdentityProvider.Tags {
		keys = append(keys, aws.String(key))
	}
	_, err := r.plan.eksClient.UntagResource(&eks.UntagResourceInput{
		ResourceArn: r.plan.currentIdentityProvider.IdentityProviderConfigArn,
		TagKeys:     keys,
	})

	if err != nil {
		return errors.Wrap(err, "untagging identity provider")
	}
	return nil
}