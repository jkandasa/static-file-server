#!/bin/bash

# container registry
REGISTRY='docker.io/jkandasa'
IMAGE_NAME="static-file-server"
PLATFORMS="linux/arm/v6,linux/arm/v7,linux/arm64,linux/amd64"
IMAGE_TAG=`git rev-parse --abbrev-ref HEAD`

# build and push to docker.io
docker buildx build --push \
  --progress=plain \
  --build-arg=GOPROXY=${GOPROXY} \
  --platform ${PLATFORMS} \
  --file Dockerfile \
  --tag ${REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG} .