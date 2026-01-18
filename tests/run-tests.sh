#!/usr/bin/env bash

# Test runner script
# Runs all tests in the correct order
# Supports both local and CI environments

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

# Detect CI environment
CI_ENV="${CI:-false}"
if [[ "${GITHUB_ACTIONS:-false}" == "true" ]]; then
    CI_ENV="true"
fi

# Colors (disabled in CI for cleaner logs)
if [[ "$CI_ENV" == "true" ]]; then
    RED=''
    GREEN=''
    YELLOW=''
    BLUE=''
    NC=''
else
    RED='\033[0;31m'
    GREEN='\033[0;32m'
    YELLOW='\033[1;33m'
    BLUE='\033[0;34m'
    NC='\033[0m'
fi

total_passed=0
total_failed=0

print_header() {
    echo ""
    if [[ "$CI_ENV" == "true" ]]; then
        echo "========================================="
        echo "$1"
        echo "========================================="
    else
        echo -e "${BLUE}========================================${NC}"
        echo -e "${BLUE}$1${NC}"
        echo -e "${BLUE}========================================${NC}"
    fi
    echo ""
}

# CI-friendly logging
log_info() {
    if [[ "$CI_ENV" == "true" ]]; then
        echo "::notice::$1"
    else
        echo -e "${BLUE}$1${NC}"
    fi
}

log_success() {
    if [[ "$CI_ENV" == "true" ]]; then
        echo "::notice::✓ $1"
    else
        echo -e "${GREEN}✓ $1${NC}"
    fi
}

log_error() {
    if [[ "$CI_ENV" == "true" ]]; then
        echo "::error::✗ $1"
    else
        echo -e "${RED}✗ $1${NC}"
    fi
}

log_warning() {
    if [[ "$CI_ENV" == "true" ]]; then
        echo "::warning::⚠ $1"
    else
        echo -e "${YELLOW}⚠ $1${NC}"
    fi
}

run_shellcheck() {
    print_header "1. Running ShellCheck (Static Analysis)"
    
    if bash "${SCRIPT_DIR}/shellcheck.sh"; then
        log_success "ShellCheck passed"
        ((total_passed++)) || true
    else
        log_error "ShellCheck failed"
        ((total_failed++)) || true
        return 1
    fi
}

run_unit_tests() {
    print_header "2. Running Unit Tests (BATS)"
    
    if ! command -v bats &>/dev/null; then
        log_warning "BATS is not installed - skipping unit tests"
        echo ""
        echo "Install BATS with:"
        echo "  macOS:    brew install bats-core"
        echo "  Ubuntu:   sudo apt install bats"
        return 0
    fi
    
    # Run with appropriate formatter for environment
    if [[ "$CI_ENV" == "true" ]]; then
        # TAP format for CI (better parsing)
        if bats "${SCRIPT_DIR}/gitstart.bats" --formatter tap; then
            log_success "Unit tests passed"
            ((total_passed++)) || true
        else
            log_error "Unit tests failed"
            ((total_failed++)) || true
            return 1
        fi
    else
        # Pretty format for local development
        if bats "${SCRIPT_DIR}/gitstart.bats"; then
            log_success "Unit tests passed"
            ((total_passed++)) || true
        else
            log_error "Unit tests failed"
            ((total_failed++)) || true
            return 1
        fi
    fi
}

