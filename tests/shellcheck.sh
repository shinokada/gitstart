#!/usr/bin/env bash

# Shell script linting and static analysis using shellcheck
# This script runs shellcheck on the gitstart script and reports issues

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
GITSTART_SCRIPT="${SCRIPT_DIR}/../gitstart"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo "Running shellcheck on gitstart script..."
echo "========================================"
echo ""

# Check if shellcheck is installed
if ! command -v shellcheck &>/dev/null; then
    echo -e "${RED}ERROR: shellcheck is not installed${NC}"
    echo ""
    echo "Install it with:"
    echo "  macOS:    brew install shellcheck"
    echo "  Ubuntu:   sudo apt install shellcheck"
    echo "  Arch:     sudo pacman -S shellcheck"
    exit 1
fi

# Check if script exists
if [[ ! -f "$GITSTART_SCRIPT" ]]; then
    echo -e "${RED}ERROR: gitstart script not found at $GITSTART_SCRIPT${NC}"
    exit 1
fi

# Run shellcheck with various severity levels
echo "Checking for errors and warnings..."
echo ""

# Run shellcheck and capture exit code
set +e
shellcheck_output=$(shellcheck \
    --format=gcc \
    --severity=style \
    --enable=all \
    --exclude=SC2034 \
    "$GITSTART_SCRIPT" 2>&1)
shellcheck_exit=$?
set -e

# Display results
if [[ $shellcheck_exit -eq 0 ]]; then
    echo -e "${GREEN}✓ No issues found!${NC}"
    echo ""
    echo "The gitstart script passes all shellcheck checks."
    exit 0
else
    echo -e "${YELLOW}Issues found:${NC}"
    echo ""
    echo "$shellcheck_output"
    echo ""
    
    # Count issues by severity
    error_count=$(echo "$shellcheck_output" | grep -c "error:" || echo "0")
    warning_count=$(echo "$shellcheck_output" | grep -c "warning:" || echo "0")
    note_count=$(echo "$shellcheck_output" | grep -c "note:" || echo "0")
    
    echo "========================================"
    echo "Summary:"
    echo "  Errors:   $error_count"
    echo "  Warnings: $warning_count"
    echo "  Notes:    $note_count"
    echo "========================================"
    
    if [[ $error_count -gt 0 ]]; then
        echo -e "${RED}✗ Critical issues found - please fix errors${NC}"
        exit 1
    elif [[ $warning_count -gt 0 ]]; then
        echo -e "${YELLOW}⚠ Warnings found - consider addressing them${NC}"
        exit 0
    else
        echo -e "${GREEN}✓ Only style suggestions found${NC}"
        exit 0
    fi
fi
