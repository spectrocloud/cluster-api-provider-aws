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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha3

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1alpha3 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha3"
	cluster_apiapiv1alpha3 "sigs.k8s.io/cluster-api/api/v1alpha3"
	"sigs.k8s.io/cluster-api/errors"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSFargateProfile) DeepCopyInto(out *AWSFargateProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSFargateProfile.
func (in *AWSFargateProfile) DeepCopy() *AWSFargateProfile {
	if in == nil {
		return nil
	}
	out := new(AWSFargateProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSFargateProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSFargateProfileList) DeepCopyInto(out *AWSFargateProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSFargateProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSFargateProfileList.
func (in *AWSFargateProfileList) DeepCopy() *AWSFargateProfileList {
	if in == nil {
		return nil
	}
	out := new(AWSFargateProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSFargateProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSLaunchTemplate) DeepCopyInto(out *AWSLaunchTemplate) {
	*out = *in
	in.AMI.DeepCopyInto(&out.AMI)
	if in.RootVolume != nil {
		in, out := &in.RootVolume, &out.RootVolume
		*out = new(apiv1alpha3.Volume)
		**out = **in
	}
	if in.SSHKeyName != nil {
		in, out := &in.SSHKeyName, &out.SSHKeyName
		*out = new(string)
		**out = **in
	}
	if in.VersionNumber != nil {
		in, out := &in.VersionNumber, &out.VersionNumber
		*out = new(int64)
		**out = **in
	}
	if in.AdditionalSecurityGroups != nil {
		in, out := &in.AdditionalSecurityGroups, &out.AdditionalSecurityGroups
		*out = make([]apiv1alpha3.AWSResourceReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSLaunchTemplate.
func (in *AWSLaunchTemplate) DeepCopy() *AWSLaunchTemplate {
	if in == nil {
		return nil
	}
	out := new(AWSLaunchTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachinePool) DeepCopyInto(out *AWSMachinePool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachinePool.
func (in *AWSMachinePool) DeepCopy() *AWSMachinePool {
	if in == nil {
		return nil
	}
	out := new(AWSMachinePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSMachinePool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachinePoolInstanceStatus) DeepCopyInto(out *AWSMachinePoolInstanceStatus) {
	*out = *in
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachinePoolInstanceStatus.
func (in *AWSMachinePoolInstanceStatus) DeepCopy() *AWSMachinePoolInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(AWSMachinePoolInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachinePoolList) DeepCopyInto(out *AWSMachinePoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSMachinePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachinePoolList.
func (in *AWSMachinePoolList) DeepCopy() *AWSMachinePoolList {
	if in == nil {
		return nil
	}
	out := new(AWSMachinePoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSMachinePoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachinePoolSpec) DeepCopyInto(out *AWSMachinePoolSpec) {
	*out = *in
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]apiv1alpha3.AWSResourceReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdditionalTags != nil {
		in, out := &in.AdditionalTags, &out.AdditionalTags
		*out = make(apiv1alpha3.Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.AWSLaunchTemplate.DeepCopyInto(&out.AWSLaunchTemplate)
	if in.MixedInstancesPolicy != nil {
		in, out := &in.MixedInstancesPolicy, &out.MixedInstancesPolicy
		*out = new(MixedInstancesPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.ProviderIDList != nil {
		in, out := &in.ProviderIDList, &out.ProviderIDList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.DefaultCoolDown = in.DefaultCoolDown
	if in.RefreshPreferences != nil {
		in, out := &in.RefreshPreferences, &out.RefreshPreferences
		*out = new(RefreshPreferences)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachinePoolSpec.
func (in *AWSMachinePoolSpec) DeepCopy() *AWSMachinePoolSpec {
	if in == nil {
		return nil
	}
	out := new(AWSMachinePoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachinePoolStatus) DeepCopyInto(out *AWSMachinePoolStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(cluster_apiapiv1alpha3.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = make([]AWSMachinePoolInstanceStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(errors.MachineStatusError)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
	if in.ASGStatus != nil {
		in, out := &in.ASGStatus, &out.ASGStatus
		*out = new(ASGStatus)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachinePoolStatus.
func (in *AWSMachinePoolStatus) DeepCopy() *AWSMachinePoolStatus {
	if in == nil {
		return nil
	}
	out := new(AWSMachinePoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSManagedCluster) DeepCopyInto(out *AWSManagedCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSManagedCluster.
func (in *AWSManagedCluster) DeepCopy() *AWSManagedCluster {
	if in == nil {
		return nil
	}
	out := new(AWSManagedCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSManagedCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSManagedClusterList) DeepCopyInto(out *AWSManagedClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSManagedCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSManagedClusterList.
func (in *AWSManagedClusterList) DeepCopy() *AWSManagedClusterList {
	if in == nil {
		return nil
	}
	out := new(AWSManagedClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSManagedClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSManagedClusterSpec) DeepCopyInto(out *AWSManagedClusterSpec) {
	*out = *in
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSManagedClusterSpec.
func (in *AWSManagedClusterSpec) DeepCopy() *AWSManagedClusterSpec {
	if in == nil {
		return nil
	}
	out := new(AWSManagedClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSManagedClusterStatus) DeepCopyInto(out *AWSManagedClusterStatus) {
	*out = *in
	if in.FailureDomains != nil {
		in, out := &in.FailureDomains, &out.FailureDomains
		*out = make(cluster_apiapiv1alpha3.FailureDomains, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSManagedClusterStatus.
func (in *AWSManagedClusterStatus) DeepCopy() *AWSManagedClusterStatus {
	if in == nil {
		return nil
	}
	out := new(AWSManagedClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSManagedMachinePool) DeepCopyInto(out *AWSManagedMachinePool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSManagedMachinePool.
func (in *AWSManagedMachinePool) DeepCopy() *AWSManagedMachinePool {
	if in == nil {
		return nil
	}
	out := new(AWSManagedMachinePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSManagedMachinePool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSManagedMachinePoolList) DeepCopyInto(out *AWSManagedMachinePoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSManagedMachinePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSManagedMachinePoolList.
func (in *AWSManagedMachinePoolList) DeepCopy() *AWSManagedMachinePoolList {
	if in == nil {
		return nil
	}
	out := new(AWSManagedMachinePoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSManagedMachinePoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSManagedMachinePoolSpec) DeepCopyInto(out *AWSManagedMachinePoolSpec) {
	*out = *in
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubnetIDs != nil {
		in, out := &in.SubnetIDs, &out.SubnetIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AdditionalTags != nil {
		in, out := &in.AdditionalTags, &out.AdditionalTags
		*out = make(apiv1alpha3.Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AMIVersion != nil {
		in, out := &in.AMIVersion, &out.AMIVersion
		*out = new(string)
		**out = **in
	}
	if in.AMIType != nil {
		in, out := &in.AMIType, &out.AMIType
		*out = new(ManagedMachineAMIType)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DiskSize != nil {
		in, out := &in.DiskSize, &out.DiskSize
		*out = new(int32)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.Scaling != nil {
		in, out := &in.Scaling, &out.Scaling
		*out = new(ManagedMachinePoolScaling)
		(*in).DeepCopyInto(*out)
	}
	if in.RemoteAccess != nil {
		in, out := &in.RemoteAccess, &out.RemoteAccess
		*out = new(ManagedRemoteAccess)
		(*in).DeepCopyInto(*out)
	}
	if in.ProviderIDList != nil {
		in, out := &in.ProviderIDList, &out.ProviderIDList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CapacityType != nil {
		in, out := &in.CapacityType, &out.CapacityType
		*out = new(ManagedMachinePoolCapacityType)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSManagedMachinePoolSpec.
func (in *AWSManagedMachinePoolSpec) DeepCopy() *AWSManagedMachinePoolSpec {
	if in == nil {
		return nil
	}
	out := new(AWSManagedMachinePoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSManagedMachinePoolStatus) DeepCopyInto(out *AWSManagedMachinePoolStatus) {
	*out = *in
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(errors.MachineStatusError)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(cluster_apiapiv1alpha3.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSManagedMachinePoolStatus.
func (in *AWSManagedMachinePoolStatus) DeepCopy() *AWSManagedMachinePoolStatus {
	if in == nil {
		return nil
	}
	out := new(AWSManagedMachinePoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingGroup) DeepCopyInto(out *AutoScalingGroup) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(apiv1alpha3.Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DesiredCapacity != nil {
		in, out := &in.DesiredCapacity, &out.DesiredCapacity
		*out = new(int32)
		**out = **in
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.DefaultCoolDown = in.DefaultCoolDown
	if in.MixedInstancesPolicy != nil {
		in, out := &in.MixedInstancesPolicy, &out.MixedInstancesPolicy
		*out = new(MixedInstancesPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = make([]apiv1alpha3.Instance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingGroup.
func (in *AutoScalingGroup) DeepCopy() *AutoScalingGroup {
	if in == nil {
		return nil
	}
	out := new(AutoScalingGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceMapping) DeepCopyInto(out *BlockDeviceMapping) {
	*out = *in
	out.Ebs = in.Ebs
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceMapping.
func (in *BlockDeviceMapping) DeepCopy() *BlockDeviceMapping {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EBS) DeepCopyInto(out *EBS) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EBS.
func (in *EBS) DeepCopy() *EBS {
	if in == nil {
		return nil
	}
	out := new(EBS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FargateProfileSpec) DeepCopyInto(out *FargateProfileSpec) {
	*out = *in
	if in.SubnetIDs != nil {
		in, out := &in.SubnetIDs, &out.SubnetIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AdditionalTags != nil {
		in, out := &in.AdditionalTags, &out.AdditionalTags
		*out = make(apiv1alpha3.Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Selectors != nil {
		in, out := &in.Selectors, &out.Selectors
		*out = make([]FargateSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FargateProfileSpec.
func (in *FargateProfileSpec) DeepCopy() *FargateProfileSpec {
	if in == nil {
		return nil
	}
	out := new(FargateProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FargateProfileStatus) DeepCopyInto(out *FargateProfileStatus) {
	*out = *in
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(errors.MachineStatusError)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(cluster_apiapiv1alpha3.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FargateProfileStatus.
func (in *FargateProfileStatus) DeepCopy() *FargateProfileStatus {
	if in == nil {
		return nil
	}
	out := new(FargateProfileStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FargateSelector) DeepCopyInto(out *FargateSelector) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FargateSelector.
func (in *FargateSelector) DeepCopy() *FargateSelector {
	if in == nil {
		return nil
	}
	out := new(FargateSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstancesDistribution) DeepCopyInto(out *InstancesDistribution) {
	*out = *in
	if in.OnDemandBaseCapacity != nil {
		in, out := &in.OnDemandBaseCapacity, &out.OnDemandBaseCapacity
		*out = new(int64)
		**out = **in
	}
	if in.OnDemandPercentageAboveBaseCapacity != nil {
		in, out := &in.OnDemandPercentageAboveBaseCapacity, &out.OnDemandPercentageAboveBaseCapacity
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstancesDistribution.
func (in *InstancesDistribution) DeepCopy() *InstancesDistribution {
	if in == nil {
		return nil
	}
	out := new(InstancesDistribution)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedMachinePoolScaling) DeepCopyInto(out *ManagedMachinePoolScaling) {
	*out = *in
	if in.MinSize != nil {
		in, out := &in.MinSize, &out.MinSize
		*out = new(int32)
		**out = **in
	}
	if in.MaxSize != nil {
		in, out := &in.MaxSize, &out.MaxSize
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedMachinePoolScaling.
func (in *ManagedMachinePoolScaling) DeepCopy() *ManagedMachinePoolScaling {
	if in == nil {
		return nil
	}
	out := new(ManagedMachinePoolScaling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedRemoteAccess) DeepCopyInto(out *ManagedRemoteAccess) {
	*out = *in
	if in.SSHKeyName != nil {
		in, out := &in.SSHKeyName, &out.SSHKeyName
		*out = new(string)
		**out = **in
	}
	if in.SourceSecurityGroups != nil {
		in, out := &in.SourceSecurityGroups, &out.SourceSecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedRemoteAccess.
func (in *ManagedRemoteAccess) DeepCopy() *ManagedRemoteAccess {
	if in == nil {
		return nil
	}
	out := new(ManagedRemoteAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MixedInstancesPolicy) DeepCopyInto(out *MixedInstancesPolicy) {
	*out = *in
	if in.InstancesDistribution != nil {
		in, out := &in.InstancesDistribution, &out.InstancesDistribution
		*out = new(InstancesDistribution)
		(*in).DeepCopyInto(*out)
	}
	if in.Overrides != nil {
		in, out := &in.Overrides, &out.Overrides
		*out = make([]Overrides, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MixedInstancesPolicy.
func (in *MixedInstancesPolicy) DeepCopy() *MixedInstancesPolicy {
	if in == nil {
		return nil
	}
	out := new(MixedInstancesPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Overrides) DeepCopyInto(out *Overrides) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Overrides.
func (in *Overrides) DeepCopy() *Overrides {
	if in == nil {
		return nil
	}
	out := new(Overrides)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RefreshPreferences) DeepCopyInto(out *RefreshPreferences) {
	*out = *in
	if in.Strategy != nil {
		in, out := &in.Strategy, &out.Strategy
		*out = new(string)
		**out = **in
	}
	if in.InstanceWarmup != nil {
		in, out := &in.InstanceWarmup, &out.InstanceWarmup
		*out = new(int64)
		**out = **in
	}
	if in.MinHealthyPercentage != nil {
		in, out := &in.MinHealthyPercentage, &out.MinHealthyPercentage
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RefreshPreferences.
func (in *RefreshPreferences) DeepCopy() *RefreshPreferences {
	if in == nil {
		return nil
	}
	out := new(RefreshPreferences)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Tags) DeepCopyInto(out *Tags) {
	{
		in := &in
		*out = make(Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tags.
func (in Tags) DeepCopy() Tags {
	if in == nil {
		return nil
	}
	out := new(Tags)
	in.DeepCopyInto(out)
	return *out
}
