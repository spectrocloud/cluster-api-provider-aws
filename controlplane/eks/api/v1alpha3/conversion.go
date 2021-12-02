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

package v1alpha3

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiconversion "k8s.io/apimachinery/pkg/conversion"
	infrav1alpha3 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha3"
	infrav1alpha4 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha4"
	"sigs.k8s.io/cluster-api-provider-aws/controlplane/eks/api/v1alpha4"
	clusterapiapiv1alpha3 "sigs.k8s.io/cluster-api/api/v1alpha3"
	clusterapiapiv1alpha4 "sigs.k8s.io/cluster-api/api/v1alpha4"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts the v1alpha3 AWSManagedControlPlane receiver to a v1alpha4 AWSManagedControlPlane.
func (r *AWSManagedControlPlane) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.AWSManagedControlPlane)

	if err := Convert_v1alpha3_AWSManagedControlPlane_To_v1alpha4_AWSManagedControlPlane(r, dst, nil); err != nil {
		return err
	}

	if r.Spec.OIDCIdentityProviderConfig != nil {
		if err := v1.Convert_Pointer_string_To_string(&r.Spec.OIDCIdentityProviderConfig.ClientId,
			&dst.Spec.OIDCIdentityProviderConfig.ClientID, nil); err != nil {
			return err
		}

		if err := v1.Convert_Pointer_string_To_string(&r.Spec.OIDCIdentityProviderConfig.IssuerUrl,
			&dst.Spec.OIDCIdentityProviderConfig.IssuerURL, nil); err != nil {
			return err
		}
	}

	return nil
}

// ConvertFrom converts the v1alpha4 AWSManagedControlPlane receiver to a v1alpha3 AWSManagedControlPlane.
func (r *AWSManagedControlPlane) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.AWSManagedControlPlane)

	if err := Convert_v1alpha4_AWSManagedControlPlane_To_v1alpha3_AWSManagedControlPlane(src, r, nil); err != nil {
		return err
	}

	if r.Spec.OIDCIdentityProviderConfig != nil {
		if err := v1.Convert_string_To_Pointer_string(&src.Spec.OIDCIdentityProviderConfig.ClientID,
			&r.Spec.OIDCIdentityProviderConfig.ClientId, nil); err != nil {
			return err
		}

		if err := v1.Convert_string_To_Pointer_string(&src.Spec.OIDCIdentityProviderConfig.IssuerURL,
			&r.Spec.OIDCIdentityProviderConfig.IssuerUrl, nil); err != nil {
			return err
		}
	}

	return nil
}

// ConvertTo converts the v1alpha3 AWSManagedControlPlaneList receiver to a v1alpha4 AWSManagedControlPlaneList.
func (r *AWSManagedControlPlaneList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.AWSManagedControlPlaneList)

	return Convert_v1alpha3_AWSManagedControlPlaneList_To_v1alpha4_AWSManagedControlPlaneList(r, dst, nil)
}

// ConvertFrom converts the v1alpha4 AWSManagedControlPlaneList receiver to a v1alpha3 AWSManagedControlPlaneList.
func (r *AWSManagedControlPlaneList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.AWSManagedControlPlaneList)

	return Convert_v1alpha4_AWSManagedControlPlaneList_To_v1alpha3_AWSManagedControlPlaneList(src, r, nil)
}

// Convert_v1alpha3_APIEndpoint_To_v1alpha4_APIEndpoint is an autogenerated conversion function.
func Convert_v1alpha3_APIEndpoint_To_v1alpha4_APIEndpoint(in *clusterapiapiv1alpha3.APIEndpoint, out *clusterapiapiv1alpha4.APIEndpoint, s apiconversion.Scope) error {
	return clusterapiapiv1alpha3.Convert_v1alpha3_APIEndpoint_To_v1alpha4_APIEndpoint(in, out, s)
}

// Convert_v1alpha4_APIEndpoint_To_v1alpha3_APIEndpoint is an autogenerated conversion function.
func Convert_v1alpha4_APIEndpoint_To_v1alpha3_APIEndpoint(in *clusterapiapiv1alpha4.APIEndpoint, out *clusterapiapiv1alpha3.APIEndpoint, s apiconversion.Scope) error {
	return clusterapiapiv1alpha3.Convert_v1alpha4_APIEndpoint_To_v1alpha3_APIEndpoint(in, out, s)
}

