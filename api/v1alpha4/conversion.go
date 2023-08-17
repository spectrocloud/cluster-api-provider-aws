/*
Copyright 2021 The Kubernetes Authors.

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

package v1alpha4

import (
	"k8s.io/apimachinery/pkg/conversion"

	"sigs.k8s.io/cluster-api-provider-aws/api/v1beta1"
)

func Convert_v1beta1_AWSClusterSpec_To_v1alpha4_AWSClusterSpec(in *v1beta1.AWSClusterSpec, out *AWSClusterSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_AWSClusterSpec_To_v1alpha4_AWSClusterSpec(in, out, s)
}

func Convert_v1beta1_IngressRule_To_v1alpha4_IngressRule(in *v1beta1.IngressRule, out *IngressRule, s conversion.Scope) error {
	return autoConvert_v1beta1_IngressRule_To_v1alpha4_IngressRule(in, out, s)
}
