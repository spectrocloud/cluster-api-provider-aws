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

package mime

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"html/template"
	"mime/multipart"
	"net/textproto"
	"os"
	"strings"

	"sigs.k8s.io/cluster-api-provider-aws/pkg/utils"
)

const (
	includePart = "file:///etc/secret-userdata.txt\n"
)

var (
	includeType = textproto.MIMEHeader{
		"content-type": {"text/x-include-url"},
	}

	boothookType = textproto.MIMEHeader{
		"content-type": {"text/cloud-boothook"},
	}

	multipartHeader = strings.Join([]string{
		"MIME-Version: 1.0",
		"Content-Type: multipart/mixed; boundary=\"%s\"",
		"\n",
	}, "\n")
)

type scriptVariables struct {
	SecretPrefix string
	Chunks       int32
	Region       string
	Endpoint     string
	CABundle     []byte
}

// GenerateInitDocument renders a given template, applies MIME properties
// and returns a series of byte chunks which put together represent a UserData
// script.
func GenerateInitDocument(secretPrefix string, chunks int32, region string, endpoint string, secretFetchScript string) ([]byte, error) {
	var secretFetchTemplate = template.Must(template.New("secret-fetch-script").Parse(secretFetchScript))

	var buf bytes.Buffer
	mpWriter := multipart.NewWriter(&buf)
	buf.WriteString(fmt.Sprintf(multipartHeader, mpWriter.Boundary()))
	scriptWriter, err := mpWriter.CreatePart(boothookType)
	if err != nil {
		return []byte{}, err
	}

	scriptVariables := scriptVariables{
		SecretPrefix: secretPrefix,
		Chunks:       chunks,
		Region:       region,
		Endpoint:     endpoint,
	}

	caBundle, err := utils.GetAWSCABundle(context.TODO(), "")
	if err != nil {
		return []byte{}, fmt.Errorf("failed to get AWS CA bundle: %w", err)
	}
	if caBundle != nil {
		scriptVariables.CABundle = caBundle
	}

	var scriptBuf bytes.Buffer
	if err := secretFetchTemplate.Execute(&scriptBuf, scriptVariables); err != nil {
		return []byte{}, err
	}
	_, err = scriptWriter.Write(scriptBuf.Bytes())
	if err != nil {
		return []byte{}, err
	}

	includeWriter, err := mpWriter.CreatePart(includeType)
	if err != nil {
		return []byte{}, err
	}

	_, err = includeWriter.Write([]byte(includePart))
	if err != nil {
		return []byte{}, err
	}

	if err := mpWriter.Close(); err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), nil
}

func getCABundlePathFromAWSConfig() string {
	awsConfigFile := os.Getenv("AWS_CONFIG_FILE")
	if awsConfigFile == "" {
		return ""
	}

	file, err := os.Open(awsConfigFile)
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "ca_bundle") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return ""
	}

	return ""
}
