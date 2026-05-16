IMAGE=docker-hits
IMAGE_VERSION=1.2.1
GOVERSION=1.26.3

target:
	buildah build --build-arg GOVERSION=${GOVERSION} -t ${IMAGE}:${IMAGE_VERSION} .
	#buildah login docker.io
	buildah push ${IMAGE}:${IMAGE_VERSION} gounix/${IMAGE}:${IMAGE_VERSION}

