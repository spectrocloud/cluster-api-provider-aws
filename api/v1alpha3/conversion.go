package v1alpha3

import (
	"k8s.io/apimachinery/pkg/conversion"
	"sigs.k8s.io/cluster-api-provider-aws/api/v1beta1"
)

func Convert_v1beta1_IngressRule_To_v1alpha3_IngressRule(rule *v1beta1.IngressRule, rule2 *IngressRule, scope conversion.Scope) error {
	return autoConvert_v1beta1_IngressRule_To_v1alpha3_IngressRule(rule, rule2, scope)
}
