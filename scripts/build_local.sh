#!/bin/bash

set -e

ROOT_DIR=$(dirname $(dirname $(realpath "$0")))
BUILD_DIR=${ROOT_DIR}/build

VERSION=$(bash $ROOT_DIR/scripts/get_version.sh)
GO_VERSION=$(go version)
PREFIX="agb/config"
LDFLAGS="-X '$PREFIX.appVersion=$VERSION' -X '$PREFIX.goVersion=$GO_VERSION'"

mkdir -p "${BUILD_DIR}"
cd "${ROOT_DIR}"/agb
go build -o "$BUILD_DIR" -ldflags "$LDFLAGS" ./cmd/agb
cd "${ROOT_DIR}"
