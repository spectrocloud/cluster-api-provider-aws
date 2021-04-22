#!/bin/bash

rm ./generated/*
rm ./global/generated/*


# kustomize infra and managed separately to avoid multiple matches error
kustomize build global/infra > ./global/generated/infra-global.yaml
kustomize build global/managed > ./global/generated/managed-global.yaml

kustomize build global > ./generated/core-global.yaml
kustomize build base > ./generated/core-base.yaml