// Convert_v1alpha3_Bastion_To_v1alpha4_Bastion is an autogenerated conversion function.
func Convert_v1alpha3_Bastion_To_v1alpha4_Bastion(in *infrav1alpha3.Bastion, out *infrav1alpha4.Bastion, s apiconversion.Scope) error {
	return infrav1alpha3.Convert_v1alpha3_Bastion_To_v1alpha4_Bastion(in, out, s)
}

// Convert_v1alpha4_Bastion_To_v1alpha3_Bastion is an autogenerated conversion function.
func Convert_v1alpha4_Bastion_To_v1alpha3_Bastion(in *infrav1alpha4.Bastion, out *infrav1alpha3.Bastion, s apiconversion.Scope) error {
	return infrav1alpha3.Convert_v1alpha4_Bastion_To_v1alpha3_Bastion(in, out, s)
}

// Convert_v1alpha3_NetworkSpec_To_v1alpha4_NetworkSpec is an autogenerated conversion function.
func Convert_v1alpha3_NetworkSpec_To_v1alpha4_NetworkSpec(in *infrav1alpha3.NetworkSpec, out *infrav1alpha4.NetworkSpec, s apiconversion.Scope) error {
	return infrav1alpha3.Convert_v1alpha3_NetworkSpec_To_v1alpha4_NetworkSpec(in, out, s)
}

// Convert_v1alpha4_NetworkSpec_To_v1alpha3_NetworkSpec is an autogenerated conversion function.
func Convert_v1alpha4_NetworkSpec_To_v1alpha3_NetworkSpec(in *infrav1alpha4.NetworkSpec, out *infrav1alpha3.NetworkSpec, s apiconversion.Scope) error {
	return infrav1alpha3.Convert_v1alpha4_NetworkSpec_To_v1alpha3_NetworkSpec(in, out, s)
}

// Convert_v1alpha4_Instance_To_v1alpha3_Instance is an autogenerated conversion function.
func Convert_v1alpha4_Instance_To_v1alpha3_Instance(in *infrav1alpha4.Instance, out *infrav1alpha3.Instance, s apiconversion.Scope) error {
	return infrav1alpha3.Convert_v1alpha4_Instance_To_v1alpha3_Instance(in, out, s)
}

// Convert_v1alpha3_Instance_To_v1alpha4_Instance is an autogenerated conversion function.
func Convert_v1alpha3_Instance_To_v1alpha4_Instance(in *infrav1alpha3.Instance, out *infrav1alpha4.Instance, s apiconversion.Scope) error {
	return infrav1alpha3.Convert_v1alpha3_Instance_To_v1alpha4_Instance(in, out, s)
}

// Convert_v1alpha3_Network_To_v1alpha4_NetworkStatus is an autogenerated conversion function.
func Convert_v1alpha3_Network_To_v1alpha4_NetworkStatus(in *infrav1alpha3.Network, out *infrav1alpha4.NetworkStatus, s apiconversion.Scope) error {
	return infrav1alpha3.Convert_v1alpha3_Network_To_v1alpha4_NetworkStatus(in, out, s)
}

// Convert_v1alpha4_NetworkStatus_To_v1alpha3_Network is an autogenerated conversion function.
func Convert_v1alpha4_NetworkStatus_To_v1alpha3_Network(in *infrav1alpha4.NetworkStatus, out *infrav1alpha3.Network, s apiconversion.Scope) error {
	return infrav1alpha3.Convert_v1alpha4_NetworkStatus_To_v1alpha3_Network(in, out, s)
}

// Convert_v1alpha4_OIDCIdentityProviderConfig_To_v1alpha3_OIDCIdentityProviderConfig is an autogenerated conversion function.
func Convert_v1alpha4_OIDCIdentityProviderConfig_To_v1alpha3_OIDCIdentityProviderConfig(in *v1alpha4.OIDCIdentityProviderConfig, out *OIDCIdentityProviderConfig, s apiconversion.Scope) error {
	return autoConvert_v1alpha4_OIDCIdentityProviderConfig_To_v1alpha3_OIDCIdentityProviderConfig(in, out, s)
}

// Convert_v1alpha3_OIDCIdentityProviderConfig_To_v1alpha4_OIDCIdentityProviderConfig is an autogenerated conversion function.
func Convert_v1alpha3_OIDCIdentityProviderConfig_To_v1alpha4_OIDCIdentityProviderConfig(in *OIDCIdentityProviderConfig, out *v1alpha4.OIDCIdentityProviderConfig, s apiconversion.Scope) error {
	return autoConvert_v1alpha3_OIDCIdentityProviderConfig_To_v1alpha4_OIDCIdentityProviderConfig(in, out, s)
}
