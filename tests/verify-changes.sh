#!/usr/bin/env bash

# Verification script - Shows all changes made

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"
GITSTART="${REPO_ROOT}/gitstart"
WORKFLOW="${REPO_ROOT}/.github/workflows/tests.yml"
RUN_TESTS="${REPO_ROOT}/tests/run-tests.sh"

echo "=========================================="
echo "Change Verification Summary"
echo "=========================================="
echo ""

echo "1. gitstart - Auth check (CRITICAL FIX)"
echo "   Line ~180: Skip gh auth in dry-run"
grep -A 3 "Skip auth check in dry-run mode" "${GITSTART}" || echo "   ❌ Not found!"
echo ""

echo "2. gitstart - Terminal checks (lines 199, 214, 244)"
echo "   Checking for '-t 0' additions..."
count=$(grep -c '\-t 0' "${GITSTART}" || echo "0")
echo "   Found ${count} instances (expected: 3)"
echo ""

echo "3. .github/workflows/tests.yml - Shell syntax"
echo "   Line 54: Using [[ ]] with quotes"
grep -A 1 'job.status' "${WORKFLOW}" || echo "   ❌ Not found!"
echo ""

echo "4. tests/run-tests.sh - Simplified conditional"
echo "   Checking CI_ENV conditional..."
grep -A 2 'CI_ENV.*true' "${RUN_TESTS}" | head -10 || echo "   ❌ Not found!"
echo ""

echo "=========================================="
echo "Verification Complete"
echo "=========================================="
echo ""
echo "Next steps:"
echo "1. Run local test: ./tests/test-dry-run-simple.sh"
echo "2. Commit changes: git add -A && git commit -m 'Fix: CI test improvements'"
echo "3. Push: git push"
echo "4. Watch CI pass all 35 tests ✓"
