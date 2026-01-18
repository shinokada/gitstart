#!/usr/bin/env bash

# Test runner script
# Runs all tests in the correct order

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

total_passed=0
total_failed=0

print_header() {
    echo ""
    echo -e "${BLUE}========================================${NC}"
    echo -e "${BLUE}$1${NC}"
    echo -e "${BLUE}========================================${NC}"
    echo ""
}

run_shellcheck() {
    print_header "1. Running ShellCheck (Static Analysis)"
    
    if bash "${SCRIPT_DIR}/shellcheck.sh"; then
        echo -e "${GREEN}✓ ShellCheck passed${NC}"
        ((total_passed++))
    else
        echo -e "${RED}✗ ShellCheck failed${NC}"
        ((total_failed++))
        return 1
    fi
}

run_unit_tests() {
    print_header "2. Running Unit Tests (BATS)"
    
    if ! command -v bats &>/dev/null; then
        echo -e "${YELLOW}⚠ BATS is not installed - skipping unit tests${NC}"
        echo ""
        echo "Install BATS with:"
        echo "  macOS:    brew install bats-core"
        echo "  Ubuntu:   sudo apt install bats"
        return 0
    fi
    
    if bats "${SCRIPT_DIR}/gitstart.bats"; then
        echo -e "${GREEN}✓ Unit tests passed${NC}"
        ((total_passed++))
    else
        echo -e "${RED}✗ Unit tests failed${NC}"
        ((total_failed++))
        return 1
    fi
}

run_integration_tests() {
    print_header "3. Running Integration Tests (Optional)"
    
    if ! command -v bats &>/dev/null; then
        echo -e "${YELLOW}⚠ BATS is not installed - skipping integration tests${NC}"
        return 0
    fi
    
    echo -e "${YELLOW}Integration tests are currently skipped (require GitHub API)${NC}"
    echo "To run integration tests manually:"
    echo "  bats tests/integration.bats"
    echo ""
    echo "Note: Integration tests will create actual GitHub repositories"
    echo "      and require cleanup afterward."
    return 0
}

verify_dependencies() {
    print_header "0. Verifying Dependencies"
    
    local missing_deps=()
    
    echo "Checking required dependencies..."
    
    if ! command -v shellcheck &>/dev/null; then
        missing_deps+=("shellcheck")
        echo -e "${YELLOW}⚠ shellcheck not found${NC}"
    else
        echo -e "${GREEN}✓ shellcheck installed${NC}"
    fi
    
    if ! command -v bats &>/dev/null; then
        missing_deps+=("bats")
        echo -e "${YELLOW}⚠ bats not found${NC}"
    else
        echo -e "${GREEN}✓ bats installed${NC}"
    fi
    
    if ! command -v gh &>/dev/null; then
        echo -e "${YELLOW}⚠ gh (GitHub CLI) not found - integration tests will fail${NC}"
    else
        echo -e "${GREEN}✓ gh (GitHub CLI) installed${NC}"
    fi
    
    if ! command -v jq &>/dev/null; then
        echo -e "${YELLOW}⚠ jq not found - required for gitstart script${NC}"
    else
        echo -e "${GREEN}✓ jq installed${NC}"
    fi
    
    if [[ ${#missing_deps[@]} -gt 0 ]]; then
        echo ""
        echo -e "${YELLOW}Missing optional dependencies: ${missing_deps[*]}${NC}"
        echo ""
        echo "Install them with:"
        echo "  macOS:    brew install ${missing_deps[*]}"
        echo "  Ubuntu:   sudo apt install ${missing_deps[*]}"
        echo ""
    fi
}

print_summary() {
    print_header "Test Summary"
    
    local total=$((total_passed + total_failed))
    
    echo "Tests run:    $total"
    echo -e "Passed:       ${GREEN}$total_passed${NC}"
    
    if [[ $total_failed -gt 0 ]]; then
        echo -e "Failed:       ${RED}$total_failed${NC}"
    else
        echo -e "Failed:       $total_failed"
    fi
    
    echo ""
    
    if [[ $total_failed -eq 0 ]]; then
        echo -e "${GREEN}========================================${NC}"
        echo -e "${GREEN}All tests passed! ✓${NC}"
        echo -e "${GREEN}========================================${NC}"
        return 0
    else
        echo -e "${RED}========================================${NC}"
        echo -e "${RED}Some tests failed ✗${NC}"
        echo -e "${RED}========================================${NC}"
        return 1
    fi
}

main() {
    cd "$PROJECT_ROOT"
    
    echo "Gitstart Test Suite"
    echo "==================="
    echo ""
    echo "Project: $(basename "$PROJECT_ROOT")"
    echo "Script:  $PROJECT_ROOT/gitstart"
    echo ""
    
    # Verify dependencies first
    verify_dependencies
    
    # Run tests
    local failed=0
    
    run_shellcheck || ((failed++))
    run_unit_tests || ((failed++))
    run_integration_tests || true  # Don't count integration tests in failure
    
    # Print summary
    echo ""
    if print_summary; then
        exit 0
    else
        exit 1
    fi
}

# Parse command line arguments
case "${1:-}" in
    --help|-h)
        echo "Usage: $0 [OPTIONS]"
        echo ""
        echo "Options:"
        echo "  --help, -h           Show this help"
        echo "  --shellcheck-only    Run only shellcheck"
        echo "  --unit-only          Run only unit tests"
        echo "  --integration-only   Run only integration tests"
        echo ""
        echo "By default, runs all tests except integration tests."
        exit 0
        ;;
    --shellcheck-only)
        verify_dependencies
        run_shellcheck
        exit $?
        ;;
    --unit-only)
        verify_dependencies
        run_unit_tests
        exit $?
        ;;
    --integration-only)
        verify_dependencies
        run_integration_tests
        exit $?
        ;;
    "")
        main
        ;;
    *)
        echo "Unknown option: $1"
        echo "Use --help for usage information"
        exit 1
        ;;
esac