run_integration_tests() {
    print_header "3. Running Integration Tests (Optional)"
    
    if ! command -v bats &>/dev/null; then
        log_warning "BATS is not installed - skipping integration tests"
        return 0
    fi
    
    log_warning "Integration tests are currently skipped (require GitHub API)"
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
    local ci_info=""
    
    if [[ "$CI_ENV" == "true" ]]; then
        ci_info=" (CI Environment Detected)"
        echo "Running in CI environment${ci_info}"
        echo "CI Runner: ${RUNNER_OS:-unknown}"
        echo ""
    fi
    
    echo "Checking required dependencies..."
    
    if ! command -v shellcheck &>/dev/null; then
        missing_deps+=("shellcheck")
        log_warning "shellcheck not found"
    else
        log_success "shellcheck installed ($(shellcheck --version | head -n 2 | tail -n 1))"
    fi
    
    if ! command -v bats &>/dev/null; then
        missing_deps+=("bats")
        log_warning "bats not found"
    else
        log_success "bats installed ($(bats --version))"
    fi
    
    if ! command -v gh &>/dev/null; then
        log_warning "gh (GitHub CLI) not found - integration tests will fail"
    else
        log_success "gh (GitHub CLI) installed ($(gh --version | head -n 1))"
    fi
    
    if ! command -v jq &>/dev/null; then
        log_warning "jq not found - required for gitstart script"
    else
        log_success "jq installed ($(jq --version))"
    fi
    
    # Check if gitstart script is executable
    if [[ ! -x "$PROJECT_ROOT/gitstart" ]]; then
        log_error "gitstart script is not executable"
        echo "Run: chmod +x gitstart"
        if [[ "$CI_ENV" == "true" ]]; then
            echo "CI environments may need to set permissions explicitly"
        fi
    else
        log_success "gitstart script is executable"
    fi
    
    if [[ ${#missing_deps[@]} -gt 0 ]]; then
        echo ""
        log_warning "Missing optional dependencies: ${missing_deps[*]}"
        echo ""
        echo "Install them with:"
        echo "  macOS:    brew install ${missing_deps[*]}"
        echo "  Ubuntu:   sudo apt install ${missing_deps[*]}"
        echo ""
        
        if [[ "$CI_ENV" == "true" ]]; then
            echo "In CI, add to your workflow:"
            echo "  - name: Install dependencies"
            echo "    run: sudo apt-get install -y ${missing_deps[*]}"
        fi
    fi
}

print_summary() {
    print_header "Test Summary"
    
    local total=$((total_passed + total_failed))
    
    echo "Tests run:    $total"
    if [[ "$CI_ENV" == "true" ]]; then
        echo "Passed:       $total_passed"
        if [[ $total_failed -gt 0 ]]; then
            echo "Failed:       $total_failed"
        else
            echo "Failed:       $total_failed"
        fi
    else
        echo -e "Passed:       ${GREEN}$total_passed${NC}"
        if [[ $total_failed -gt 0 ]]; then
            echo -e "Failed:       ${RED}$total_failed${NC}"
        else
            echo -e "Failed:       $total_failed"
        fi
    fi
    
    echo ""
    
    if [[ $total_failed -eq 0 ]]; then
        if [[ "$CI_ENV" == "true" ]]; then
            echo "========================================="
            echo "All tests passed! ✓"
            echo "========================================="
        else
            echo -e "${GREEN}========================================${NC}"
            echo -e "${GREEN}All tests passed! ✓${NC}"
            echo -e "${GREEN}========================================${NC}"
        fi
        return 0
    else
        if [[ "$CI_ENV" == "true" ]]; then
            echo "========================================="
            echo "Some tests failed ✗"
            echo "========================================="
            echo "::error::Test suite failed with $total_failed failure(s)"
        else
            echo -e "${RED}========================================${NC}"
            echo -e "${RED}Some tests failed ✗${NC}"
            echo -e "${RED}========================================${NC}"
        fi
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
    if [[ "$CI_ENV" == "true" ]]; then
        echo "Environment: CI"
        echo "Runner: ${RUNNER_OS:-unknown}"
    else
        echo "Environment: Local"
    fi
    echo ""
    
    # Verify dependencies first
    verify_dependencies
    
    # Run tests
    local failed=0
    
    run_shellcheck || failed=$((failed + 1))
    run_unit_tests || failed=$((failed + 1))
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
        echo "Environment Variables:"
        echo "  CI=true              Enable CI mode (auto-detected)"
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
