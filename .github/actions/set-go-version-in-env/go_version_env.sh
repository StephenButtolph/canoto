#!/usr/bin/env bash

set -euo pipefail

# Prints the go version defined in the repo's go.mod. This is useful for
# configuring the correct version of go to install in CI.
#
# `go list -m -f '{{.GoVersion}}'` should be preferred outside of CI when go is
# already installed.

# 3 directories above this script
REPO_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd ../../.. && pwd )

echo GO_VERSION="~$(sed -n -e 's/^go //p' "${REPO_PATH}"/go.mod)"
