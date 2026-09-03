/*
Copyright 2026 The Kubernetes Authors.

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

// AL2023 userdata for the EKSConfig bootstrap provider. Separate from the nodeadm
// renderer in nodeadm.go, which NodeadmConfig uses. The rendered bytes must not change:
// any difference rewrites the launch template and repaves every node in the pool.

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"k8s.io/klog/v2"
	"k8s.io/utils/ptr"

	expinfrav1 "sigs.k8s.io/cluster-api-provider-aws/v2/exp/api/v1beta2"
)

const (
	// Shell script part template for AL2023
	al2023ShellScriptPartTemplate = `--{{.Boundary}}
Content-Type: text/x-shellscript; charset="us-ascii"

#!/bin/bash
set -o errexit
set -o pipefail
set -o nounset
{{- if or .PreBootstrapCommands .PostBootstrapCommands }}

{{- range .PreBootstrapCommands}}
{{.}}
{{- end}}
{{- range .PostBootstrapCommands}}
{{.}}
{{- end}}
{{- end}}`

	// Node config part template for AL2023.
	al2023NodeConfigPartTemplate = `
--{{.Boundary}}
Content-Type: application/node.eks.aws

---
apiVersion: node.eks.aws/v1alpha1
kind: NodeConfig
spec:
  cluster:
    name: {{.ClusterName}}
    apiServerEndpoint: {{.APIServerEndpoint}}
    certificateAuthority: {{.CACert}}
    cidr: {{if .ClusterCIDR}}{{.ClusterCIDR}}{{else}}10.96.0.0/12{{end}}
  kubelet:
    config:
      maxPods: {{.MaxPods}}
      clusterDNS:
      - {{.DNSClusterIP}}
    flags:
    - "--node-labels={{if and .KubeletExtraArgs (index .KubeletExtraArgs "node-labels")}}{{index .KubeletExtraArgs "node-labels"}}{{else}}eks.amazonaws.com/nodegroup-image={{if .AMIImageID}}{{.AMIImageID}}{{end}},eks.amazonaws.com/capacityType={{if .CapacityType}}{{.CapacityType}}{{else}}ON_DEMAND{{end}},eks.amazonaws.com/nodegroup={{.NodeGroupName}}{{end}}"

--{{.Boundary}}--`
)

// AL2023Input is the context for rendering AL2023 nodeadm userdata. It carries only the
// fields the two templates above read, plus UseMaxPods for the MaxPods default.
type AL2023Input struct {
	ClusterName           string
	KubeletExtraArgs      map[string]string
	PreBootstrapCommands  []string
	PostBootstrapCommands []string
	DNSClusterIP          *string
	UseMaxPods            *bool

	APIServerEndpoint string
	CACert            string
	NodeGroupName     string
	AMIImageID        string
	CapacityType      *expinfrav1.ManagedMachinePoolCapacityType
	MaxPods           *int32
	Boundary          string
	ClusterCIDR       string
}

// NewAL2023Node returns the userdata to bootstrap an AL2023 EKS node with nodeadm.
func NewAL2023Node(input *AL2023Input) ([]byte, error) {
	if err := validateAL2023Input(input); err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	// Write MIME header
	if _, err := buf.WriteString(fmt.Sprintf("MIME-Version: 1.0\nContent-Type: multipart/mixed; boundary=\"%s\"\n\n", input.Boundary)); err != nil {
		return nil, fmt.Errorf("failed to write MIME header: %v", err)
	}

	// Write shell script part if needed
	if len(input.PreBootstrapCommands) > 0 || len(input.PostBootstrapCommands) > 0 {
		shellScriptTemplate := template.Must(template.New("shell").Parse(al2023ShellScriptPartTemplate))
		if err := shellScriptTemplate.Execute(&buf, input); err != nil {
			return nil, fmt.Errorf("failed to execute shell script template: %v", err)
		}
		if _, err := buf.WriteString("\n"); err != nil {
			return nil, fmt.Errorf("failed to write newline: %v", err)
		}
	}

	// Write node config part
	nodeConfigTemplate := template.Must(template.New("node").Parse(al2023NodeConfigPartTemplate))
	if err := nodeConfigTemplate.Execute(&buf, input); err != nil {
		return nil, fmt.Errorf("failed to execute node config template: %v", err)
	}

	return buf.Bytes(), nil
}

// getCapacityTypeString returns the string representation of the capacity type.
func (ni *AL2023Input) getCapacityTypeString() string {
	if ni.CapacityType == nil {
		return "ON_DEMAND"
	}
	switch *ni.CapacityType {
	case expinfrav1.ManagedMachinePoolCapacityTypeSpot:
		return "SPOT"
	case expinfrav1.ManagedMachinePoolCapacityTypeOnDemand:
		return "ON_DEMAND"
	default:
		return strings.ToUpper(string(*ni.CapacityType))
	}
}

// validateAL2023Input validates the input for AL2023 user data generation.
func validateAL2023Input(input *AL2023Input) error {
	if input.APIServerEndpoint == "" {
		return fmt.Errorf("API server endpoint is required for AL2023")
	}
	if input.CACert == "" {
		return fmt.Errorf("CA certificate is required for AL2023")
	}
	if input.ClusterName == "" {
		return fmt.Errorf("cluster name is required for AL2023")
	}
	if input.NodeGroupName == "" {
		return fmt.Errorf("node group name is required for AL2023")
	}

	if input.MaxPods == nil {
		if input.UseMaxPods != nil && *input.UseMaxPods {
			input.MaxPods = ptr.To[int32](58)
		} else {
			input.MaxPods = ptr.To[int32](110)
		}
	}
	if input.DNSClusterIP == nil {
		if input.ClusterCIDR != "" {
			input.DNSClusterIP = ptr.To(calculateDNSFromServiceCIDR(input.ClusterCIDR))
		} else {
			input.DNSClusterIP = ptr.To[string]("10.96.0.10")
		}
	}

	if input.Boundary == "" {
		input.Boundary = boundary
	}

	klog.V(2).Infof("AL2023 Userdata Generation - maxPods: %d, clusterDNS: %s, amiID: %s, capacityType: %s",
		*input.MaxPods, *input.DNSClusterIP, input.AMIImageID, input.getCapacityTypeString())

	return nil
}

// calculateDNSFromServiceCIDR calculates the DNS cluster IP by replacing the last octet with 10.
func calculateDNSFromServiceCIDR(cidr string) string {
	ipStr := strings.Split(strings.TrimSpace(cidr), "/")[0]
	lastDotIdx := strings.LastIndex(ipStr, ".")
	return ipStr[:lastDotIdx] + ".10"
}
