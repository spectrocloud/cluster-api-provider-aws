# Common set of functions
# Error check is done with set -e command . Build will fail if any of the commands fail

# Variables expected from CI - PULL_NUMBER , JOB_TYPE , ARTIFACTS , SONAR_SCAN_TOKEN, SONARQUBE_URL, DOCKER_REGISTRY

print_step() {
	text_val=$1
	set +x
	echo " "
	echo "###################################################
#  ${text_val}
###################################################"
	echo " "
	set -x
}

set_image_tag() {
	IMG_TAG="latest"
	if [[ ${JOB_TYPE} == 'presubmit' ]]; then
	    IMG_TAG=${PULL_NUMBER}
	    IMG_LOC='pr'
        fi
	if [[ ${JOB_TYPE} == 'periodic' ]]; then
	    IMG_TAG=$(date +%Y%m%d.%H%M)
	    IMG_LOC='daily'
	fi
	if [[ ${SPECTRO_RELEASE} == "yes" ]]; then
	    IMG_TAG=${VERSION}
	    IMG_LOC='release'
	fi
	export IMG_TAG
}

build_code() {
	print_step "Building Code"
	make all
}

create_images() {
	print_step "Create and Push the images"
	make binaries
	make release
}

delete_images() {
	print_step "Delete local images"
	echo make docker-rmi
}


create_manifest() {
	project_name=$1
	print_step "Create manifest files and copy to artifacts folder"
	# Manifest output has all secrets printed. Mask the output
	make manifest > /dev/null 2>&1

	mkdir -p ${ARTIFACTS}/${project_name}/build
	cp -r build/kustomize ${ARTIFACTS}/${project_name}/build/kustomize

	if [[ -d _build/manifests ]]; then
		cp -r _build/manifests ${ARTIFACTS}/manifests
	fi
}

run_lint() {
	print_step "Running Lint check"
	golangci-lint run    ./...  --timeout 10m  --tests=false
}



#----------------------------------------------/
# Scan containers with Anchore and Trivy       /
# Variables required are set in CI             /
#----------------------------------------------/
run_container_scan() {
	set +e
	print_step 'Run container scan'
	COMPL_DIR=${ARTIFACTS}/compliance
	CONTAINER_SCAN_DIR=${COMPL_DIR}/container_scan
	TRIVY_LIST=${CONTAINER_SCAN_DIR}/trivy_vulnerability.txt
	TRIVY_JSON=${CONTAINER_SCAN_DIR}/trivy_vulnerability.json
	mkdir -p ${CONTAINER_SCAN_DIR}

	for EACH_IMAGE in ${IMAGES_LIST}
	do
		trivy --download-db-only
 		echo "Image Name: ${EACH_IMAGE} " >> ${TRIVY_LIST}
		trivy ${EACH_IMAGE} >> ${TRIVY_LIST}
 	        trivy -f json ${EACH_IMAGE} >> ${TRIVY_JSON}
	done
	set -e
}


export REPO_NAME=cluster-api-provider-aws
export VERSION_SUFFIX=""
set_image_tag
export CONTROLLER_IMG=${DOCKER_REGISTRY}/${IMG_LOC}/cluster-api-aws/cluster-api-aws-controller
IMAGES_LIST="${CONTROLLER_IMG}"