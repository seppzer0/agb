#!/bin/bash

set -e

ROOT_DIR=$(dirname $(dirname $(realpath "$0")))
BUILD_DIR=${ROOT_DIR}/build

mkdir -p "${BUILD_DIR}"
cd "${ROOT_DIR}"/agb
go build -o "$BUILD_DIR" ./cmd/agb
cd "${ROOT_DIR}"
