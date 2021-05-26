package identity_provider

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/aws/aws-sdk-go/service/eks/eksiface"
	"sigs.k8s.io/cluster-api-provider-aws/pkg/planner"
)

// NewPlan creates a new Plan to manage EKS addons
func NewPlan(clusterName string, currentIdentityProvider, desiredIdentityProvider *OidcIdentityProviderConfig, client eksiface.EKSAPI) planner.Plan {
	return &plan{
		currentIdentityProvider: currentIdentityProvider,
		desiredIdentityProvider:   desiredIdentityProvider,
		eksClient:       client,
		clusterName:     clusterName,
	}
}

// Plan is a plan that will manage EKS addons
type plan struct {
	currentIdentityProvider *OidcIdentityProviderConfig
	desiredIdentityProvider   *OidcIdentityProviderConfig
	eksClient       eksiface.EKSAPI
	clusterName     string
}


func (p *plan) Create(ctx context.Context) ([]planner.Procedure, error) {
	var procedures []planner.Procedure

	if p.desiredIdentityProvider == nil && p.currentIdentityProvider == nil {
		return procedures, nil
	}

	// no config is mentioned deleted provider if we have one
	if p.desiredIdentityProvider == nil {
		// disassociation will also also trigger deletion hence
		// we do nothing in case of ConfigStatusDeleting as it will happen eventually
		if aws.StringValue(p.currentIdentityProvider.Status) == eks.ConfigStatusActive {
			procedures = append(procedures, &DisassociateIdentityProviderConfig{plan: p})
		}

		return procedures, nil
	}

	// create case
	if p.currentIdentityProvider == nil {
		procedures  = append(procedures, &AssociateIdentityProviderProcedure{plan: p})
		return procedures, nil
	}

	if p.currentIdentityProvider.IsEqual(p.desiredIdentityProvider) {
		tagsDiff := p.currentIdentityProvider.Tags.Difference(p.desiredIdentityProvider.Tags)
		if len(tagsDiff) > 0 {
			procedures = append(procedures, &UpdatedIdentityProviderTagsProcedure{plan: p})
		}
		switch aws.StringValue(p.currentIdentityProvider.Status) {
		case eks.ConfigStatusActive:
			// config active no work to be done
			return procedures, nil
		case eks.ConfigStatusCreating:
			// no change need wait for association to complete
			procedures = append(procedures, &WaitIdentityProviderAssociatedProcudure{plan: p})
		}
	} else {
		procedures = append(procedures, &DisassociateIdentityProviderConfig{plan: p})
	}

	return procedures, nil
}