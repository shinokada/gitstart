#!/usr/bin/env bash
set -euo pipefail

# Run a quick local test
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TEST_SCRIPT="${SCRIPT_DIR}/test-dry-run-simple.sh"
chmod +x "${TEST_SCRIPT}"
"${TEST_SCRIPT}"
