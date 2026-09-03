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

package userdata

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
	"k8s.io/utils/ptr"

	eksbootstrapv1 "sigs.k8s.io/cluster-api-provider-aws/v2/bootstrap/eks/api/v1beta2"
	expinfrav1 "sigs.k8s.io/cluster-api-provider-aws/v2/exp/api/v1beta2"
)

func TestNewNode(t *testing.T) {
	format.TruncatedDiff = false
	g := NewWithT(t)

	type args struct {
		input *NodeInput
	}

	tests := []struct {
		name          string
		args          args
		expectedBytes []byte
		expectErr     bool
	}{
		{
			name: "only cluster name",
			args: args{
				input: &NodeInput{
					ClusterName: "test-cluster",
				},
			},
			expectedBytes: []byte(`#cloud-config
runcmd:
  - /etc/eks/bootstrap.sh test-cluster
`),
			expectErr: false,
		},
		{
			name: "sample-with-values",
			args: args{
				input: &NodeInput{
					ClusterName: "test-cluster",
					KubeletExtraArgs: map[string]string{
						"node-labels":          "node-role.undistro.io/infra=true",
						"register-with-taints": "dedicated=infra:NoSchedule",
					},
				},
			},
			expectedBytes: []byte(`#cloud-config
runcmd:
  - /etc/eks/bootstrap.sh test-cluster --kubelet-extra-args '--node-labels=node-role.undistro.io/infra=true --register-with-taints=dedicated=infra:NoSchedule'
`),
		},
		{
			name: "with container runtime",
			args: args{
				input: &NodeInput{
					ClusterName:      "test-cluster",
					ContainerRuntime: ptr.To[string]("containerd"),
				},
			},
			expectedBytes: []byte(`#cloud-config
runcmd:
  - /etc/eks/bootstrap.sh test-cluster --container-runtime containerd
`),
		},
		{
			name: "with kubelet extra args and container runtime",
			args: args{
				input: &NodeInput{
					ClusterName: "test-cluster",
					KubeletExtraArgs: map[string]string{
						"node-labels":          "node-role.undistro.io/infra=true",
						"register-with-taints": "dedicated=infra:NoSchedule",
					},
					ContainerRuntime: ptr.To[string]("containerd"),
				},
			},
			expectedBytes: []byte(`#cloud-config
runcmd:
  - /etc/eks/bootstrap.sh test-cluster --kubelet-extra-args '--node-labels=node-role.undistro.io/infra=true --register-with-taints=dedicated=infra:NoSchedule' --container-runtime containerd
`),
		},
		{
			name: "with ipv6",
			args: args{
				input: &NodeInput{
					ClusterName:     "test-cluster",
					ServiceIPV6Cidr: ptr.To[string]("fe80:0000:0000:0000:0204:61ff:fe9d:f156/24"),
					IPFamily:        ptr.To[string]("ipv6"),
				},
			},
			expectedBytes: []byte(`#cloud-config
runcmd:
  - /etc/eks/bootstrap.sh test-cluster --ip-family ipv6 --service-ipv6-cidr fe80:0000:0000:0000:0204:61ff:fe9d:f156/24
`),
		},
		{
			name: "without max pods",
			args: args{
				input: &NodeInput{
					ClusterName: "test-cluster",
					UseMaxPods:  ptr.To[bool](false),
				},
			},
			expectedBytes: []byte(`#cloud-config
runcmd:
  - /etc/eks/bootstrap.sh test-cluster --use-max-pods false
`),
		},
		{
			name: "with api retry attempts",
			args: args{
				input: &NodeInput{
					ClusterName:      "test-cluster",
					APIRetryAttempts: ptr.To[int](5),
				},
			},
			expectedBytes: []byte(`#cloud-config
runcmd:
  - /etc/eks/bootstrap.sh test-cluster --aws-api-retry-attempts 5
`),
		},
		{
			name: "with pause container",
			args: args{
				input: &NodeInput{
					ClusterName:           "test-cluster",
					PauseContainerAccount: ptr.To[string]("12345678"),
					PauseContainerVersion: ptr.To[string]("v1"),
				},
			},
			expectedBytes: []byte(`#cloud-config
runcmd:
  - /etc/eks/bootstrap.sh test-cluster --pause-container-account 12345678 --pause-container-version v1
`),
		},
		{
			name: "with dns cluster ip",
			args: args{
				input: &NodeInput{
					ClusterName:  "test-cluster",
					DNSClusterIP: ptr.To[string]("192.168.0.1"),
				},
			},
			expectedBytes: []byte(`#cloud-config
runcmd:
  - /etc/eks/bootstrap.sh test-cluster --dns-cluster-ip 192.168.0.1
`),
		},
		{
			name: "with docker json",
			args: args{
				input: &NodeInput{
					ClusterName:      "test-cluster",
					DockerConfigJSON: ptr.To[string]("{\"debug\":true}"),
				},
			},
			expectedBytes: []byte(`#cloud-config
runcmd:
  - /etc/eks/bootstrap.sh test-cluster --docker-config-json '{"debug":true}'
`),
		},
		{
			name: "with pre-bootstrap command",
			args: args{
				input: &NodeInput{
					ClusterName:          "test-cluster",
					PreBootstrapCommands: []string{"date", "echo \"testing\""},
				},
			},
			expectedBytes: []byte(`#cloud-config
runcmd:
  - "date"
  - "echo \"testing\""
  - /etc/eks/bootstrap.sh test-cluster
`),
		},
		{
			name: "with post-bootstrap command",
			args: args{
				input: &NodeInput{
					ClusterName:           "test-cluster",
					PostBootstrapCommands: []string{"date", "echo \"testing\""},
				},
			},
			expectedBytes: []byte(`#cloud-config
runcmd:
  - /etc/eks/bootstrap.sh test-cluster
  - "date"
  - "echo \"testing\""
`),
		},
		{
			name: "with pre & post-bootstrap command",
			args: args{
				input: &NodeInput{
					ClusterName:           "test-cluster",
					PreBootstrapCommands:  []string{"echo \"testing pre\""},
					PostBootstrapCommands: []string{"echo \"testing post\""},
				},
			},
			expectedBytes: []byte(`#cloud-config
runcmd:
  - "echo \"testing pre\""
  - /etc/eks/bootstrap.sh test-cluster
  - "echo \"testing post\""
`),
		},
		{
			name: "with bootstrap override command",
			args: args{
				input: &NodeInput{
					ClusterName:              "test-cluster",
					BootstrapCommandOverride: ptr.To[string]("/custom/mybootstrap.sh"),
				},
			},
			expectedBytes: []byte(`#cloud-config
runcmd:
  - /custom/mybootstrap.sh test-cluster
`),
		},
		{
			name: "with disk setup and mount points",
			args: args{
				input: &NodeInput{
					ClusterName: "test-cluster",
					DiskSetup: &eksbootstrapv1.DiskSetup{
						Filesystems: []eksbootstrapv1.Filesystem{
							{
								Device:     "/dev/sdb",
								Filesystem: "ext4",
								Label:      "vol2",
							},
						},
						Partitions: []eksbootstrapv1.Partition{
							{
								Device: "/dev/sdb",
								Layout: true,
							},
						},
					},
					Mounts: []eksbootstrapv1.MountPoints{
						[]string{"LABEL=vol2", "/mnt/vol2", "ext4", "defaults"},
						[]string{"LABEL=vol2", "/opt/data", "ext4", "defaults"},
					},
				},
			},
			expectedBytes: []byte(`#cloud-config
runcmd:
  - /etc/eks/bootstrap.sh test-cluster
disk_setup:
  /dev/sdb:
    layout: true
fs_setup:
  - label: vol2
    filesystem: ext4
    device: /dev/sdb
mounts:
  -
    - LABEL=vol2
    - /mnt/vol2
    - ext4
    - defaults
  -
    - LABEL=vol2
    - /opt/data
    - ext4
    - defaults
`),
		},
		{
			name: "with files",
			args: args{
				input: &NodeInput{
					ClusterName: "test-cluster",
					Files: []eksbootstrapv1.File{
						{
							Path:    "/etc/sysctl.d/91-fs-inotify.conf",
							Content: "fs.inotify.max_user_instances=256",
						},
					},
				},
			},
			expectedBytes: []byte(`#cloud-config
write_files:
  - path: /etc/sysctl.d/91-fs-inotify.conf
    content: |
      fs.inotify.max_user_instances=256
runcmd:
  - /etc/eks/bootstrap.sh test-cluster
`),
		},
		{
			name: "with empty files",
			args: args{
				input: &NodeInput{
					ClusterName: "test-cluster",
					Files:       []eksbootstrapv1.File{},
				},
			},
			expectedBytes: []byte(`#cloud-config
runcmd:
  - /etc/eks/bootstrap.sh test-cluster
`),
		},
		{
			name: "with nil files",
			args: args{
				input: &NodeInput{
					ClusterName: "test-cluster",
				},
			},
			expectedBytes: []byte(`#cloud-config
runcmd:
  - /etc/eks/bootstrap.sh test-cluster
`),
		},
		{
			name: "with ntp",
			args: args{
				input: &NodeInput{
					ClusterName: "test-cluster",
					NTP: &eksbootstrapv1.NTP{
						Enabled: aws.Bool(true),
						Servers: []string{"time1.google.com", "time2.google.com", "time3.google.com", "time4.google.com"},
					},
				},
			},
			expectedBytes: []byte(`#cloud-config
runcmd:
  - /etc/eks/bootstrap.sh test-cluster
ntp:
  enabled: true
  servers:
    - time1.google.com
    - time2.google.com
    - time3.google.com
    - time4.google.com
`),
		},
		{
			name: "with users",
			args: args{
				input: &NodeInput{
					ClusterName: "test-cluster",
					Users: []eksbootstrapv1.User{
						{
							Name:  "testuser",
							Shell: aws.String("/bin/bash"),
						},
					},
				},
			},
			expectedBytes: []byte(`#cloud-config
runcmd:
  - /etc/eks/bootstrap.sh test-cluster
users:
  - name: testuser
    shell: /bin/bash
`),
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			bytes, err := NewNode(testcase.args.input)
			if testcase.expectErr {
				g.Expect(err).To(HaveOccurred())
				return
			}

			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(string(bytes)).To(Equal(string(testcase.expectedBytes)))
		})
	}
}

