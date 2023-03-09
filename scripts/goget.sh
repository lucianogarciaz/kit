#! /bin/bash

set -x -e

TMP_DIR="$(mktemp -d)"
cd "${TMP_DIR}"
go mod init tmp
go get "$1"
rm -rf "${TMP_DIR}"
