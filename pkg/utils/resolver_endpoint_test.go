package utils

import (
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/onsi/gomega"
	"strings"
	"testing"
)

func TestResolverEndpointAWSGov(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	resolver := CustomEndpointResolverForAWS()
	result, err := resolver(endpoints.CloudformationServiceID, endpoints.UsGovEast1RegionID, endpoints.UseFIPSEndpointOption)
	g.Expect(err).ToNot(gomega.HaveOccurred())
	g.Expect(strings.Contains(result.URL, "cloudformation.us-gov-east-1.amazonaws.com")).To(gomega.BeTrue())

	result, err = resolver(endpoints.CloudformationServiceID, endpoints.UsGovWest1RegionID, endpoints.UseFIPSEndpointOption)
	g.Expect(err).ToNot(gomega.HaveOccurred())

	g.Expect(strings.Contains(result.URL, "cloudformation.us-gov-west-1.amazonaws.com")).To(gomega.BeTrue())

	result, err = resolver(endpoints.StsServiceID, endpoints.UsGovWest1RegionID, endpoints.UseFIPSEndpointOption)
	g.Expect(err).ToNot(gomega.HaveOccurred())

	g.Expect(strings.Contains(result.URL, "sts.us-gov-west-1.amazonaws.com")).To(gomega.BeTrue())

	result, err = resolver(endpoints.IamServiceID, endpoints.UsGovWest1RegionID, endpoints.UseFIPSEndpointOption)
	g.Expect(err).ToNot(gomega.HaveOccurred())

	g.Expect(strings.Contains(result.URL, "iam.us-gov.amazonaws.com")).To(gomega.BeTrue())
}
