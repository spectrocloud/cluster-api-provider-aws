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

package credentials

import (
	"github.com/spf13/cobra"
	"sigs.k8s.io/cluster-api-provider-aws/v2/cmd/clusterawsadm/cmd/util"

	"sigs.k8s.io/cluster-api-provider-aws/v2/cmd/clusterawsadm/controller/credentials"
)

// ZeroCredentialsCmd is a CLI command that will zero credentials the controller is started with.
func ZeroCredentialsCmd() *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "zero-credentials",
		Short: "zero credentials the controller is started with",
		Long: util.LongDesc(`
			Zero credentials the controller is started with
		`),
		Example: util.Examples(`
		# zero credentials
		# Kubeconfig file will be searched in default locations
		clusterawsadm controller zero-credentials --namespace=capa-system
		# Provided kubeconfig file will be used
		clusterawsadm controller zero-credentials --kubeconfig=kubeconfig  --namespace=capa-system
		# Kubeconfig in the default location will be retrieved and the provided context will be used
		clusterawsadm controller zero-credentials --kubeconfig-context=mgmt-cluster  --namespace=capa-system
		`),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return credentials.ZeroCredentials(credentials.ZeroCredentialsInput{
				KubeconfigPath:    kubeconfigPath,
				KubeconfigContext: kubeconfigContext,
				Namespace:         namespace,
			})
		},
	}
	addKubeconfigFlag(newCmd)
	addKubeconfigContextFlag(newCmd)
	addNamespaceFlag(newCmd)
	return newCmd
}
