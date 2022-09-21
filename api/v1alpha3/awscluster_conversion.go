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
	"fmt"

	apiconversion "k8s.io/apimachinery/pkg/conversion"
	clusterv1alpha3 "sigs.k8s.io/cluster-api/api/v1alpha3"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	utilconversion "sigs.k8s.io/cluster-api/util/conversion"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	infrav1 "sigs.k8s.io/cluster-api-provider-aws/api/v1beta1"
)

// ConvertTo converts the v1alpha3 AWSCluster receiver to a v1beta1 AWSCluster.
func (r *AWSCluster) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infrav1.AWSCluster)

	if err := Convert_v1alpha3_AWSCluster_To_v1beta1_AWSCluster(r, dst, nil); err != nil {
		return err
	}
	// Manually restore data.
	restored := &infrav1.AWSCluster{}
	if ok, err := utilconversion.UnmarshalData(r, restored); err != nil || !ok {
		return err
	}

	if restored.Status.Bastion != nil {
		if dst.Status.Bastion == nil {
			dst.Status.Bastion = &infrav1.Instance{}
		}
		restoreInstance(restored.Status.Bastion, dst.Status.Bastion)
	}

	if restored.Spec.ControlPlaneLoadBalancer != nil {
		if dst.Spec.ControlPlaneLoadBalancer == nil {
			dst.Spec.ControlPlaneLoadBalancer = &infrav1.AWSLoadBalancerSpec{}
		}
		restoreControlPlaneLoadBalancer(restored.Spec.ControlPlaneLoadBalancer, dst.Spec.ControlPlaneLoadBalancer)
	}

	dst.Spec.S3Bucket = restored.Spec.S3Bucket
	if restored.Spec.NetworkSpec.VPC.IPv6 != nil {
		if dst.Spec.NetworkSpec.VPC.IPv6 == nil {
			dst.Spec.NetworkSpec.VPC.IPv6 = &infrav1.IPv6{}
		}
		dst.Spec.NetworkSpec.VPC.IPv6.EgressOnlyInternetGatewayID = restored.Spec.NetworkSpec.VPC.IPv6.EgressOnlyInternetGatewayID
		dst.Spec.NetworkSpec.VPC.IPv6.CidrBlock = restored.Spec.NetworkSpec.VPC.IPv6.CidrBlock
		dst.Spec.NetworkSpec.VPC.IPv6.PoolID = restored.Spec.NetworkSpec.VPC.IPv6.PoolID
	}

	for i := range dst.Spec.NetworkSpec.Subnets {
		var found bool
		for k := range restored.Spec.NetworkSpec.Subnets {
			if dst.Spec.NetworkSpec.Subnets[i].ID == restored.Spec.NetworkSpec.Subnets[k].ID {
				dst.Spec.NetworkSpec.Subnets[i].IsIPv6 = restored.Spec.NetworkSpec.Subnets[i].IsIPv6
				dst.Spec.NetworkSpec.Subnets[i].IPv6CidrBlock = restored.Spec.NetworkSpec.Subnets[i].IPv6CidrBlock
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("subnet with id %s not found amongst restored subnets", dst.Spec.NetworkSpec.Subnets[i].ID)
		}
	}

	for k, v := range restored.Status.Network.SecurityGroups {
		if dst.Status.Network.SecurityGroups == nil {
			dst.Status.Network.SecurityGroups = make(map[infrav1.SecurityGroupRole]infrav1.SecurityGroup)
		}
		// This will overwrite the whole rule set... But maybe that's fine?
		dst.Status.Network.SecurityGroups[k] = v
	}

	return nil
}

// restoreControlPlaneLoadBalancer manually restores the control plane loadbalancer data.
// Assumes restored and dst are non-nil.
func restoreControlPlaneLoadBalancer(restored, dst *infrav1.AWSLoadBalancerSpec) {
	dst.Name = restored.Name
	dst.HealthCheckProtocol = restored.HealthCheckProtocol
}

// ConvertFrom converts the v1beta1 AWSCluster receiver to a v1alpha3 AWSCluster.
func (r *AWSCluster) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infrav1.AWSCluster)

	if err := Convert_v1beta1_AWSCluster_To_v1alpha3_AWSCluster(src, r, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion.
	if err := utilconversion.MarshalData(src, r); err != nil {
		return err
	}

	return nil
}

// ConvertTo converts the v1alpha3 AWSClusterList receiver to a v1beta1 AWSClusterList.
func (r *AWSClusterList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infrav1.AWSClusterList)

	return Convert_v1alpha3_AWSClusterList_To_v1beta1_AWSClusterList(r, dst, nil)
}

