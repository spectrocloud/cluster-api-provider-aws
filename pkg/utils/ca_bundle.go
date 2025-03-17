package utils

import (
	"context"
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

const (
	defaultProfile = "default"
)

func GetAWSCABundlePath(ctx context.Context, profile string) (string, error) {
	configFile := os.Getenv("AWS_CONFIG_FILE")
	if configFile == "" {
		return "", nil
	}

	cfg, err := ini.Load(configFile)
	if err != nil {
		return "", fmt.Errorf("failed to load AWS config: %w", err)
	}

	if profile == "" {
		profile = defaultProfile
	}

	section, err := cfg.GetSection(profile)
	if err != nil {
		return "", fmt.Errorf("failed to get profile %s: %w", profile, err)
	}

	cabundlePath := section.Key("ca_bundle")
	if cabundlePath == nil || cabundlePath.String() == "" {
		return "", nil
	}

	return cabundlePath.String(), nil
}

func GetAWSCABundle(ctx context.Context, profile string) ([]byte, error) {
	cabundlePath, err := GetAWSCABundlePath(ctx, profile)
	if err != nil {
		return nil, fmt.Errorf("failed to get AWS CA bundle path: %w", err)
	}

	if cabundlePath == "" {
		return nil, nil
	}

	cabundle, err := os.ReadFile(cabundlePath.String())
	if err != nil {
		return nil, fmt.Errorf("failed to read ca bundle file: %w", err)
	}

	return cabundle, nil
}
