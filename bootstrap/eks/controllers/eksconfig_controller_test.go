/*
Copyright 2020 The Kubernetes Authors.

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

package controllers

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	ekscontrolplanev1 "sigs.k8s.io/cluster-api-provider-aws/v2/controlplane/eks/api/v1beta2"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	bsutil "sigs.k8s.io/cluster-api/bootstrap/util"
)

func TestEKSConfigReconcilerReturnEarlyIfClusterInfraNotReady(t *testing.T) {
	g := NewWithT(t)

	cluster := newCluster("cluster")
	machine := newMachine(cluster, "machine")
	config := newEKSConfig(machine)

	cluster.Status = clusterv1.ClusterStatus{
		InfrastructureReady: false,
	}

	reconciler := EKSConfigReconciler{
		Client: testEnv.Client,
	}

	g.Eventually(func(gomega Gomega) {
		result, err := reconciler.joinWorker(context.Background(), cluster, config, configOwner("Machine"))
		gomega.Expect(err).NotTo(HaveOccurred())
		gomega.Expect(result.Requeue).To(BeFalse())
	}).Should(Succeed())
}

func TestEKSConfigReconcilerReturnEarlyIfClusterControlPlaneNotInitialized(t *testing.T) {
	g := NewWithT(t)

	cluster := newCluster("cluster")
	machine := newMachine(cluster, "machine")
	config := newEKSConfig(machine)

	cluster.Status = clusterv1.ClusterStatus{
		InfrastructureReady: true,
	}

	reconciler := EKSConfigReconciler{
		Client: testEnv.Client,
	}

	g.Eventually(func(gomega Gomega) {
		result, err := reconciler.joinWorker(context.Background(), cluster, config, configOwner("Machine"))
		gomega.Expect(err).NotTo(HaveOccurred())
		gomega.Expect(result.Requeue).To(BeFalse())
	}).Should(Succeed())
}

func TestDetermineClusterCIDR(t *testing.T) {
	g := NewWithT(t)

	cluster := newCluster("service-cidr-cluster")
	controlPlane := newAMCP("service-cidr-cluster")
	controlPlane.Spec.NetworkSpec.VPC.CidrBlock = "10.0.0.0/16"

	cluster.Spec.ClusterNetwork = &clusterv1.ClusterNetwork{
		Services: &clusterv1.NetworkRanges{
			CIDRBlocks: []string{"192.168.0.0/16"},
		},
	}

	g.Expect(determineClusterCIDR(cluster, controlPlane)).To(Equal("192.168.0.0/16"))

	cluster.Spec.ClusterNetwork.Services.CIDRBlocks = nil
	g.Expect(determineClusterCIDR(cluster, controlPlane)).To(Equal("10.0.0.0/16"))

	controlPlane.Spec.NetworkSpec.VPC.CidrBlock = ""
	g.Expect(determineClusterCIDR(&clusterv1.Cluster{}, &ekscontrolplanev1.AWSManagedControlPlane{})).To(Equal(""))
}

func TestDeriveDNSFromCIDR(t *testing.T) {
	g := NewWithT(t)

	ip, err := deriveDNSFromCIDR("192.168.0.0/16")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(ip).To(Equal("192.168.0.10"))

	ip, err = deriveDNSFromCIDR("fd00::/112")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(ip).To(Equal("fd00::a"))

	_, err = deriveDNSFromCIDR("not-a-cidr")
	g.Expect(err).To(HaveOccurred())
}

func configOwner(kind string) *bsutil.ConfigOwner {
	unstructuredOwner := unstructured.Unstructured{
		Object: map[string]interface{}{"kind": kind},
	}
	configOwner := bsutil.ConfigOwner{Unstructured: &unstructuredOwner}
	return &configOwner
}
