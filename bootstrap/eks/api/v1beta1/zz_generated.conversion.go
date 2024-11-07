//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

package v1beta1

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1beta2 "sigs.k8s.io/cluster-api-provider-aws/v2/bootstrap/eks/api/v1beta2"
	apiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*EKSConfig)(nil), (*v1beta2.EKSConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EKSConfig_To_v1beta2_EKSConfig(a.(*EKSConfig), b.(*v1beta2.EKSConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.EKSConfig)(nil), (*EKSConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_EKSConfig_To_v1beta1_EKSConfig(a.(*v1beta2.EKSConfig), b.(*EKSConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EKSConfigList)(nil), (*v1beta2.EKSConfigList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EKSConfigList_To_v1beta2_EKSConfigList(a.(*EKSConfigList), b.(*v1beta2.EKSConfigList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.EKSConfigList)(nil), (*EKSConfigList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_EKSConfigList_To_v1beta1_EKSConfigList(a.(*v1beta2.EKSConfigList), b.(*EKSConfigList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EKSConfigSpec)(nil), (*v1beta2.EKSConfigSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EKSConfigSpec_To_v1beta2_EKSConfigSpec(a.(*EKSConfigSpec), b.(*v1beta2.EKSConfigSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EKSConfigStatus)(nil), (*v1beta2.EKSConfigStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EKSConfigStatus_To_v1beta2_EKSConfigStatus(a.(*EKSConfigStatus), b.(*v1beta2.EKSConfigStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.EKSConfigStatus)(nil), (*EKSConfigStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_EKSConfigStatus_To_v1beta1_EKSConfigStatus(a.(*v1beta2.EKSConfigStatus), b.(*EKSConfigStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EKSConfigTemplate)(nil), (*v1beta2.EKSConfigTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EKSConfigTemplate_To_v1beta2_EKSConfigTemplate(a.(*EKSConfigTemplate), b.(*v1beta2.EKSConfigTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.EKSConfigTemplate)(nil), (*EKSConfigTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_EKSConfigTemplate_To_v1beta1_EKSConfigTemplate(a.(*v1beta2.EKSConfigTemplate), b.(*EKSConfigTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EKSConfigTemplateList)(nil), (*v1beta2.EKSConfigTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EKSConfigTemplateList_To_v1beta2_EKSConfigTemplateList(a.(*EKSConfigTemplateList), b.(*v1beta2.EKSConfigTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.EKSConfigTemplateList)(nil), (*EKSConfigTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_EKSConfigTemplateList_To_v1beta1_EKSConfigTemplateList(a.(*v1beta2.EKSConfigTemplateList), b.(*EKSConfigTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EKSConfigTemplateResource)(nil), (*v1beta2.EKSConfigTemplateResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EKSConfigTemplateResource_To_v1beta2_EKSConfigTemplateResource(a.(*EKSConfigTemplateResource), b.(*v1beta2.EKSConfigTemplateResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.EKSConfigTemplateResource)(nil), (*EKSConfigTemplateResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_EKSConfigTemplateResource_To_v1beta1_EKSConfigTemplateResource(a.(*v1beta2.EKSConfigTemplateResource), b.(*EKSConfigTemplateResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EKSConfigTemplateSpec)(nil), (*v1beta2.EKSConfigTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EKSConfigTemplateSpec_To_v1beta2_EKSConfigTemplateSpec(a.(*EKSConfigTemplateSpec), b.(*v1beta2.EKSConfigTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.EKSConfigTemplateSpec)(nil), (*EKSConfigTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_EKSConfigTemplateSpec_To_v1beta1_EKSConfigTemplateSpec(a.(*v1beta2.EKSConfigTemplateSpec), b.(*EKSConfigTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PauseContainer)(nil), (*v1beta2.PauseContainer)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_PauseContainer_To_v1beta2_PauseContainer(a.(*PauseContainer), b.(*v1beta2.PauseContainer), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.PauseContainer)(nil), (*PauseContainer)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_PauseContainer_To_v1beta1_PauseContainer(a.(*v1beta2.PauseContainer), b.(*PauseContainer), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta2.EKSConfigSpec)(nil), (*EKSConfigSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_EKSConfigSpec_To_v1beta1_EKSConfigSpec(a.(*v1beta2.EKSConfigSpec), b.(*EKSConfigSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_EKSConfig_To_v1beta2_EKSConfig(in *EKSConfig, out *v1beta2.EKSConfig, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_EKSConfigSpec_To_v1beta2_EKSConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_EKSConfigStatus_To_v1beta2_EKSConfigStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_EKSConfig_To_v1beta2_EKSConfig is an autogenerated conversion function.
func Convert_v1beta1_EKSConfig_To_v1beta2_EKSConfig(in *EKSConfig, out *v1beta2.EKSConfig, s conversion.Scope) error {
	return autoConvert_v1beta1_EKSConfig_To_v1beta2_EKSConfig(in, out, s)
}

func autoConvert_v1beta2_EKSConfig_To_v1beta1_EKSConfig(in *v1beta2.EKSConfig, out *EKSConfig, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta2_EKSConfigSpec_To_v1beta1_EKSConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta2_EKSConfigStatus_To_v1beta1_EKSConfigStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta2_EKSConfig_To_v1beta1_EKSConfig is an autogenerated conversion function.
func Convert_v1beta2_EKSConfig_To_v1beta1_EKSConfig(in *v1beta2.EKSConfig, out *EKSConfig, s conversion.Scope) error {
	return autoConvert_v1beta2_EKSConfig_To_v1beta1_EKSConfig(in, out, s)
}

func autoConvert_v1beta1_EKSConfigList_To_v1beta2_EKSConfigList(in *EKSConfigList, out *v1beta2.EKSConfigList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta2.EKSConfig, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_EKSConfig_To_v1beta2_EKSConfig(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_EKSConfigList_To_v1beta2_EKSConfigList is an autogenerated conversion function.
func Convert_v1beta1_EKSConfigList_To_v1beta2_EKSConfigList(in *EKSConfigList, out *v1beta2.EKSConfigList, s conversion.Scope) error {
	return autoConvert_v1beta1_EKSConfigList_To_v1beta2_EKSConfigList(in, out, s)
}

func autoConvert_v1beta2_EKSConfigList_To_v1beta1_EKSConfigList(in *v1beta2.EKSConfigList, out *EKSConfigList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EKSConfig, len(*in))
		for i := range *in {
			if err := Convert_v1beta2_EKSConfig_To_v1beta1_EKSConfig(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta2_EKSConfigList_To_v1beta1_EKSConfigList is an autogenerated conversion function.
func Convert_v1beta2_EKSConfigList_To_v1beta1_EKSConfigList(in *v1beta2.EKSConfigList, out *EKSConfigList, s conversion.Scope) error {
	return autoConvert_v1beta2_EKSConfigList_To_v1beta1_EKSConfigList(in, out, s)
}

func autoConvert_v1beta1_EKSConfigSpec_To_v1beta2_EKSConfigSpec(in *EKSConfigSpec, out *v1beta2.EKSConfigSpec, s conversion.Scope) error {
	out.KubeletExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.KubeletExtraArgs))
	out.ContainerRuntime = (*string)(unsafe.Pointer(in.ContainerRuntime))
	out.DNSClusterIP = (*string)(unsafe.Pointer(in.DNSClusterIP))
	out.DockerConfigJSON = (*string)(unsafe.Pointer(in.DockerConfigJSON))
	out.APIRetryAttempts = (*int)(unsafe.Pointer(in.APIRetryAttempts))
	out.PauseContainer = (*v1beta2.PauseContainer)(unsafe.Pointer(in.PauseContainer))
	out.UseMaxPods = (*bool)(unsafe.Pointer(in.UseMaxPods))
	out.ServiceIPV6Cidr = (*string)(unsafe.Pointer(in.ServiceIPV6Cidr))
	return nil
}

// Convert_v1beta1_EKSConfigSpec_To_v1beta2_EKSConfigSpec is an autogenerated conversion function.
func Convert_v1beta1_EKSConfigSpec_To_v1beta2_EKSConfigSpec(in *EKSConfigSpec, out *v1beta2.EKSConfigSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_EKSConfigSpec_To_v1beta2_EKSConfigSpec(in, out, s)
}

func autoConvert_v1beta2_EKSConfigSpec_To_v1beta1_EKSConfigSpec(in *v1beta2.EKSConfigSpec, out *EKSConfigSpec, s conversion.Scope) error {
	out.KubeletExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.KubeletExtraArgs))
	out.ContainerRuntime = (*string)(unsafe.Pointer(in.ContainerRuntime))
	out.DNSClusterIP = (*string)(unsafe.Pointer(in.DNSClusterIP))
	out.DockerConfigJSON = (*string)(unsafe.Pointer(in.DockerConfigJSON))
	out.APIRetryAttempts = (*int)(unsafe.Pointer(in.APIRetryAttempts))
	out.PauseContainer = (*PauseContainer)(unsafe.Pointer(in.PauseContainer))
	out.UseMaxPods = (*bool)(unsafe.Pointer(in.UseMaxPods))
	out.ServiceIPV6Cidr = (*string)(unsafe.Pointer(in.ServiceIPV6Cidr))
	// WARNING: in.PreBootstrapCommands requires manual conversion: does not exist in peer-type
	// WARNING: in.PostBootstrapCommands requires manual conversion: does not exist in peer-type
	// WARNING: in.BootstrapCommandOverride requires manual conversion: does not exist in peer-type
	// WARNING: in.Files requires manual conversion: does not exist in peer-type
	// WARNING: in.DiskSetup requires manual conversion: does not exist in peer-type
	// WARNING: in.Mounts requires manual conversion: does not exist in peer-type
	// WARNING: in.Users requires manual conversion: does not exist in peer-type
	// WARNING: in.NTP requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1beta1_EKSConfigStatus_To_v1beta2_EKSConfigStatus(in *EKSConfigStatus, out *v1beta2.EKSConfigStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	out.DataSecretName = (*string)(unsafe.Pointer(in.DataSecretName))
	out.FailureReason = in.FailureReason
	out.FailureMessage = in.FailureMessage
	out.ObservedGeneration = in.ObservedGeneration
	out.Conditions = *(*apiv1beta1.Conditions)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v1beta1_EKSConfigStatus_To_v1beta2_EKSConfigStatus is an autogenerated conversion function.
func Convert_v1beta1_EKSConfigStatus_To_v1beta2_EKSConfigStatus(in *EKSConfigStatus, out *v1beta2.EKSConfigStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_EKSConfigStatus_To_v1beta2_EKSConfigStatus(in, out, s)
}

func autoConvert_v1beta2_EKSConfigStatus_To_v1beta1_EKSConfigStatus(in *v1beta2.EKSConfigStatus, out *EKSConfigStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	out.DataSecretName = (*string)(unsafe.Pointer(in.DataSecretName))
	out.FailureReason = in.FailureReason
	out.FailureMessage = in.FailureMessage
	out.ObservedGeneration = in.ObservedGeneration
	out.Conditions = *(*apiv1beta1.Conditions)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v1beta2_EKSConfigStatus_To_v1beta1_EKSConfigStatus is an autogenerated conversion function.
func Convert_v1beta2_EKSConfigStatus_To_v1beta1_EKSConfigStatus(in *v1beta2.EKSConfigStatus, out *EKSConfigStatus, s conversion.Scope) error {
	return autoConvert_v1beta2_EKSConfigStatus_To_v1beta1_EKSConfigStatus(in, out, s)
}

func autoConvert_v1beta1_EKSConfigTemplate_To_v1beta2_EKSConfigTemplate(in *EKSConfigTemplate, out *v1beta2.EKSConfigTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_EKSConfigTemplateSpec_To_v1beta2_EKSConfigTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_EKSConfigTemplate_To_v1beta2_EKSConfigTemplate is an autogenerated conversion function.
func Convert_v1beta1_EKSConfigTemplate_To_v1beta2_EKSConfigTemplate(in *EKSConfigTemplate, out *v1beta2.EKSConfigTemplate, s conversion.Scope) error {
	return autoConvert_v1beta1_EKSConfigTemplate_To_v1beta2_EKSConfigTemplate(in, out, s)
}

func autoConvert_v1beta2_EKSConfigTemplate_To_v1beta1_EKSConfigTemplate(in *v1beta2.EKSConfigTemplate, out *EKSConfigTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta2_EKSConfigTemplateSpec_To_v1beta1_EKSConfigTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta2_EKSConfigTemplate_To_v1beta1_EKSConfigTemplate is an autogenerated conversion function.
func Convert_v1beta2_EKSConfigTemplate_To_v1beta1_EKSConfigTemplate(in *v1beta2.EKSConfigTemplate, out *EKSConfigTemplate, s conversion.Scope) error {
	return autoConvert_v1beta2_EKSConfigTemplate_To_v1beta1_EKSConfigTemplate(in, out, s)
}

func autoConvert_v1beta1_EKSConfigTemplateList_To_v1beta2_EKSConfigTemplateList(in *EKSConfigTemplateList, out *v1beta2.EKSConfigTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta2.EKSConfigTemplate, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_EKSConfigTemplate_To_v1beta2_EKSConfigTemplate(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_EKSConfigTemplateList_To_v1beta2_EKSConfigTemplateList is an autogenerated conversion function.
func Convert_v1beta1_EKSConfigTemplateList_To_v1beta2_EKSConfigTemplateList(in *EKSConfigTemplateList, out *v1beta2.EKSConfigTemplateList, s conversion.Scope) error {
	return autoConvert_v1beta1_EKSConfigTemplateList_To_v1beta2_EKSConfigTemplateList(in, out, s)
}

func autoConvert_v1beta2_EKSConfigTemplateList_To_v1beta1_EKSConfigTemplateList(in *v1beta2.EKSConfigTemplateList, out *EKSConfigTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EKSConfigTemplate, len(*in))
		for i := range *in {
			if err := Convert_v1beta2_EKSConfigTemplate_To_v1beta1_EKSConfigTemplate(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta2_EKSConfigTemplateList_To_v1beta1_EKSConfigTemplateList is an autogenerated conversion function.
func Convert_v1beta2_EKSConfigTemplateList_To_v1beta1_EKSConfigTemplateList(in *v1beta2.EKSConfigTemplateList, out *EKSConfigTemplateList, s conversion.Scope) error {
	return autoConvert_v1beta2_EKSConfigTemplateList_To_v1beta1_EKSConfigTemplateList(in, out, s)
}

func autoConvert_v1beta1_EKSConfigTemplateResource_To_v1beta2_EKSConfigTemplateResource(in *EKSConfigTemplateResource, out *v1beta2.EKSConfigTemplateResource, s conversion.Scope) error {
	if err := Convert_v1beta1_EKSConfigSpec_To_v1beta2_EKSConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_EKSConfigTemplateResource_To_v1beta2_EKSConfigTemplateResource is an autogenerated conversion function.
func Convert_v1beta1_EKSConfigTemplateResource_To_v1beta2_EKSConfigTemplateResource(in *EKSConfigTemplateResource, out *v1beta2.EKSConfigTemplateResource, s conversion.Scope) error {
	return autoConvert_v1beta1_EKSConfigTemplateResource_To_v1beta2_EKSConfigTemplateResource(in, out, s)
}

func autoConvert_v1beta2_EKSConfigTemplateResource_To_v1beta1_EKSConfigTemplateResource(in *v1beta2.EKSConfigTemplateResource, out *EKSConfigTemplateResource, s conversion.Scope) error {
	if err := Convert_v1beta2_EKSConfigSpec_To_v1beta1_EKSConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta2_EKSConfigTemplateResource_To_v1beta1_EKSConfigTemplateResource is an autogenerated conversion function.
func Convert_v1beta2_EKSConfigTemplateResource_To_v1beta1_EKSConfigTemplateResource(in *v1beta2.EKSConfigTemplateResource, out *EKSConfigTemplateResource, s conversion.Scope) error {
	return autoConvert_v1beta2_EKSConfigTemplateResource_To_v1beta1_EKSConfigTemplateResource(in, out, s)
}

func autoConvert_v1beta1_EKSConfigTemplateSpec_To_v1beta2_EKSConfigTemplateSpec(in *EKSConfigTemplateSpec, out *v1beta2.EKSConfigTemplateSpec, s conversion.Scope) error {
	if err := Convert_v1beta1_EKSConfigTemplateResource_To_v1beta2_EKSConfigTemplateResource(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_EKSConfigTemplateSpec_To_v1beta2_EKSConfigTemplateSpec is an autogenerated conversion function.
func Convert_v1beta1_EKSConfigTemplateSpec_To_v1beta2_EKSConfigTemplateSpec(in *EKSConfigTemplateSpec, out *v1beta2.EKSConfigTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_EKSConfigTemplateSpec_To_v1beta2_EKSConfigTemplateSpec(in, out, s)
}

func autoConvert_v1beta2_EKSConfigTemplateSpec_To_v1beta1_EKSConfigTemplateSpec(in *v1beta2.EKSConfigTemplateSpec, out *EKSConfigTemplateSpec, s conversion.Scope) error {
	if err := Convert_v1beta2_EKSConfigTemplateResource_To_v1beta1_EKSConfigTemplateResource(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta2_EKSConfigTemplateSpec_To_v1beta1_EKSConfigTemplateSpec is an autogenerated conversion function.
func Convert_v1beta2_EKSConfigTemplateSpec_To_v1beta1_EKSConfigTemplateSpec(in *v1beta2.EKSConfigTemplateSpec, out *EKSConfigTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1beta2_EKSConfigTemplateSpec_To_v1beta1_EKSConfigTemplateSpec(in, out, s)
}

func autoConvert_v1beta1_PauseContainer_To_v1beta2_PauseContainer(in *PauseContainer, out *v1beta2.PauseContainer, s conversion.Scope) error {
	out.AccountNumber = in.AccountNumber
	out.Version = in.Version
	return nil
}

// Convert_v1beta1_PauseContainer_To_v1beta2_PauseContainer is an autogenerated conversion function.
func Convert_v1beta1_PauseContainer_To_v1beta2_PauseContainer(in *PauseContainer, out *v1beta2.PauseContainer, s conversion.Scope) error {
	return autoConvert_v1beta1_PauseContainer_To_v1beta2_PauseContainer(in, out, s)
}

func autoConvert_v1beta2_PauseContainer_To_v1beta1_PauseContainer(in *v1beta2.PauseContainer, out *PauseContainer, s conversion.Scope) error {
	out.AccountNumber = in.AccountNumber
	out.Version = in.Version
	return nil
}

// Convert_v1beta2_PauseContainer_To_v1beta1_PauseContainer is an autogenerated conversion function.
func Convert_v1beta2_PauseContainer_To_v1beta1_PauseContainer(in *v1beta2.PauseContainer, out *PauseContainer, s conversion.Scope) error {
	return autoConvert_v1beta2_PauseContainer_To_v1beta1_PauseContainer(in, out, s)
}
