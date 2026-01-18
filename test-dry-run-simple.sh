#!/usr/bin/env bash

# Quick test for dry-run mode in non-interactive CI environment

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
GITSTART="${SCRIPT_DIR}/gitstart"

# Create temporary config
export TEST_DIR="$(mktemp -d)"
export XDG_CONFIG_HOME="${TEST_DIR}/.config"
mkdir -p "${XDG_CONFIG_HOME}/gitstart"
echo "testuser" > "${XDG_CONFIG_HOME}/gitstart/config"

echo "Testing dry-run mode..."
echo "======================="

# Test 1: Basic dry-run
echo ""
echo "Test 1: Basic dry-run"
if "${GITSTART}" -d test-repo --dry-run; then
    echo "✓ Test 1 passed"
else
    echo "✗ Test 1 failed"
    exit 1
fi

# Test 2: Dry-run with all options
echo ""
echo "Test 2: Dry-run with all options"
if "${GITSTART}" -d test-repo -l python -p -b develop -m "Test commit" --description "Test description" --dry-run; then
    echo "✓ Test 2 passed"
else
    echo "✗ Test 2 failed"
    exit 1
fi

# Test 3: Dry-run without config file
echo ""
echo "Test 3: Dry-run without config (should use placeholder)"
rm -f "${XDG_CONFIG_HOME}/gitstart/config"
if "${GITSTART}" -d test-repo --dry-run 2>&1 | grep -q "<username>"; then
    echo "✓ Test 3 passed (placeholder username used)"
else
    echo "✗ Test 3 failed"
    exit 1
fi

# Cleanup
rm -rf "${TEST_DIR}"

echo ""
echo "======================="
echo "All tests passed! ✓"
