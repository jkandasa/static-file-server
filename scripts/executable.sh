#!/bin/bash


# update version details
source ./scripts/update_version.sh

EXECUTABLES_DIR=binary

mkdir -p ${EXECUTABLES_DIR}

# platforms to build
PLATFORMS=("linux/arm" "linux/arm64" "linux/386" "linux/amd64" "windows/386" "windows/amd64" "darwin/amd64" "linux/ppc64" "linux/s390x")

# compile
for platform in "${PLATFORMS[@]}"
do
  platform_raw=(${platform//\// })
  GOOS=${platform_raw[0]}
  GOARCH=${platform_raw[1]}
  EXECUTABLE="static-file-server-${VERSION}-${GOOS}-${GOARCH}"

  if [ $GOOS = "windows" ]; then
    EXECUTABLE+='.exe'
  fi

  env GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${EXECUTABLES_DIR}/${EXECUTABLE} -ldflags "$LD_FLAGS" cmd/main.go
  if [ $? -ne 0 ]; then
    echo 'an error has occurred. aborting the build process'
    exit 1
  fi
done

