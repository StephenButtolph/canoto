#!/usr/bin/env bash

set -euo pipefail

go install -v github.com/rhysd/actionlint/cmd/actionlint@v1.7.1

actionlint
