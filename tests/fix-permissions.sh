#!/usr/bin/env bash

# Comprehensive fix script for gitstart project
# Restores executable permissions after file edits

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

echo "Fixing executable permissions..."
echo "================================"
echo ""

# Fix main script
if [[ -f "$SCRIPT_DIR/../gitstart" ]]; then
    chmod +x "$SCRIPT_DIR/../gitstart"
    echo "✓ Fixed: gitstart"
else
    echo "✗ Not found: gitstart"
fi

# Fix test scripts
for script in \
    "run-tests.sh" \
    "shellcheck.sh" \
    "test-dry-run.sh" \
    "test-dry-run-simple.sh" \
    "test-validation.sh" \
    "test-path-handling.sh" \
    "quick-test.sh" \
    "verify-changes.sh" \
    "fix-permissions.sh"
do
    if [[ -f "$SCRIPT_DIR/$script" ]]; then
        chmod +x "$SCRIPT_DIR/$script"
        echo "✓ Fixed: tests/$script"
    else
        echo "⚠ Not found: tests/$script"
    fi
done

echo ""
echo "================================"
echo "Permissions fixed!"
echo ""
echo "You can now run:"
echo "  ./gitstart --help"
echo "  ./tests/run-tests.sh"
echo ""