// ConvertFrom converts the v1beta1 AWSClusterList receiver to a v1alpha3 AWSClusterList.
func (r *AWSClusterList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infrav1.AWSClusterList)

	return Convert_v1beta1_AWSClusterList_To_v1alpha3_AWSClusterList(src, r, nil)
}

// Convert_v1alpha3_APIEndpoint_To_v1beta1_APIEndpoint .
func Convert_v1alpha3_APIEndpoint_To_v1beta1_APIEndpoint(in *clusterv1alpha3.APIEndpoint, out *clusterv1.APIEndpoint, s apiconversion.Scope) error {
	return clusterv1alpha3.Convert_v1alpha3_APIEndpoint_To_v1beta1_APIEndpoint(in, out, s)
}

// Convert_v1beta1_APIEndpoint_To_v1alpha3_APIEndpoint .
func Convert_v1beta1_APIEndpoint_To_v1alpha3_APIEndpoint(in *clusterv1.APIEndpoint, out *clusterv1alpha3.APIEndpoint, s apiconversion.Scope) error {
	return clusterv1alpha3.Convert_v1beta1_APIEndpoint_To_v1alpha3_APIEndpoint(in, out, s)
}

// Convert_v1alpha3_Network_To_v1alpha4_NetworkStatus is based on the autogenerated function and handles the renaming of the Network struct to NetworkStatus
func Convert_v1alpha3_Network_To_v1beta1_NetworkStatus(in *Network, out *infrav1.NetworkStatus, s apiconversion.Scope) error {
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make(map[infrav1.SecurityGroupRole]infrav1.SecurityGroup, len(*in))
		for key, val := range *in {
			newVal := new(infrav1.SecurityGroup)
			if err := Convert_v1alpha3_SecurityGroup_To_v1beta1_SecurityGroup(&val, newVal, s); err != nil {
				return err
			}
			(*out)[infrav1.SecurityGroupRole(key)] = *newVal
		}
	} else {
		out.SecurityGroups = nil
	}
	if err := Convert_v1alpha3_ClassicELB_To_v1beta1_ClassicELB(&in.APIServerELB, &out.APIServerELB, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha4_NetworkStatus_To_v1alpha3_Network is based on the autogenerated function and handles the renaming of the NetworkStatus struct to Network
func Convert_v1beta1_NetworkStatus_To_v1alpha3_Network(in *infrav1.NetworkStatus, out *Network, s apiconversion.Scope) error {
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make(map[SecurityGroupRole]SecurityGroup, len(*in))
		for key, val := range *in {
			newVal := new(SecurityGroup)
			if err := Convert_v1beta1_SecurityGroup_To_v1alpha3_SecurityGroup(&val, newVal, s); err != nil {
				return err
			}
			(*out)[SecurityGroupRole(key)] = *newVal
		}
	} else {
		out.SecurityGroups = nil
	}
	if err := Convert_v1beta1_ClassicELB_To_v1alpha3_ClassicELB(&in.APIServerELB, &out.APIServerELB, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1beta1_AWSLoadBalancerSpec_To_v1alpha3_AWSLoadBalancerSpec(in *infrav1.AWSLoadBalancerSpec, out *AWSLoadBalancerSpec, s apiconversion.Scope) error {
	return autoConvert_v1beta1_AWSLoadBalancerSpec_To_v1alpha3_AWSLoadBalancerSpec(in, out, s)
}

func Convert_v1beta1_AWSClusterSpec_To_v1alpha3_AWSClusterSpec(in *infrav1.AWSClusterSpec, out *AWSClusterSpec, s apiconversion.Scope) error {
	return autoConvert_v1beta1_AWSClusterSpec_To_v1alpha3_AWSClusterSpec(in, out, s)
}

func Convert_v1beta1_VPCSpec_To_v1alpha3_VPCSpec(in *infrav1.VPCSpec, out *VPCSpec, s apiconversion.Scope) error {
	return autoConvert_v1beta1_VPCSpec_To_v1alpha3_VPCSpec(in, out, s)
}

func Convert_v1beta1_SubnetSpec_To_v1alpha3_SubnetSpec(in *infrav1.SubnetSpec, out *SubnetSpec, s apiconversion.Scope) error {
	return autoConvert_v1beta1_SubnetSpec_To_v1alpha3_SubnetSpec(in, out, s)
}

func Convert_v1beta1_IngressRule_To_v1alpha3_IngressRule(in *infrav1.IngressRule, out *IngressRule, s apiconversion.Scope) error {
	return autoConvert_v1beta1_IngressRule_To_v1alpha3_IngressRule(in, out, s)
}
