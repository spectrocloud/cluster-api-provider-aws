package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

const (
	permissionsBoundaryFile = "/home/.aws/permissionsBoundary"

	// partitionAwsIso and partitionAwsIsoB are the AWS partition IDs
	// for the US ISO (TopSecret) and US ISOB (SecretRegion) partitions.
	// These string literals match aws-sdk-go v1's
	// endpoints.AwsIsoPartitionID / AwsIsoBPartitionID and the partition IDs
	// that the aws-sdk-go-v2 endpoint resolver returns.
	partitionAwsIso  = "aws-iso"
	partitionAwsIsoB = "aws-iso-b"
)

var (
	cachedPermissionsBoundary string
	cacheLock                 sync.Mutex
)

func isSecretPartition(partitionID string) bool {
	return partitionID == partitionAwsIso || partitionID == partitionAwsIsoB
}

func GetPermissionsBoundary(partitionID string) (string, error) {

	if !isSecretPartition(partitionID) {
		return "", nil
	}

	cacheLock.Lock()
	defer cacheLock.Unlock()

	if cachedPermissionsBoundary != "" {
		return cachedPermissionsBoundary, nil
	}

	permissionsBoundary, err := readPermissionsBoundaryFromFile(permissionsBoundaryFile)
	if err != nil {
		return "", err
	}

	cachedPermissionsBoundary = permissionsBoundary
	return cachedPermissionsBoundary, nil
}

func readPermissionsBoundaryFromFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if os.IsNotExist(err) {
		return "", nil
	} else if err != nil {
		return "", fmt.Errorf("failed to open permissions boundary file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			return line, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("failed to read permissions boundary file: %w", err)
	}

	return "", nil
}
