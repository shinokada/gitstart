#!/usr/bin/env bash

# Test absolute path handling

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"
GITSTART="${REPO_ROOT}/gitstart"

echo "Testing Absolute Path Handling"
echo "==============================="
echo ""

# Setup test config
export TEST_DIR="$(mktemp -d)"
export XDG_CONFIG_HOME="${TEST_DIR}/.config"
mkdir -p "${XDG_CONFIG_HOME}/gitstart"
echo "testuser" > "${XDG_CONFIG_HOME}/gitstart/config"

cleanup() {
    rm -rf "${TEST_DIR}"
}
trap cleanup EXIT

echo "Test 1: Relative path"
echo "---------------------"
output=$("${GITSTART}" -d myrepo --dry-run 2>&1)
if echo "$output" | grep -q "$(pwd)/myrepo"; then
    echo "✓ Relative path correctly becomes $(pwd)/myrepo"
else
    echo "✗ Relative path handling failed"
    exit 1
fi
echo ""

echo "Test 2: Current directory (.)"
echo "-----------------------------"
output=$("${GITSTART}" -d . --dry-run 2>&1)
if echo "$output" | grep -q "$(pwd)"; then
    echo "✓ Current directory (.) correctly becomes $(pwd)"
else
    echo "✗ Current directory handling failed"
    exit 1
fi
echo ""

echo "Test 3: Absolute path"
echo "--------------------"
output=$("${GITSTART}" -d /tmp/test-absolute-repo --dry-run 2>&1)
if echo "$output" | grep -q "Directory:       /tmp/test-absolute-repo"; then
    echo "✓ Absolute path /tmp/test-absolute-repo kept as-is"
else
    echo "✗ Absolute path handling failed"
    echo "Output was:"
    echo "$output"
    exit 1
fi
echo ""

echo "Test 4: Empty commit message"
echo "---------------------------"
output=$("${GITSTART}" -d test-repo -m "" --dry-run 2>&1 || true)
if grep -Fq "Commit message cannot be empty" <<<"$output"; then
    echo "✓ Empty commit message rejected"
else
    echo "✗ Empty commit message not caught"
    echo "Output was: $output"
    exit 1
fi
echo ""

echo "==============================="
echo "All path handling tests passed! ✓"
