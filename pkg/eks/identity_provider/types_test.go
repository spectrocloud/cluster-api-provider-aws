package identity_provider

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/onsi/gomega"
	"testing"
)

func TestIdentityProviderEqual(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	tests := []struct{
		orig *OidcIdentityProviderConfig
		other *OidcIdentityProviderConfig
	}{
		{
			orig: &OidcIdentityProviderConfig{
				ClientId:                   aws.String("a"),
				IdentityProviderConfigName: aws.String("b"),
				IssuerUrl:                  aws.String("c"),
			},
			other: &OidcIdentityProviderConfig{
				ClientId:                   aws.String("a"),
				IdentityProviderConfigName: aws.String("b"),
				IssuerUrl:                  aws.String("c"),
				Status: aws.String("e"),
			},
		},
	}

	for _, test := range tests {
		g.Expect(test.orig.IsEqual(test.other)).To(gomega.BeTrue())
	}
}
