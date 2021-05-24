package identity_provider

import (
	"reflect"
	infrav1 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha3"
)

// An object that represents the configuration for an OpenID Connect (OIDC)
// identity provider.
type OidcIdentityProviderConfig struct {
	ClientId *string
	GroupsClaim *string
	GroupsPrefix *string
	IdentityProviderConfigArn *string
	IdentityProviderConfigName *string
	IssuerUrl *string
	RequiredClaims map[string]*string
	Status *string
	Tags infrav1.Tags
	UsernameClaim *string
	UsernamePrefix *string
}

func (o *OidcIdentityProviderConfig) IsEqual(other *OidcIdentityProviderConfig) bool {
	if o == other {
		return true
	}

	if !reflect.DeepEqual(o.ClientId, other.ClientId) {
		return false
	}

	if !reflect.DeepEqual(o.GroupsClaim, other.GroupsClaim) {
		return false
	}

	if !reflect.DeepEqual(o.GroupsPrefix, other.GroupsPrefix) {
		return false
	}

	if !reflect.DeepEqual(o.IdentityProviderConfigName, other.IdentityProviderConfigName) {
		return false
	}

	if !reflect.DeepEqual(o.IssuerUrl, other.IssuerUrl) {
		return false
	}

	if !reflect.DeepEqual(o.RequiredClaims, other.RequiredClaims) {
		return false
	}

	if !reflect.DeepEqual(o.Tags, other.Tags) {
		return false
	}

	if !reflect.DeepEqual(o.UsernameClaim, other.UsernameClaim) {
		return false
	}

	if !reflect.DeepEqual(o.UsernamePrefix, other.UsernamePrefix) {
		return false
	}

	return true
}
