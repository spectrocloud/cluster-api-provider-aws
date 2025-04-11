package utils

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
)

type HttpHandler struct {
	CaCert             []byte
	InsecureSkipVerify bool
}

// NewHttpHandler creates a new HttpHandler with the given CA certificate and insecure skip verify flag.
func NewHttpHandler(caCert []byte, insecureSkipVerify bool) (*HttpHandler, error) {
	return &HttpHandler{
		CaCert:             caCert,
		InsecureSkipVerify: insecureSkipVerify,
	}, nil
}

// NewHttpHandlerWithAWSCABundle creates a new HttpHandler with the AWS CA bundle for the specified profile.
func NewHttpHandlerWithAWSCABundle(ctx context.Context, profile string) (*HttpHandler, error) {
	caBundle, err := GetAWSCABundle(ctx, profile)
	if err != nil {
		return nil, fmt.Errorf("failed to get AWS CA bundle: %w", err)
	}

	return NewHttpHandler(caBundle, caBundle == nil)
}

// GetHttpClient returns an HTTP client with the specified CA certificate and insecure skip verify flag.
func (h *HttpHandler) GetHttpClient() (*http.Client, error) {
	var tlsCfg tls.Config

	if h.InsecureSkipVerify {
		tlsCfg = tls.Config{
			InsecureSkipVerify: true,
			MinVersion:         tls.VersionTLS12,
		}
	} else {
		caCertPool, err := x509.SystemCertPool()
		if err != nil {
			return nil, fmt.Errorf("failed to load system cert pool: %w", err)
		} else if caCertPool == nil {
			caCertPool = x509.NewCertPool()
		}
		if len(h.CaCert) > 0 {
			caCertPool.AppendCertsFromPEM(h.CaCert)
		}

		tlsCfg = tls.Config{
			RootCAs:            caCertPool,
			MinVersion:         tls.VersionTLS12,
			InsecureSkipVerify: false,
		}
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tlsCfg,
			Proxy:           http.ProxyFromEnvironment,
		},
	}

	return httpClient, nil
}
