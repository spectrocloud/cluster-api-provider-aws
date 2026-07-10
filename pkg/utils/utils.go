// Package utils has the common functions that can be used for cluster-api-provider-aws repo.
package utils

import (
	"context"
	"fmt"
	"math"
	"strings"

	"k8s.io/utils/ptr"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"

	clusterv1 "sigs.k8s.io/cluster-api/api/core/v1beta2"
)

// PartitionForRegion returns the AWS partition ID for a given region.
// It mirrors aws-sdk-go v1's endpoints.PartitionForRegion() using string
// prefix matching — the aws-sdk-go-v2 endpoint resolver no longer exposes
// a static region→partition table. The returned partition IDs match the
// partitionAwsIso*/partitionAwsIsoB constants in permissions_boundary.go
// and the v2 SDK's ARN partition strings, so TopSecret-region features
// (permissions boundary, IAM/EKS role creation) can key off the result.
func PartitionForRegion(region string) (string, error) {
	if region == "" {
		return "", fmt.Errorf("region is empty")
	}
	switch {
	case strings.HasPrefix(region, "us-gov-"):
		return "aws-us-gov", nil
	case strings.HasPrefix(region, "us-isob-"):
		return "aws-iso-b", nil
	case strings.HasPrefix(region, "us-iso-"):
		return "aws-iso", nil
	case strings.HasPrefix(region, "cn-"):
		return "aws-cn", nil
	default:
		return "aws", nil
	}
}

// GetMachinePools belong to a cluster.
func GetMachinePools(ctx context.Context, client crclient.Client, clusterName string, clusterNS string) ([]clusterv1.MachinePool, error) {
	machinePoolList := clusterv1.MachinePoolList{}
	listOptions := []crclient.ListOption{
		crclient.InNamespace(clusterNS),
		crclient.MatchingLabels(map[string]string{clusterv1.ClusterNameLabel: clusterName}),
	}

	if err := client.List(ctx, &machinePoolList, listOptions...); err != nil {
		return []clusterv1.MachinePool{}, fmt.Errorf("failed to list machine pools for cluster %s: %v", clusterName, err)
	}

	return machinePoolList.Items, nil
}

// ToInt64Pointer converts an int32 pointer to an int64 pointer.
func ToInt64Pointer(to *int32) *int64 {
	if to == nil {
		return nil
	}
	return ptr.To(int64(*to))
}

// ToInt64Value converts an int32 pointer to an int64 value.
func ToInt64Value(to *int32) int64 {
	if to == nil {
		panic("Cannot convert nil pointer to int64")
	}
	return int64(*to)
}

// ToInt32Pointer converts an int64 pointer to an int32 pointer.
func ToInt32Pointer(to *int64) *int32 {
	if to == nil {
		return nil
	}

	if *to > math.MaxInt32 || *to < math.MinInt32 {
		panic(fmt.Sprintf("Value %d is out of int32 range", *to))
	}

	return ptr.To(int32(*to))
}
