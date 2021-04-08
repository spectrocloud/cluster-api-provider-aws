/*
Copyright 2018 The Kubernetes Authors.

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

package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"sigs.k8s.io/cluster-api-provider-aws/cmd/clusterawsadm/cmd/alpha"
	"sigs.k8s.io/cluster-api-provider-aws/cmd/clusterawsadm/cmd/ami"
	"sigs.k8s.io/cluster-api-provider-aws/cmd/clusterawsadm/cmd/bootstrap"
	"sigs.k8s.io/cluster-api-provider-aws/cmd/clusterawsadm/cmd/eks"
	"sigs.k8s.io/cluster-api-provider-aws/cmd/clusterawsadm/cmd/version"
	"sigs.k8s.io/cluster-api/cmd/clusterctl/cmd"
)

// RootCmd is the Cobra root command
func RootCmd() *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "clusterawsadm",
		Short: "Kubernetes Cluster API Provider AWS Management Utility",
		Long: cmd.LongDesc(`
			clusterawsadm provides helpers for bootstrapping Kubernetes Cluster
			API Provider AWS. Use clusterawsadm to view required AWS Identity and Access Management
			(IAM) policies as JSON docs, or create IAM roles and instance profiles automatically
			using AWS CloudFormation.

			clusterawsadm additionally helps provide credentials for use with clusterctl.
		`),
		Example: cmd.Examples(`
			# Create AWS Identity and Access Management (IAM) roles for use with
			# Kubernetes Cluster API Provider AWS.
			clusterawsadm bootstrap iam create-cloudformation-stack

			# Encode credentials for use with clusterctl init
			export AWS_B64ENCODED_CREDENTIALS=$(clusterawsadm bootstrap credentials encode-as-profile)
			clusterctl init --infrastructure aws
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Help(); err != nil {
				return err
			}
			return nil
		},
	}
	newCmd.AddCommand(alpha.AlphaCmd())
	newCmd.AddCommand(bootstrap.RootCmd())
	newCmd.AddCommand(version.VersionCmd(os.Stdout))
	newCmd.AddCommand(ami.RootCmd())
	newCmd.AddCommand(eks.RootCmd())

	return newCmd
}

// Execute starts the process
func Execute() {
	if err := flag.CommandLine.Parse([]string{}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, "")
		os.Exit(1)
	}
	if err := RootCmd().Execute(); err != nil {
		os.Exit(1)
	}

	// Honor glog flags for verbosity control
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
}
