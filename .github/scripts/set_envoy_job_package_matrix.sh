#!/usr/bin/env bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

set -euo pipefail
export RUNNER_COUNT=$1

matrix="$(find ./test/integration/connect/envoy -maxdepth 1 -type d -print0 | xargs -0 -n 1 basename | jq --raw-input --argjson runnercount "$RUNNER_COUNT" -cM '[ inputs ] | [_nwise(length / $runnercount | floor)]')"

echo "envoy-matrix=${matrix}" >> "${GITHUB_OUTPUT}"
