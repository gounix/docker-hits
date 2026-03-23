IMAGE=docker-hits
IMAGE_VERSION=1.2
GOVERSION=1.25.8

target:
	buildah build --build-arg GOVERSION=${GOVERSION} -t ${IMAGE}:${IMAGE_VERSION} .
	#buildah login docker.io
	buildah push ${IMAGE}:${IMAGE_VERSION} gounix/${IMAGE}:${IMAGE_VERSION}

