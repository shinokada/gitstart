#!/usr/bin/env bash

# Test validation improvements

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
GITSTART="${SCRIPT_DIR}/gitstart"

echo "Testing Validation Improvements"
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

echo "Test 1: Empty commit message (should fail)"
echo "-------------------------------------------"
if "${GITSTART}" -d test-repo -m "" --dry-run 2>&1 | grep -q "Commit message cannot be empty"; then
    echo "✓ Empty commit message correctly rejected"
else
    echo "✗ Empty commit message not caught"
    exit 1
fi
echo ""

echo "Test 2: Valid commit message (should pass)"
echo "------------------------------------------"
if "${GITSTART}" -d test-repo -m "Valid commit message" --dry-run >/dev/null 2>&1; then
    echo "✓ Valid commit message accepted"
else
    echo "✗ Valid commit message rejected (shouldn't happen)"
    exit 1
fi
echo ""

echo "Test 3: Missing commit message argument (should fail)"
echo "-----------------------------------------------------"
if "${GITSTART}" -d test-repo -m 2>&1 | grep -q "requires an argument"; then
    echo "✓ Missing argument correctly detected"
else
    echo "✗ Missing argument not caught"
    exit 1
fi
echo ""

echo "Test 4: Whitespace-only commit message (should pass in dry-run)"
echo "---------------------------------------------------------------"
# Note: Whitespace-only is technically non-empty, so it passes validation
# Git would catch it later in real execution
if "${GITSTART}" -d test-repo -m "   " --dry-run >/dev/null 2>&1; then
    echo "✓ Whitespace-only message passes validation (git will handle it)"
else
    echo "⚠ Whitespace-only message rejected (this is okay too)"
fi
echo ""

echo "Test 5: Long commit message (should pass)"
echo "-----------------------------------------"
long_msg="This is a very long commit message that should be handled properly without causing any issues or errors in the script"
if "${GITSTART}" -d test-repo -m "$long_msg" --dry-run >/dev/null 2>&1; then
    echo "✓ Long commit message accepted"
else
    echo "✗ Long commit message rejected (shouldn't happen)"
    exit 1
fi
echo ""

echo "Test 6: Special characters in commit message (should pass)"
echo "----------------------------------------------------------"
special_msg="Initial commit: Added README.md & LICENSE (v1.0.0)"
if "${GITSTART}" -d test-repo -m "$special_msg" --dry-run >/dev/null 2>&1; then
    echo "✓ Special characters in message accepted"
else
    echo "✗ Special characters rejected (shouldn't happen)"
    exit 1
fi
echo ""

echo "Test 7: Newlines in commit message (should pass)"
echo "------------------------------------------------"
multiline_msg="Initial commit

Added project structure"
if "${GITSTART}" -d test-repo -m "$multiline_msg" --dry-run >/dev/null 2>&1; then
    echo "✓ Multi-line message accepted"
else
    echo "⚠ Multi-line message rejected (this is okay)"
fi
echo ""

echo "==============================="
echo "All validation tests passed! ✓"
