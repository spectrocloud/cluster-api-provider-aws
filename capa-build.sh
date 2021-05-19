#!/bin/bash

if [ -z "$1" ]
  then
    echo "Please provide tag as the argument"
    exit 1
fi

export TAG=$1
export REGISTRY=gcr.io/spectro-images-public/cluster-api-aws/release

make docker-build

CORE_IMAGE_NAME=cluster-api-aws-controller
CORE_CONTROLLER_IMG=${REGISTRY}/${CORE_IMAGE_NAME}
EKS_CONTROLPLANE_IMAGE_NAME=eks-controlplane-controller
EKS_CONTROLPLANE_CONTROLLER_IMG=${REGISTRY}/${EKS_CONTROLPLANE_IMAGE_NAME}
EKS_BOOTSTRAP_IMAGE_NAME=eks-bootstrap-controller
EKS_BOOTSTRAP_CONTROLLER_IMG=${REGISTRY}/${EKS_BOOTSTRAP_IMAGE_NAME}

ARCH=amd64

docker tag ${CORE_CONTROLLER_IMG}-${ARCH}:${TAG} ${CORE_CONTROLLER_IMG}:${TAG}
docker tag ${EKS_BOOTSTRAP_CONTROLLER_IMG}-${ARCH}:${TAG} ${EKS_BOOTSTRAP_CONTROLLER_IMG}:${TAG}
docker tag ${EKS_CONTROLPLANE_CONTROLLER_IMG}-${ARCH}:${TAG} ${EKS_CONTROLPLANE_CONTROLLER_IMG}:${TAG}


docker push ${CORE_CONTROLLER_IMG}:${TAG}
docker push ${EKS_BOOTSTRAP_CONTROLLER_IMG}:${TAG}
docker push ${EKS_CONTROLPLANE_CONTROLLER_IMG}:${TAG}
