// +build !ignore_autogenerated_conversions

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha3

import (
	unsafe "unsafe"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1alpha3 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha3"
	apiv1beta1 "sigs.k8s.io/cluster-api-provider-aws/api/v1beta1"
	v1beta1 "sigs.k8s.io/cluster-api-provider-aws/controlplane/eks/api/v1beta1"
	clusterapiapiv1alpha3 "sigs.k8s.io/cluster-api/api/v1alpha3"
	clusterapiapiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*AWSManagedControlPlane)(nil), (*v1beta1.AWSManagedControlPlane)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_AWSManagedControlPlane_To_v1beta1_AWSManagedControlPlane(a.(*AWSManagedControlPlane), b.(*v1beta1.AWSManagedControlPlane), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.AWSManagedControlPlane)(nil), (*AWSManagedControlPlane)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AWSManagedControlPlane_To_v1alpha3_AWSManagedControlPlane(a.(*v1beta1.AWSManagedControlPlane), b.(*AWSManagedControlPlane), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AWSManagedControlPlaneList)(nil), (*v1beta1.AWSManagedControlPlaneList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_AWSManagedControlPlaneList_To_v1beta1_AWSManagedControlPlaneList(a.(*AWSManagedControlPlaneList), b.(*v1beta1.AWSManagedControlPlaneList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.AWSManagedControlPlaneList)(nil), (*AWSManagedControlPlaneList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AWSManagedControlPlaneList_To_v1alpha3_AWSManagedControlPlaneList(a.(*v1beta1.AWSManagedControlPlaneList), b.(*AWSManagedControlPlaneList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AWSManagedControlPlaneSpec)(nil), (*v1beta1.AWSManagedControlPlaneSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_AWSManagedControlPlaneSpec_To_v1beta1_AWSManagedControlPlaneSpec(a.(*AWSManagedControlPlaneSpec), b.(*v1beta1.AWSManagedControlPlaneSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AWSManagedControlPlaneStatus)(nil), (*v1beta1.AWSManagedControlPlaneStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_AWSManagedControlPlaneStatus_To_v1beta1_AWSManagedControlPlaneStatus(a.(*AWSManagedControlPlaneStatus), b.(*v1beta1.AWSManagedControlPlaneStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Addon)(nil), (*v1beta1.Addon)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_Addon_To_v1beta1_Addon(a.(*Addon), b.(*v1beta1.Addon), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.Addon)(nil), (*Addon)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Addon_To_v1alpha3_Addon(a.(*v1beta1.Addon), b.(*Addon), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AddonIssue)(nil), (*v1beta1.AddonIssue)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_AddonIssue_To_v1beta1_AddonIssue(a.(*AddonIssue), b.(*v1beta1.AddonIssue), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.AddonIssue)(nil), (*AddonIssue)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AddonIssue_To_v1alpha3_AddonIssue(a.(*v1beta1.AddonIssue), b.(*AddonIssue), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AddonState)(nil), (*v1beta1.AddonState)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_AddonState_To_v1beta1_AddonState(a.(*AddonState), b.(*v1beta1.AddonState), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.AddonState)(nil), (*AddonState)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AddonState_To_v1alpha3_AddonState(a.(*v1beta1.AddonState), b.(*AddonState), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControlPlaneLoggingSpec)(nil), (*v1beta1.ControlPlaneLoggingSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_ControlPlaneLoggingSpec_To_v1beta1_ControlPlaneLoggingSpec(a.(*ControlPlaneLoggingSpec), b.(*v1beta1.ControlPlaneLoggingSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.ControlPlaneLoggingSpec)(nil), (*ControlPlaneLoggingSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ControlPlaneLoggingSpec_To_v1alpha3_ControlPlaneLoggingSpec(a.(*v1beta1.ControlPlaneLoggingSpec), b.(*ControlPlaneLoggingSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EncryptionConfig)(nil), (*v1beta1.EncryptionConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_EncryptionConfig_To_v1beta1_EncryptionConfig(a.(*EncryptionConfig), b.(*v1beta1.EncryptionConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.EncryptionConfig)(nil), (*EncryptionConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EncryptionConfig_To_v1alpha3_EncryptionConfig(a.(*v1beta1.EncryptionConfig), b.(*EncryptionConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EndpointAccess)(nil), (*v1beta1.EndpointAccess)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_EndpointAccess_To_v1beta1_EndpointAccess(a.(*EndpointAccess), b.(*v1beta1.EndpointAccess), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.EndpointAccess)(nil), (*EndpointAccess)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EndpointAccess_To_v1alpha3_EndpointAccess(a.(*v1beta1.EndpointAccess), b.(*EndpointAccess), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*IAMAuthenticatorConfig)(nil), (*v1beta1.IAMAuthenticatorConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_IAMAuthenticatorConfig_To_v1beta1_IAMAuthenticatorConfig(a.(*IAMAuthenticatorConfig), b.(*v1beta1.IAMAuthenticatorConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.IAMAuthenticatorConfig)(nil), (*IAMAuthenticatorConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_IAMAuthenticatorConfig_To_v1alpha3_IAMAuthenticatorConfig(a.(*v1beta1.IAMAuthenticatorConfig), b.(*IAMAuthenticatorConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubernetesMapping)(nil), (*v1beta1.KubernetesMapping)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_KubernetesMapping_To_v1beta1_KubernetesMapping(a.(*KubernetesMapping), b.(*v1beta1.KubernetesMapping), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.KubernetesMapping)(nil), (*KubernetesMapping)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubernetesMapping_To_v1alpha3_KubernetesMapping(a.(*v1beta1.KubernetesMapping), b.(*KubernetesMapping), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*OIDCProviderStatus)(nil), (*v1beta1.OIDCProviderStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_OIDCProviderStatus_To_v1beta1_OIDCProviderStatus(a.(*OIDCProviderStatus), b.(*v1beta1.OIDCProviderStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.OIDCProviderStatus)(nil), (*OIDCProviderStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_OIDCProviderStatus_To_v1alpha3_OIDCProviderStatus(a.(*v1beta1.OIDCProviderStatus), b.(*OIDCProviderStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RoleMapping)(nil), (*v1beta1.RoleMapping)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_RoleMapping_To_v1beta1_RoleMapping(a.(*RoleMapping), b.(*v1beta1.RoleMapping), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.RoleMapping)(nil), (*RoleMapping)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_RoleMapping_To_v1alpha3_RoleMapping(a.(*v1beta1.RoleMapping), b.(*RoleMapping), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*UserMapping)(nil), (*v1beta1.UserMapping)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_UserMapping_To_v1beta1_UserMapping(a.(*UserMapping), b.(*v1beta1.UserMapping), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.UserMapping)(nil), (*UserMapping)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_UserMapping_To_v1alpha3_UserMapping(a.(*v1beta1.UserMapping), b.(*UserMapping), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*apiv1alpha3.Bastion)(nil), (*apiv1beta1.Bastion)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_Bastion_To_v1beta1_Bastion(a.(*apiv1alpha3.Bastion), b.(*apiv1beta1.Bastion), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*apiv1alpha3.Instance)(nil), (*apiv1beta1.Instance)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_Instance_To_v1beta1_Instance(a.(*apiv1alpha3.Instance), b.(*apiv1beta1.Instance), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*apiv1alpha3.NetworkSpec)(nil), (*apiv1beta1.NetworkSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_NetworkSpec_To_v1beta1_NetworkSpec(a.(*apiv1alpha3.NetworkSpec), b.(*apiv1beta1.NetworkSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta1.AWSManagedControlPlaneSpec)(nil), (*AWSManagedControlPlaneSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AWSManagedControlPlaneSpec_To_v1alpha3_AWSManagedControlPlaneSpec(a.(*v1beta1.AWSManagedControlPlaneSpec), b.(*AWSManagedControlPlaneSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta1.AWSManagedControlPlaneStatus)(nil), (*AWSManagedControlPlaneStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AWSManagedControlPlaneStatus_To_v1alpha3_AWSManagedControlPlaneStatus(a.(*v1beta1.AWSManagedControlPlaneStatus), b.(*AWSManagedControlPlaneStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*apiv1beta1.Bastion)(nil), (*apiv1alpha3.Bastion)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Bastion_To_v1alpha3_Bastion(a.(*apiv1beta1.Bastion), b.(*apiv1alpha3.Bastion), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*apiv1beta1.NetworkSpec)(nil), (*apiv1alpha3.NetworkSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_NetworkSpec_To_v1alpha3_NetworkSpec(a.(*apiv1beta1.NetworkSpec), b.(*apiv1alpha3.NetworkSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha3_AWSManagedControlPlane_To_v1beta1_AWSManagedControlPlane(in *AWSManagedControlPlane, out *v1beta1.AWSManagedControlPlane, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha3_AWSManagedControlPlaneSpec_To_v1beta1_AWSManagedControlPlaneSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha3_AWSManagedControlPlaneStatus_To_v1beta1_AWSManagedControlPlaneStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_AWSManagedControlPlane_To_v1beta1_AWSManagedControlPlane is an autogenerated conversion function.
func Convert_v1alpha3_AWSManagedControlPlane_To_v1beta1_AWSManagedControlPlane(in *AWSManagedControlPlane, out *v1beta1.AWSManagedControlPlane, s conversion.Scope) error {
	return autoConvert_v1alpha3_AWSManagedControlPlane_To_v1beta1_AWSManagedControlPlane(in, out, s)
}

func autoConvert_v1beta1_AWSManagedControlPlane_To_v1alpha3_AWSManagedControlPlane(in *v1beta1.AWSManagedControlPlane, out *AWSManagedControlPlane, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_AWSManagedControlPlaneSpec_To_v1alpha3_AWSManagedControlPlaneSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_AWSManagedControlPlaneStatus_To_v1alpha3_AWSManagedControlPlaneStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_AWSManagedControlPlane_To_v1alpha3_AWSManagedControlPlane is an autogenerated conversion function.
func Convert_v1beta1_AWSManagedControlPlane_To_v1alpha3_AWSManagedControlPlane(in *v1beta1.AWSManagedControlPlane, out *AWSManagedControlPlane, s conversion.Scope) error {
	return autoConvert_v1beta1_AWSManagedControlPlane_To_v1alpha3_AWSManagedControlPlane(in, out, s)
}

func autoConvert_v1alpha3_AWSManagedControlPlaneList_To_v1beta1_AWSManagedControlPlaneList(in *AWSManagedControlPlaneList, out *v1beta1.AWSManagedControlPlaneList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta1.AWSManagedControlPlane, len(*in))
		for i := range *in {
			if err := Convert_v1alpha3_AWSManagedControlPlane_To_v1beta1_AWSManagedControlPlane(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha3_AWSManagedControlPlaneList_To_v1beta1_AWSManagedControlPlaneList is an autogenerated conversion function.
func Convert_v1alpha3_AWSManagedControlPlaneList_To_v1beta1_AWSManagedControlPlaneList(in *AWSManagedControlPlaneList, out *v1beta1.AWSManagedControlPlaneList, s conversion.Scope) error {
	return autoConvert_v1alpha3_AWSManagedControlPlaneList_To_v1beta1_AWSManagedControlPlaneList(in, out, s)
}

func autoConvert_v1beta1_AWSManagedControlPlaneList_To_v1alpha3_AWSManagedControlPlaneList(in *v1beta1.AWSManagedControlPlaneList, out *AWSManagedControlPlaneList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSManagedControlPlane, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_AWSManagedControlPlane_To_v1alpha3_AWSManagedControlPlane(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_AWSManagedControlPlaneList_To_v1alpha3_AWSManagedControlPlaneList is an autogenerated conversion function.
func Convert_v1beta1_AWSManagedControlPlaneList_To_v1alpha3_AWSManagedControlPlaneList(in *v1beta1.AWSManagedControlPlaneList, out *AWSManagedControlPlaneList, s conversion.Scope) error {
	return autoConvert_v1beta1_AWSManagedControlPlaneList_To_v1alpha3_AWSManagedControlPlaneList(in, out, s)
}

func autoConvert_v1alpha3_AWSManagedControlPlaneSpec_To_v1beta1_AWSManagedControlPlaneSpec(in *AWSManagedControlPlaneSpec, out *v1beta1.AWSManagedControlPlaneSpec, s conversion.Scope) error {
	out.EKSClusterName = in.EKSClusterName
	out.IdentityRef = (*apiv1beta1.AWSIdentityReference)(unsafe.Pointer(in.IdentityRef))
	if err := Convert_v1alpha3_NetworkSpec_To_v1beta1_NetworkSpec(&in.NetworkSpec, &out.NetworkSpec, s); err != nil {
		return err
	}
	out.SecondaryCidrBlock = (*string)(unsafe.Pointer(in.SecondaryCidrBlock))
	out.Region = in.Region
	out.SSHKeyName = (*string)(unsafe.Pointer(in.SSHKeyName))
	out.Version = (*string)(unsafe.Pointer(in.Version))
	out.RoleName = (*string)(unsafe.Pointer(in.RoleName))
	out.RoleAdditionalPolicies = (*[]string)(unsafe.Pointer(in.RoleAdditionalPolicies))
	out.Logging = (*v1beta1.ControlPlaneLoggingSpec)(unsafe.Pointer(in.Logging))
	out.EncryptionConfig = (*v1beta1.EncryptionConfig)(unsafe.Pointer(in.EncryptionConfig))
	out.AdditionalTags = *(*apiv1beta1.Tags)(unsafe.Pointer(&in.AdditionalTags))
	out.IAMAuthenticatorConfig = (*v1beta1.IAMAuthenticatorConfig)(unsafe.Pointer(in.IAMAuthenticatorConfig))
	if err := Convert_v1alpha3_EndpointAccess_To_v1beta1_EndpointAccess(&in.EndpointAccess, &out.EndpointAccess, s); err != nil {
		return err
	}
	if err := clusterapiapiv1alpha3.Convert_v1alpha3_APIEndpoint_To_v1beta1_APIEndpoint(&in.ControlPlaneEndpoint, &out.ControlPlaneEndpoint, s); err != nil {
		return err
	}
	out.ImageLookupFormat = in.ImageLookupFormat
	out.ImageLookupOrg = in.ImageLookupOrg
	out.ImageLookupBaseOS = in.ImageLookupBaseOS
	if err := Convert_v1alpha3_Bastion_To_v1beta1_Bastion(&in.Bastion, &out.Bastion, s); err != nil {
		return err
	}
	out.TokenMethod = (*v1beta1.EKSTokenMethod)(unsafe.Pointer(in.TokenMethod))
	out.AssociateOIDCProvider = in.AssociateOIDCProvider
	out.Addons = (*[]v1beta1.Addon)(unsafe.Pointer(in.Addons))
	out.DisableVPCCNI = in.DisableVPCCNI
	return nil
}

// Convert_v1alpha3_AWSManagedControlPlaneSpec_To_v1beta1_AWSManagedControlPlaneSpec is an autogenerated conversion function.
func Convert_v1alpha3_AWSManagedControlPlaneSpec_To_v1beta1_AWSManagedControlPlaneSpec(in *AWSManagedControlPlaneSpec, out *v1beta1.AWSManagedControlPlaneSpec, s conversion.Scope) error {
	return autoConvert_v1alpha3_AWSManagedControlPlaneSpec_To_v1beta1_AWSManagedControlPlaneSpec(in, out, s)
}

func autoConvert_v1beta1_AWSManagedControlPlaneSpec_To_v1alpha3_AWSManagedControlPlaneSpec(in *v1beta1.AWSManagedControlPlaneSpec, out *AWSManagedControlPlaneSpec, s conversion.Scope) error {
	out.EKSClusterName = in.EKSClusterName
	out.IdentityRef = (*apiv1alpha3.AWSIdentityReference)(unsafe.Pointer(in.IdentityRef))
	if err := Convert_v1beta1_NetworkSpec_To_v1alpha3_NetworkSpec(&in.NetworkSpec, &out.NetworkSpec, s); err != nil {
		return err
	}
	out.SecondaryCidrBlock = (*string)(unsafe.Pointer(in.SecondaryCidrBlock))
	out.Region = in.Region
	out.SSHKeyName = (*string)(unsafe.Pointer(in.SSHKeyName))
	out.Version = (*string)(unsafe.Pointer(in.Version))
	out.RoleName = (*string)(unsafe.Pointer(in.RoleName))
	out.RoleAdditionalPolicies = (*[]string)(unsafe.Pointer(in.RoleAdditionalPolicies))
	out.Logging = (*ControlPlaneLoggingSpec)(unsafe.Pointer(in.Logging))
	out.EncryptionConfig = (*EncryptionConfig)(unsafe.Pointer(in.EncryptionConfig))
	out.AdditionalTags = *(*apiv1alpha3.Tags)(unsafe.Pointer(&in.AdditionalTags))
	out.IAMAuthenticatorConfig = (*IAMAuthenticatorConfig)(unsafe.Pointer(in.IAMAuthenticatorConfig))
	if err := Convert_v1beta1_EndpointAccess_To_v1alpha3_EndpointAccess(&in.EndpointAccess, &out.EndpointAccess, s); err != nil {
		return err
	}
	if err := clusterapiapiv1alpha3.Convert_v1beta1_APIEndpoint_To_v1alpha3_APIEndpoint(&in.ControlPlaneEndpoint, &out.ControlPlaneEndpoint, s); err != nil {
		return err
	}
	out.ImageLookupFormat = in.ImageLookupFormat
	out.ImageLookupOrg = in.ImageLookupOrg
	out.ImageLookupBaseOS = in.ImageLookupBaseOS
	if err := Convert_v1beta1_Bastion_To_v1alpha3_Bastion(&in.Bastion, &out.Bastion, s); err != nil {
		return err
	}
	out.TokenMethod = (*EKSTokenMethod)(unsafe.Pointer(in.TokenMethod))
	out.AssociateOIDCProvider = in.AssociateOIDCProvider
	out.Addons = (*[]Addon)(unsafe.Pointer(in.Addons))
	if in.OIDCIdentityProviderConfig != nil {
		in, out := &in.OIDCIdentityProviderConfig, &out.OIDCIdentityProviderConfig
		*out = new(OIDCIdentityProviderConfig)
		if err := Convert_v1alpha4_OIDCIdentityProviderConfig_To_v1alpha3_OIDCIdentityProviderConfig(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.OIDCIdentityProviderConfig = nil
	}
	out.DisableVPCCNI = in.DisableVPCCNI
	return nil
}

func autoConvert_v1alpha3_AWSManagedControlPlaneStatus_To_v1beta1_AWSManagedControlPlaneStatus(in *AWSManagedControlPlaneStatus, out *v1beta1.AWSManagedControlPlaneStatus, s conversion.Scope) error {
	if err := apiv1alpha3.Convert_v1alpha3_Network_To_v1beta1_NetworkStatus(&in.Network, &out.Network, s); err != nil {
		return err
	}
	if in.FailureDomains != nil {
		in, out := &in.FailureDomains, &out.FailureDomains
		*out = make(clusterapiapiv1beta1.FailureDomains, len(*in))
		for key, val := range *in {
			newVal := new(clusterapiapiv1beta1.FailureDomainSpec)
			if err := clusterapiapiv1alpha3.Convert_v1alpha3_FailureDomainSpec_To_v1beta1_FailureDomainSpec(&val, newVal, s); err != nil {
				return err
			}
			(*out)[key] = *newVal
		}
	} else {
		out.FailureDomains = nil
	}
	if in.Bastion != nil {
		in, out := &in.Bastion, &out.Bastion
		*out = new(apiv1beta1.Instance)
		if err := Convert_v1alpha3_Instance_To_v1beta1_Instance(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Bastion = nil
	}
	if err := Convert_v1alpha3_OIDCProviderStatus_To_v1beta1_OIDCProviderStatus(&in.OIDCProvider, &out.OIDCProvider, s); err != nil {
		return err
	}
	out.ExternalManagedControlPlane = (*bool)(unsafe.Pointer(in.ExternalManagedControlPlane))
	out.Initialized = in.Initialized
	out.Ready = in.Ready
	out.FailureMessage = (*string)(unsafe.Pointer(in.FailureMessage))
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(clusterapiapiv1beta1.Conditions, len(*in))
		for i := range *in {
			if err := clusterapiapiv1alpha3.Convert_v1alpha3_Condition_To_v1beta1_Condition(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	out.Addons = *(*[]v1beta1.AddonState)(unsafe.Pointer(&in.Addons))
	return nil
}

// Convert_v1alpha3_AWSManagedControlPlaneStatus_To_v1beta1_AWSManagedControlPlaneStatus is an autogenerated conversion function.
func Convert_v1alpha3_AWSManagedControlPlaneStatus_To_v1beta1_AWSManagedControlPlaneStatus(in *AWSManagedControlPlaneStatus, out *v1beta1.AWSManagedControlPlaneStatus, s conversion.Scope) error {
	return autoConvert_v1alpha3_AWSManagedControlPlaneStatus_To_v1beta1_AWSManagedControlPlaneStatus(in, out, s)
}

func autoConvert_v1beta1_AWSManagedControlPlaneStatus_To_v1alpha3_AWSManagedControlPlaneStatus(in *v1beta1.AWSManagedControlPlaneStatus, out *AWSManagedControlPlaneStatus, s conversion.Scope) error {
	if err := apiv1alpha3.Convert_v1beta1_NetworkStatus_To_v1alpha3_Network(&in.Network, &out.Network, s); err != nil {
		return err
	}
	if in.FailureDomains != nil {
		in, out := &in.FailureDomains, &out.FailureDomains
		*out = make(clusterapiapiv1alpha3.FailureDomains, len(*in))
		for key, val := range *in {
			newVal := new(clusterapiapiv1alpha3.FailureDomainSpec)
			if err := clusterapiapiv1alpha3.Convert_v1beta1_FailureDomainSpec_To_v1alpha3_FailureDomainSpec(&val, newVal, s); err != nil {
				return err
			}
			(*out)[key] = *newVal
		}
	} else {
		out.FailureDomains = nil
	}
	if in.Bastion != nil {
		in, out := &in.Bastion, &out.Bastion
		*out = new(apiv1alpha3.Instance)
		if err := apiv1alpha3.Convert_v1beta1_Instance_To_v1alpha3_Instance(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Bastion = nil
	}
	if err := Convert_v1beta1_OIDCProviderStatus_To_v1alpha3_OIDCProviderStatus(&in.OIDCProvider, &out.OIDCProvider, s); err != nil {
		return err
	}
	out.ExternalManagedControlPlane = (*bool)(unsafe.Pointer(in.ExternalManagedControlPlane))
	out.Initialized = in.Initialized
	out.Ready = in.Ready
	out.FailureMessage = (*string)(unsafe.Pointer(in.FailureMessage))
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(clusterapiapiv1alpha3.Conditions, len(*in))
		for i := range *in {
			if err := clusterapiapiv1alpha3.Convert_v1beta1_Condition_To_v1alpha3_Condition(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	out.Addons = *(*[]AddonState)(unsafe.Pointer(&in.Addons))
	if err := Convert_v1alpha4_IdentityProviderStatus_To_v1alpha3_IdentityProviderStatus(&in.IdentityProviderStatus, &out.IdentityProviderStatus, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha3_Addon_To_v1beta1_Addon(in *Addon, out *v1beta1.Addon, s conversion.Scope) error {
	out.Name = in.Name
	out.Version = in.Version
	out.ConflictResolution = (*v1beta1.AddonResolution)(unsafe.Pointer(in.ConflictResolution))
	out.ServiceAccountRoleArn = (*string)(unsafe.Pointer(in.ServiceAccountRoleArn))
	return nil
}

// Convert_v1alpha3_Addon_To_v1beta1_Addon is an autogenerated conversion function.
func Convert_v1alpha3_Addon_To_v1beta1_Addon(in *Addon, out *v1beta1.Addon, s conversion.Scope) error {
	return autoConvert_v1alpha3_Addon_To_v1beta1_Addon(in, out, s)
}

func autoConvert_v1beta1_Addon_To_v1alpha3_Addon(in *v1beta1.Addon, out *Addon, s conversion.Scope) error {
	out.Name = in.Name
	out.Version = in.Version
	out.ConflictResolution = (*AddonResolution)(unsafe.Pointer(in.ConflictResolution))
	out.ServiceAccountRoleArn = (*string)(unsafe.Pointer(in.ServiceAccountRoleArn))
	return nil
}

// Convert_v1beta1_Addon_To_v1alpha3_Addon is an autogenerated conversion function.
func Convert_v1beta1_Addon_To_v1alpha3_Addon(in *v1beta1.Addon, out *Addon, s conversion.Scope) error {
	return autoConvert_v1beta1_Addon_To_v1alpha3_Addon(in, out, s)
}

func autoConvert_v1alpha3_AddonIssue_To_v1beta1_AddonIssue(in *AddonIssue, out *v1beta1.AddonIssue, s conversion.Scope) error {
	out.Code = (*string)(unsafe.Pointer(in.Code))
	out.Message = (*string)(unsafe.Pointer(in.Message))
	out.ResourceIDs = *(*[]string)(unsafe.Pointer(&in.ResourceIDs))
	return nil
}

// Convert_v1alpha3_AddonIssue_To_v1beta1_AddonIssue is an autogenerated conversion function.
func Convert_v1alpha3_AddonIssue_To_v1beta1_AddonIssue(in *AddonIssue, out *v1beta1.AddonIssue, s conversion.Scope) error {
	return autoConvert_v1alpha3_AddonIssue_To_v1beta1_AddonIssue(in, out, s)
}

func autoConvert_v1beta1_AddonIssue_To_v1alpha3_AddonIssue(in *v1beta1.AddonIssue, out *AddonIssue, s conversion.Scope) error {
	out.Code = (*string)(unsafe.Pointer(in.Code))
	out.Message = (*string)(unsafe.Pointer(in.Message))
	out.ResourceIDs = *(*[]string)(unsafe.Pointer(&in.ResourceIDs))
	return nil
}

// Convert_v1beta1_AddonIssue_To_v1alpha3_AddonIssue is an autogenerated conversion function.
func Convert_v1beta1_AddonIssue_To_v1alpha3_AddonIssue(in *v1beta1.AddonIssue, out *AddonIssue, s conversion.Scope) error {
	return autoConvert_v1beta1_AddonIssue_To_v1alpha3_AddonIssue(in, out, s)
}

func autoConvert_v1alpha3_AddonState_To_v1beta1_AddonState(in *AddonState, out *v1beta1.AddonState, s conversion.Scope) error {
	out.Name = in.Name
	out.Version = in.Version
	out.ARN = in.ARN
	out.ServiceAccountRoleArn = (*string)(unsafe.Pointer(in.ServiceAccountRoleArn))
	out.CreatedAt = in.CreatedAt
	out.ModifiedAt = in.ModifiedAt
	out.Status = (*string)(unsafe.Pointer(in.Status))
	out.Issues = *(*[]v1beta1.AddonIssue)(unsafe.Pointer(&in.Issues))
	return nil
}

// Convert_v1alpha3_AddonState_To_v1beta1_AddonState is an autogenerated conversion function.
func Convert_v1alpha3_AddonState_To_v1beta1_AddonState(in *AddonState, out *v1beta1.AddonState, s conversion.Scope) error {
	return autoConvert_v1alpha3_AddonState_To_v1beta1_AddonState(in, out, s)
}

func autoConvert_v1beta1_AddonState_To_v1alpha3_AddonState(in *v1beta1.AddonState, out *AddonState, s conversion.Scope) error {
	out.Name = in.Name
	out.Version = in.Version
	out.ARN = in.ARN
	out.ServiceAccountRoleArn = (*string)(unsafe.Pointer(in.ServiceAccountRoleArn))
	out.CreatedAt = in.CreatedAt
	out.ModifiedAt = in.ModifiedAt
	out.Status = (*string)(unsafe.Pointer(in.Status))
	out.Issues = *(*[]AddonIssue)(unsafe.Pointer(&in.Issues))
	return nil
}

// Convert_v1beta1_AddonState_To_v1alpha3_AddonState is an autogenerated conversion function.
func Convert_v1beta1_AddonState_To_v1alpha3_AddonState(in *v1beta1.AddonState, out *AddonState, s conversion.Scope) error {
	return autoConvert_v1beta1_AddonState_To_v1alpha3_AddonState(in, out, s)
}

func autoConvert_v1alpha3_ControlPlaneLoggingSpec_To_v1beta1_ControlPlaneLoggingSpec(in *ControlPlaneLoggingSpec, out *v1beta1.ControlPlaneLoggingSpec, s conversion.Scope) error {
	out.APIServer = in.APIServer
	out.Audit = in.Audit
	out.Authenticator = in.Authenticator
	out.ControllerManager = in.ControllerManager
	out.Scheduler = in.Scheduler
	return nil
}

// Convert_v1alpha3_ControlPlaneLoggingSpec_To_v1beta1_ControlPlaneLoggingSpec is an autogenerated conversion function.
func Convert_v1alpha3_ControlPlaneLoggingSpec_To_v1beta1_ControlPlaneLoggingSpec(in *ControlPlaneLoggingSpec, out *v1beta1.ControlPlaneLoggingSpec, s conversion.Scope) error {
	return autoConvert_v1alpha3_ControlPlaneLoggingSpec_To_v1beta1_ControlPlaneLoggingSpec(in, out, s)
}

func autoConvert_v1beta1_ControlPlaneLoggingSpec_To_v1alpha3_ControlPlaneLoggingSpec(in *v1beta1.ControlPlaneLoggingSpec, out *ControlPlaneLoggingSpec, s conversion.Scope) error {
	out.APIServer = in.APIServer
	out.Audit = in.Audit
	out.Authenticator = in.Authenticator
	out.ControllerManager = in.ControllerManager
	out.Scheduler = in.Scheduler
	return nil
}

// Convert_v1beta1_ControlPlaneLoggingSpec_To_v1alpha3_ControlPlaneLoggingSpec is an autogenerated conversion function.
func Convert_v1beta1_ControlPlaneLoggingSpec_To_v1alpha3_ControlPlaneLoggingSpec(in *v1beta1.ControlPlaneLoggingSpec, out *ControlPlaneLoggingSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_ControlPlaneLoggingSpec_To_v1alpha3_ControlPlaneLoggingSpec(in, out, s)
}

func autoConvert_v1alpha3_EncryptionConfig_To_v1beta1_EncryptionConfig(in *EncryptionConfig, out *v1beta1.EncryptionConfig, s conversion.Scope) error {
	out.Provider = (*string)(unsafe.Pointer(in.Provider))
	out.Resources = *(*[]*string)(unsafe.Pointer(&in.Resources))
	return nil
}

// Convert_v1alpha3_EncryptionConfig_To_v1beta1_EncryptionConfig is an autogenerated conversion function.
func Convert_v1alpha3_EncryptionConfig_To_v1beta1_EncryptionConfig(in *EncryptionConfig, out *v1beta1.EncryptionConfig, s conversion.Scope) error {
	return autoConvert_v1alpha3_EncryptionConfig_To_v1beta1_EncryptionConfig(in, out, s)
}

func autoConvert_v1beta1_EncryptionConfig_To_v1alpha3_EncryptionConfig(in *v1beta1.EncryptionConfig, out *EncryptionConfig, s conversion.Scope) error {
	out.Provider = (*string)(unsafe.Pointer(in.Provider))
	out.Resources = *(*[]*string)(unsafe.Pointer(&in.Resources))
	return nil
}

// Convert_v1beta1_EncryptionConfig_To_v1alpha3_EncryptionConfig is an autogenerated conversion function.
func Convert_v1beta1_EncryptionConfig_To_v1alpha3_EncryptionConfig(in *v1beta1.EncryptionConfig, out *EncryptionConfig, s conversion.Scope) error {
	return autoConvert_v1beta1_EncryptionConfig_To_v1alpha3_EncryptionConfig(in, out, s)
}

func autoConvert_v1alpha3_EndpointAccess_To_v1beta1_EndpointAccess(in *EndpointAccess, out *v1beta1.EndpointAccess, s conversion.Scope) error {
	out.Public = (*bool)(unsafe.Pointer(in.Public))
	out.PublicCIDRs = *(*[]*string)(unsafe.Pointer(&in.PublicCIDRs))
	out.Private = (*bool)(unsafe.Pointer(in.Private))
	return nil
}

// Convert_v1alpha3_EndpointAccess_To_v1beta1_EndpointAccess is an autogenerated conversion function.
func Convert_v1alpha3_EndpointAccess_To_v1beta1_EndpointAccess(in *EndpointAccess, out *v1beta1.EndpointAccess, s conversion.Scope) error {
	return autoConvert_v1alpha3_EndpointAccess_To_v1beta1_EndpointAccess(in, out, s)
}

func autoConvert_v1beta1_EndpointAccess_To_v1alpha3_EndpointAccess(in *v1beta1.EndpointAccess, out *EndpointAccess, s conversion.Scope) error {
	out.Public = (*bool)(unsafe.Pointer(in.Public))
	out.PublicCIDRs = *(*[]*string)(unsafe.Pointer(&in.PublicCIDRs))
	out.Private = (*bool)(unsafe.Pointer(in.Private))
	return nil
}

// Convert_v1beta1_EndpointAccess_To_v1alpha3_EndpointAccess is an autogenerated conversion function.
func Convert_v1beta1_EndpointAccess_To_v1alpha3_EndpointAccess(in *v1beta1.EndpointAccess, out *EndpointAccess, s conversion.Scope) error {
	return autoConvert_v1beta1_EndpointAccess_To_v1alpha3_EndpointAccess(in, out, s)
}

func autoConvert_v1alpha3_IAMAuthenticatorConfig_To_v1beta1_IAMAuthenticatorConfig(in *IAMAuthenticatorConfig, out *v1beta1.IAMAuthenticatorConfig, s conversion.Scope) error {
	out.RoleMappings = *(*[]v1beta1.RoleMapping)(unsafe.Pointer(&in.RoleMappings))
	out.UserMappings = *(*[]v1beta1.UserMapping)(unsafe.Pointer(&in.UserMappings))
	return nil
}

// Convert_v1alpha3_IAMAuthenticatorConfig_To_v1beta1_IAMAuthenticatorConfig is an autogenerated conversion function.
func Convert_v1alpha3_IAMAuthenticatorConfig_To_v1beta1_IAMAuthenticatorConfig(in *IAMAuthenticatorConfig, out *v1beta1.IAMAuthenticatorConfig, s conversion.Scope) error {
	return autoConvert_v1alpha3_IAMAuthenticatorConfig_To_v1beta1_IAMAuthenticatorConfig(in, out, s)
}

func autoConvert_v1beta1_IAMAuthenticatorConfig_To_v1alpha3_IAMAuthenticatorConfig(in *v1beta1.IAMAuthenticatorConfig, out *IAMAuthenticatorConfig, s conversion.Scope) error {
	out.RoleMappings = *(*[]RoleMapping)(unsafe.Pointer(&in.RoleMappings))
	out.UserMappings = *(*[]UserMapping)(unsafe.Pointer(&in.UserMappings))
	return nil
}

// Convert_v1beta1_IAMAuthenticatorConfig_To_v1alpha3_IAMAuthenticatorConfig is an autogenerated conversion function.
func Convert_v1beta1_IAMAuthenticatorConfig_To_v1alpha3_IAMAuthenticatorConfig(in *v1beta1.IAMAuthenticatorConfig, out *IAMAuthenticatorConfig, s conversion.Scope) error {
	return autoConvert_v1beta1_IAMAuthenticatorConfig_To_v1alpha3_IAMAuthenticatorConfig(in, out, s)
}

func autoConvert_v1alpha3_KubernetesMapping_To_v1beta1_KubernetesMapping(in *KubernetesMapping, out *v1beta1.KubernetesMapping, s conversion.Scope) error {
	out.UserName = in.UserName
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	return nil
}

// Convert_v1alpha3_KubernetesMapping_To_v1beta1_KubernetesMapping is an autogenerated conversion function.
func Convert_v1alpha3_KubernetesMapping_To_v1beta1_KubernetesMapping(in *KubernetesMapping, out *v1beta1.KubernetesMapping, s conversion.Scope) error {
	return autoConvert_v1alpha3_KubernetesMapping_To_v1beta1_KubernetesMapping(in, out, s)
}

func autoConvert_v1beta1_KubernetesMapping_To_v1alpha3_KubernetesMapping(in *v1beta1.KubernetesMapping, out *KubernetesMapping, s conversion.Scope) error {
	out.UserName = in.UserName
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	return nil
}

// Convert_v1beta1_KubernetesMapping_To_v1alpha3_KubernetesMapping is an autogenerated conversion function.
func Convert_v1beta1_KubernetesMapping_To_v1alpha3_KubernetesMapping(in *v1beta1.KubernetesMapping, out *KubernetesMapping, s conversion.Scope) error {
	return autoConvert_v1beta1_KubernetesMapping_To_v1alpha3_KubernetesMapping(in, out, s)
}

func autoConvert_v1alpha3_OIDCProviderStatus_To_v1beta1_OIDCProviderStatus(in *OIDCProviderStatus, out *v1beta1.OIDCProviderStatus, s conversion.Scope) error {
	out.ARN = in.ARN
	out.TrustPolicy = in.TrustPolicy
	return nil
}

// Convert_v1alpha3_OIDCProviderStatus_To_v1beta1_OIDCProviderStatus is an autogenerated conversion function.
func Convert_v1alpha3_OIDCProviderStatus_To_v1beta1_OIDCProviderStatus(in *OIDCProviderStatus, out *v1beta1.OIDCProviderStatus, s conversion.Scope) error {
	return autoConvert_v1alpha3_OIDCProviderStatus_To_v1beta1_OIDCProviderStatus(in, out, s)
}

func autoConvert_v1beta1_OIDCProviderStatus_To_v1alpha3_OIDCProviderStatus(in *v1beta1.OIDCProviderStatus, out *OIDCProviderStatus, s conversion.Scope) error {
	out.ARN = in.ARN
	out.TrustPolicy = in.TrustPolicy
	return nil
}

// Convert_v1beta1_OIDCProviderStatus_To_v1alpha3_OIDCProviderStatus is an autogenerated conversion function.
func Convert_v1beta1_OIDCProviderStatus_To_v1alpha3_OIDCProviderStatus(in *v1beta1.OIDCProviderStatus, out *OIDCProviderStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_OIDCProviderStatus_To_v1alpha3_OIDCProviderStatus(in, out, s)
}

func autoConvert_v1alpha3_RoleMapping_To_v1beta1_RoleMapping(in *RoleMapping, out *v1beta1.RoleMapping, s conversion.Scope) error {
	out.RoleARN = in.RoleARN
	if err := Convert_v1alpha3_KubernetesMapping_To_v1beta1_KubernetesMapping(&in.KubernetesMapping, &out.KubernetesMapping, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_RoleMapping_To_v1beta1_RoleMapping is an autogenerated conversion function.
func Convert_v1alpha3_RoleMapping_To_v1beta1_RoleMapping(in *RoleMapping, out *v1beta1.RoleMapping, s conversion.Scope) error {
	return autoConvert_v1alpha3_RoleMapping_To_v1beta1_RoleMapping(in, out, s)
}

func autoConvert_v1beta1_RoleMapping_To_v1alpha3_RoleMapping(in *v1beta1.RoleMapping, out *RoleMapping, s conversion.Scope) error {
	out.RoleARN = in.RoleARN
	if err := Convert_v1beta1_KubernetesMapping_To_v1alpha3_KubernetesMapping(&in.KubernetesMapping, &out.KubernetesMapping, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_RoleMapping_To_v1alpha3_RoleMapping is an autogenerated conversion function.
func Convert_v1beta1_RoleMapping_To_v1alpha3_RoleMapping(in *v1beta1.RoleMapping, out *RoleMapping, s conversion.Scope) error {
	return autoConvert_v1beta1_RoleMapping_To_v1alpha3_RoleMapping(in, out, s)
}

func autoConvert_v1alpha3_UserMapping_To_v1beta1_UserMapping(in *UserMapping, out *v1beta1.UserMapping, s conversion.Scope) error {
	out.UserARN = in.UserARN
	if err := Convert_v1alpha3_KubernetesMapping_To_v1beta1_KubernetesMapping(&in.KubernetesMapping, &out.KubernetesMapping, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_UserMapping_To_v1beta1_UserMapping is an autogenerated conversion function.
func Convert_v1alpha3_UserMapping_To_v1beta1_UserMapping(in *UserMapping, out *v1beta1.UserMapping, s conversion.Scope) error {
	return autoConvert_v1alpha3_UserMapping_To_v1beta1_UserMapping(in, out, s)
}

func autoConvert_v1beta1_UserMapping_To_v1alpha3_UserMapping(in *v1beta1.UserMapping, out *UserMapping, s conversion.Scope) error {
	out.UserARN = in.UserARN
	if err := Convert_v1beta1_KubernetesMapping_To_v1alpha3_KubernetesMapping(&in.KubernetesMapping, &out.KubernetesMapping, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_UserMapping_To_v1alpha3_UserMapping is an autogenerated conversion function.
func Convert_v1beta1_UserMapping_To_v1alpha3_UserMapping(in *v1beta1.UserMapping, out *UserMapping, s conversion.Scope) error {
	return autoConvert_v1beta1_UserMapping_To_v1alpha3_UserMapping(in, out, s)
}
