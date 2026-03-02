#!/usr/bin/env bash

# Quick test to verify dry-run works non-interactively

set -euo pipefail

TEST_DIR="$(mktemp -d)"
GITSTART_SCRIPT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)/gitstart"
TEST_CONFIG_DIR="${TEST_DIR}/.config/gitstart"
export XDG_CONFIG_HOME="${TEST_DIR}/.config"

# Set up cleanup trap
cleanup() {
    rm -rf "$TEST_DIR"
}
trap cleanup EXIT

# Create config directory and username
mkdir -p "$TEST_CONFIG_DIR"
echo "testuser" > "$TEST_CONFIG_DIR/config"

# Change to test directory
cd "$TEST_DIR" || exit 1

echo "Running dry-run test..."
echo "========================"

# Run the script in dry-run mode (disable set -e temporarily)
set +e
"$GITSTART_SCRIPT" -d test-repo --dry-run
exit_code=$?
set -e

echo ""
echo "========================"
echo "Test completed with exit code: $exit_code"

if [[ $exit_code -eq 0 ]]; then
    echo "✓ Success: dry-run works without interaction"
    exit 0
else
    echo "✗ Failed: dry-run had issues"
    exit 1
fi
