#!/bin/bash

set -e

ROOT_DIR=$(dirname $(dirname $(realpath "$0")))

export APP_VERSION=$(bash $ROOT_DIR/scripts/get_version.sh)
export GO_VERSION=$(go version)

cd $ROOT_DIR/agb
goreleaser build --snapshot --clean
