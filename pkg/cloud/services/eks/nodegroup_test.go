package eks

import (
	"github.com/aws/aws-sdk-go/service/eks"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/cluster-api-provider-aws/exp/api/v1beta1"
	"sigs.k8s.io/cluster-api-provider-aws/pkg/cloud/scope"
	"testing"
)

func TestSetStatus(t *testing.T) {
	g := NewWithT(t)
	degraded := eks.NodegroupStatusDegraded
	code := eks.NodegroupIssueCodeAsgInstanceLaunchFailures
	message := "failed due to vcpu limits"
	resourceId := "my-worker-nodes"

	s := &NodegroupService{
		scope: &scope.ManagedMachinePoolScope{
			ManagedMachinePool: &v1beta1.AWSManagedMachinePool{
				Status: v1beta1.AWSManagedMachinePoolStatus{
					Ready: false,
				},
			},
		},
	}

	issue := &eks.Issue{
		Code:        &code,
		Message:     &message,
		ResourceIds: []*string{&resourceId},
	}
	ng := &eks.Nodegroup{
		Status: &degraded,
		Health: &eks.NodegroupHealth{
			Issues: []*eks.Issue{issue},
		},
	}

	err := s.setStatus(ng)
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(ContainSubstring(issue.GoString()))
}
