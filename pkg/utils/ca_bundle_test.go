package utils

import (
	"context"
	"os"
	"testing"

	. "github.com/onsi/gomega"
	"gopkg.in/ini.v1"
)

func TestGetAWSCABundle(t *testing.T) {
	g := NewWithT(t)

	tests := []struct {
		name          string
		profile       string
		setupEnv      func()
		setupConfig   func() string
		expectedError string
		expectedData  []byte
	}{
		{
			name:    "Should return nil if AWS_CONFIG_FILE is not set",
			profile: "default",
			setupEnv: func() {
				os.Unsetenv("AWS_CONFIG_FILE")
			},
			expectedData: nil,
		},
		{
			name:    "Should return error if AWS config file cannot be loaded",
			profile: "default",
			setupEnv: func() {
				os.Setenv("AWS_CONFIG_FILE", "invalid_path")
			},
			expectedError: "failed to load AWS config",
		},
		{
			name:    "Should return error if profile section is not found",
			profile: "nonexistent",
			setupEnv: func() {
				os.Setenv("AWS_CONFIG_FILE", "test_config.ini")
			},
			setupConfig: func() string {
				cfg := ini.Empty()
				cfg.Section("default").Key("ca_bundle").SetValue("test_bundle.pem")
				cfg.SaveTo("test_config.ini")
				return "test_config.ini"
			},
			expectedError: "failed to get profile nonexistent",
		},
		{
			name:    "Should return nil if ca_bundle is not set in profile",
			profile: "default",
			setupEnv: func() {
				os.Setenv("AWS_CONFIG_FILE", "test_config.ini")
			},
			setupConfig: func() string {
				cfg := ini.Empty()
				cfg.Section("default")
				cfg.SaveTo("test_config.ini")
				return "test_config.ini"
			},
			expectedData: nil,
		},
		{
			name:    "Should return error if ca_bundle file cannot be read",
			profile: "default",
			setupEnv: func() {
				os.Setenv("AWS_CONFIG_FILE", "test_config.ini")
			},
			setupConfig: func() string {
				cfg := ini.Empty()
				cfg.Section("default").Key("ca_bundle").SetValue("invalid_path")
				cfg.SaveTo("test_config.ini")
				return "test_config.ini"
			},
			expectedError: "failed to read ca bundle file",
		},
		{
			name:    "Should return ca_bundle content if everything is correct",
			profile: "default",
			setupEnv: func() {
				os.Setenv("AWS_CONFIG_FILE", "test_config.ini")
			},
			setupConfig: func() string {
				cfg := ini.Empty()
				cfg.Section("default").Key("ca_bundle").SetValue("test_bundle.pem")
				cfg.SaveTo("test_config.ini")
				os.WriteFile("test_bundle.pem", []byte("test content"), 0644)
				return "test_config.ini"
			},
			expectedData: []byte("test content"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupEnv != nil {
				tt.setupEnv()
			}
			if tt.setupConfig != nil {
				tt.setupConfig()
			}

			data, err := GetAWSCABundle(context.Background(), tt.profile)
			if tt.expectedError != "" {
				g.Expect(err).To(HaveOccurred())
				g.Expect(err.Error()).To(ContainSubstring(tt.expectedError))
			} else {
				g.Expect(err).NotTo(HaveOccurred())
			}
			g.Expect(data).To(Equal(tt.expectedData))

			// Clean up
			os.Remove("test_config.ini")
			os.Remove("test_bundle.pem")
		})
	}
}