// TestNewAL2023Node locks the AL2023 EKSConfig userdata layout. These bytes go into the
// pool's launch template, so any change here creates a new template version and repaves
// every node. Update the expectations only alongside a deliberate repave.
func TestNewAL2023Node(t *testing.T) {
	format.TruncatedDiff = false

	baseInput := func() *AL2023Input {
		return &AL2023Input{
			ClusterName:       "test-cluster",
			APIServerEndpoint: "https://ABC123.gr7.eu-west-1.eks.amazonaws.com",
			CACert:            "dGVzdC1jYQ==",
			NodeGroupName:     "worker-pool",
			ClusterCIDR:       "10.96.0.0/12",
		}
	}

	t.Run("no bootstrap commands emits only the node config part", func(t *testing.T) {
		g := NewWithT(t)

		out, err := NewAL2023Node(baseInput())

		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(string(out)).To(Equal(`MIME-Version: 1.0
Content-Type: multipart/mixed; boundary="//"


--//
Content-Type: application/node.eks.aws

---
apiVersion: node.eks.aws/v1alpha1
kind: NodeConfig
spec:
  cluster:
    name: test-cluster
    apiServerEndpoint: https://ABC123.gr7.eu-west-1.eks.amazonaws.com
    certificateAuthority: dGVzdC1jYQ==
    cidr: 10.96.0.0/12
  kubelet:
    config:
      maxPods: 110
      clusterDNS:
      - 10.96.0.10
    flags:
    - "--node-labels=eks.amazonaws.com/nodegroup-image=,eks.amazonaws.com/capacityType=ON_DEMAND,eks.amazonaws.com/nodegroup=worker-pool"

--//--`))
	})

	t.Run("bootstrap commands add a shell script part", func(t *testing.T) {
		g := NewWithT(t)

		input := baseInput()
		input.PreBootstrapCommands = []string{"echo pre1", "echo pre2"}
		input.PostBootstrapCommands = []string{"echo post1"}

		out, err := NewAL2023Node(input)

		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(string(out)).To(HavePrefix(`MIME-Version: 1.0
Content-Type: multipart/mixed; boundary="//"

--//
Content-Type: text/x-shellscript; charset="us-ascii"

#!/bin/bash
set -o errexit
set -o pipefail
set -o nounset
echo pre1
echo pre2
echo post1

--//
Content-Type: application/node.eks.aws
`))
	})

	t.Run("node labels come from kubeletExtraArgs when set", func(t *testing.T) {
		g := NewWithT(t)

		input := baseInput()
		input.KubeletExtraArgs = map[string]string{"node-labels": "custom=label"}

		out, err := NewAL2023Node(input)

		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(string(out)).To(ContainSubstring(`    - "--node-labels=custom=label"`))
	})

	t.Run("ami id and capacity type feed the default node labels", func(t *testing.T) {
		g := NewWithT(t)

		input := baseInput()
		input.AMIImageID = "ami-123"
		input.CapacityType = ptr.To(expinfrav1.ManagedMachinePoolCapacityTypeSpot)

		out, err := NewAL2023Node(input)

		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(string(out)).To(ContainSubstring(
			`eks.amazonaws.com/nodegroup-image=ami-123,eks.amazonaws.com/capacityType=spot,eks.amazonaws.com/nodegroup=worker-pool`))
	})

	t.Run("useMaxPods lowers the maxPods default", func(t *testing.T) {
		g := NewWithT(t)

		input := baseInput()
		input.UseMaxPods = ptr.To(true)

		out, err := NewAL2023Node(input)

		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(string(out)).To(ContainSubstring("maxPods: 58"))
	})

	t.Run("clusterDNS derives from the cluster cidr", func(t *testing.T) {
		g := NewWithT(t)

		input := baseInput()
		input.ClusterCIDR = "172.20.0.0/16"

		out, err := NewAL2023Node(input)

		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(string(out)).To(ContainSubstring("cidr: 172.20.0.0/16"))
		g.Expect(string(out)).To(ContainSubstring("- 172.20.0.10"))
	})

	t.Run("explicit dnsClusterIP wins over the derived one", func(t *testing.T) {
		g := NewWithT(t)

		input := baseInput()
		input.DNSClusterIP = ptr.To("10.100.0.53")

		out, err := NewAL2023Node(input)

		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(string(out)).To(ContainSubstring("- 10.100.0.53"))
	})

	t.Run("missing required fields are rejected", func(t *testing.T) {
		tests := []struct {
			name    string
			mutate  func(*AL2023Input)
			wantErr string
		}{
			{"no endpoint", func(i *AL2023Input) { i.APIServerEndpoint = "" }, "API server endpoint is required for AL2023"},
			{"no ca cert", func(i *AL2023Input) { i.CACert = "" }, "CA certificate is required for AL2023"},
			{"no cluster name", func(i *AL2023Input) { i.ClusterName = "" }, "cluster name is required for AL2023"},
			{"no nodegroup name", func(i *AL2023Input) { i.NodeGroupName = "" }, "node group name is required for AL2023"},
		}

		for _, testcase := range tests {
			t.Run(testcase.name, func(t *testing.T) {
				g := NewWithT(t)

				input := baseInput()
				testcase.mutate(input)

				_, err := NewAL2023Node(input)

				g.Expect(err).To(MatchError(testcase.wantErr))
			})
		}
	})
}
