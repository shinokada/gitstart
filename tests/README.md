# Gitstart Test Suite

This directory contains all testing infrastructure for the gitstart project.

## Test Files

### Static Analysis
- **`shellcheck.sh`** - Runs ShellCheck static analysis on the gitstart script
  - Checks for common shell scripting errors
  - Validates best practices
  - Reports errors, warnings, and notes

### Unit Tests (BATS)
- **`gitstart.bats`** - Unit tests using BATS (Bash Automated Testing System)
  - Tests individual functions and features
  - Validates argument parsing
  - Tests error handling
  - Checks dry-run mode
  
- **`integration.bats`** - Integration tests using BATS
  - Tests complete workflows
  - Validates end-to-end functionality

### Functional Tests
- **`test-validation.sh`** - Tests validation logic
  - Empty commit message handling
  - Valid commit message acceptance
  - Missing argument detection
  - Special characters handling
  - Multi-line messages
  
- **`test-path-handling.sh`** - Tests path resolution
  - Relative path conversion
  - Current directory (.) handling
  - Absolute path preservation
  
- **`test-dry-run.sh`** - Tests dry-run mode
  - Preview output validation
  - No side effects verification
  
- **`test-dry-run-simple.sh`** - Quick dry-run smoke test

### Utility Scripts
- **`run-tests.sh`** - Main test runner
  - Runs all tests in sequence
  - Provides comprehensive test report
  
- **`quick-test.sh`** - Quick smoke test
  - Fast sanity check
  - Useful for pre-commit validation
  
- **`verify-changes.sh`** - Verifies code changes
  - Checks for regressions
  - Validates fixes
  
- **`fix-permissions.sh`** - Fixes file permissions
  - Ensures scripts are executable
  - Maintains proper file modes

## Running Tests

### Run All Tests
```bash
./tests/run-tests.sh
```

### Run Individual Test Suites

**Static Analysis:**
```bash
./tests/shellcheck.sh
```

**Unit Tests:**
```bash
bats ./tests/gitstart.bats
bats ./tests/integration.bats
```

**Functional Tests:**
```bash
./tests/test-validation.sh
./tests/test-path-handling.sh
./tests/test-dry-run.sh
```

**Quick Smoke Test:**
```bash
./tests/quick-test.sh
```

## Test Requirements

### Dependencies
- `bash` 4.0 or higher
- `bats-core` - For BATS tests
- `shellcheck` - For static analysis
- `gh` (GitHub CLI) - For integration tests
- `jq` - For JSON processing

### Installation

**macOS:**
```bash
brew install bash bats-core shellcheck gh jq
```

**Ubuntu/Debian:**
```bash
sudo apt install bash bats shellcheck gh jq
```

**Arch Linux:**
```bash
sudo pacman -S bash bats shellcheck github-cli jq
```

## Test Structure

### Setup and Teardown
All tests use temporary directories for isolation:
- `TEST_DIR` - Temporary directory for each test
- `XDG_CONFIG_HOME` - Test config directory
- Cleanup happens automatically after each test

### Test Environment
Tests set up a controlled environment:
```bash
export TEST_DIR="$(mktemp -d)"
export XDG_CONFIG_HOME="${TEST_DIR}/.config"
mkdir -p "${XDG_CONFIG_HOME}/gitstart"
echo "testuser" > "${XDG_CONFIG_HOME}/gitstart/config"
```

### Cleanup
All tests clean up after themselves:
```bash
cleanup() {
    rm -rf "${TEST_DIR}"
}
trap cleanup EXIT
```

## Test Coverage

### Argument Parsing ✓
- All options (short and long forms)
- Required arguments
- Missing arguments
- Invalid options
- Multiple options

### Validation ✓
- Empty commit messages
- Missing directory
- Home directory protection
- Invalid paths
- Special characters

### Path Handling ✓
- Relative paths
- Absolute paths
- Current directory (.)
- Path normalization

### Dry-Run Mode ✓
- Preview output
- No file creation
- No GitHub interaction
- Configuration display

### Error Handling ✓
- Missing dependencies
- Invalid options
- Missing required arguments
- Edge cases

### Integration ✓
- Complete workflows
- Real-world scenarios
- Multi-option combinations

## CI/CD Integration

Tests are designed to run in CI/CD environments:
- Non-interactive mode supported
- Exit codes indicate success/failure
- Verbose output for debugging
- Parallel execution safe

### GitHub Actions
The tests run automatically on:
- Pull requests
- Push to main branch
- Manual workflow dispatch

## Adding New Tests

### BATS Test Template
```bash
@test "description of what you're testing" {
    # Setup
    setup_test_environment
    
    # Execute
    run "$GITSTART_SCRIPT" [options]
    
    # Assert
    [[ "$status" -eq 0 ]]
    [[ "$output" =~ "expected output" ]]
}
```

### Functional Test Template
```bash
#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
GITSTART="${SCRIPT_DIR}/../gitstart"

# Setup
export TEST_DIR="$(mktemp -d)"
cleanup() { rm -rf "${TEST_DIR}"; }
trap cleanup EXIT

# Test
echo "Test: Description"
output=$("${GITSTART}" [options] 2>&1 || true)
if [[ "$output" =~ "expected" ]]; then
    echo "✓ Test passed"
else
    echo "✗ Test failed"
    exit 1
fi
```

## Best Practices

1. **Isolation**: Each test should be independent
2. **Cleanup**: Always clean up test artifacts
3. **Deterministic**: Tests should produce consistent results
4. **Fast**: Keep tests quick for rapid feedback
5. **Clear**: Use descriptive test names and messages
6. **Comprehensive**: Cover happy paths and edge cases

## Debugging Tests

### Verbose Mode
Most test scripts support verbose output:
```bash
set -x  # Enable bash debug mode
./tests/test-validation.sh
```

### Individual Test Execution
Run specific BATS tests:
```bash
bats -f "test name pattern" ./tests/gitstart.bats
```

### Manual Testing
You can also test manually:
```bash
./gitstart -d test-repo --dry-run
```

## Troubleshooting

### Common Issues

**BATS not found:**
```bash
# Install BATS
brew install bats-core  # macOS
sudo apt install bats   # Ubuntu
```

**Permission denied:**
```bash
# Fix permissions
./tests/fix-permissions.sh
```

**Test failures:**
1. Check dependencies are installed
2. Ensure gitstart script is executable
3. Review test output for details
4. Run tests individually to isolate issues

## Contributing

When adding new features:
1. Write tests first (TDD approach)
2. Run all tests before committing
3. Update test documentation
4. Add new test files to this README

## Recent Changes

### 2026-01-18: Major Cleanup
- Moved all test files to `tests/` directory
- Fixed empty commit message validation
- Fixed ShellCheck arithmetic errors
- Fixed BATS HOME directory test
- Fixed pipeline handling in test scripts
- Consolidated test infrastructure

## References

- [BATS Documentation](https://bats-core.readthedocs.io/)
- [ShellCheck Wiki](https://www.shellcheck.net/)
- [Bash Testing Best Practices](https://github.com/bats-core/bats-core#writing-tests)
