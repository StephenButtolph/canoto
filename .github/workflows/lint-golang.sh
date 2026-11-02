#!/usr/bin/env bash

set -euo pipefail

if ! [[ "$0" =~ .github/workflows/lint-golang.sh ]]; then
  echo "must be run from repository root"
  exit 255
fi

go install -v github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.15.0

golangci-lint run --config .golangci.yml ./...
