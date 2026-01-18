# Test Suite for Gitstart

This directory contains automated tests for the gitstart project using **shellcheck** (static analysis) and **bats** (functional testing).

## Test Structure

```
tests/
├── gitstart.bats          # Unit tests (BATS framework)
├── integration.bats       # Integration tests (requires GitHub)
├── shellcheck.sh          # Static analysis runner
├── run-tests.sh           # Main test runner
└── README.md             # This file
```

## Prerequisites

### Required
- **shellcheck** - Shell script static analysis
- **bats-core** - Bash Automated Testing System

### Installation

**macOS:**
```bash
brew install shellcheck bats-core
```

**Ubuntu/Debian:**
```bash
sudo apt install shellcheck bats
```

**Arch Linux:**
```bash
sudo pacman -S shellcheck bats
```

### Optional (for integration tests)
- **gh** - GitHub CLI (authenticated)
- **jq** - JSON processor

## Running Tests

### Run All Tests
```bash
# From project root
./tests/run-tests.sh

# Or from tests directory
cd tests
./run-tests.sh
```

### Run Specific Test Suites

**ShellCheck only:**
```bash
./tests/run-tests.sh --shellcheck-only
```

**Unit tests only:**
```bash
./tests/run-tests.sh --unit-only
```

**Integration tests only:**
```bash
./tests/run-tests.sh --integration-only
```

### Run Individual Tests

**ShellCheck:**
```bash
./tests/shellcheck.sh
```

**BATS unit tests:**
```bash
bats tests/gitstart.bats
```

**BATS integration tests:**
```bash
bats tests/integration.bats
```

**Run specific test:**
```bash
bats tests/gitstart.bats --filter "version"
```

## Test Categories

### 1. Static Analysis (shellcheck.sh)

Checks for:
- Syntax errors
- Common bash pitfalls
- Code style issues
- Security issues
- POSIX compliance

**Example output:**
```
Running shellcheck on gitstart script...
========================================

Checking for errors and warnings...

✓ No issues found!

The gitstart script passes all shellcheck checks.
```

### 2. Unit Tests (gitstart.bats)

Tests script functionality without external dependencies:
- Command-line argument parsing
- Version and help output
- Dry-run mode
- Configuration file handling
- Input validation
- Error handling
- Option combinations

**Example:**
```bash
✓ gitstart script exists and is executable
✓ gitstart -v returns version
✓ gitstart -h shows help
✓ gitstart without arguments shows error
✓ gitstart --dry-run shows preview without creating
✓ gitstart refuses to create repo in home directory
```

### 3. Integration Tests (integration.bats)

Tests real GitHub interactions (currently skipped by default):
- Creating actual repositories
- Pushing to GitHub
- Repository visibility
- File creation and push
- Error handling with GitHub API

**Note:** Integration tests require:
- GitHub CLI authentication (`gh auth login`)
- Network connectivity
- Manual cleanup of test repositories

## Test Coverage

Current test coverage includes:

### Command-Line Interface
- ✅ Version flag (`-v`)
- ✅ Help flag (`-h`)
- ✅ Required arguments validation
- ✅ All option flags (short and long)
- ✅ Option combinations
- ✅ Unknown option handling

### Configuration
- ✅ Config file creation
- ✅ Config file reading
- ✅ XDG directory compliance

### Validation
- ✅ Home directory protection
- ✅ Dependency checks
- ✅ Input validation

### Features
- ✅ Dry-run mode
- ✅ Quiet mode
- ✅ Custom branch names
- ✅ Custom commit messages
- ✅ Repository descriptions
- ✅ Private repositories
- ✅ Language selection

### Edge Cases
- ✅ Special characters in names
- ✅ Empty strings
- ✅ Long descriptions
- ✅ Multiple flags

## Writing New Tests

### BATS Test Structure

```bash
@test "description of test" {
    # Setup (optional)
    local test_var="value"
    
    # Run command
    run command_to_test arg1 arg2
    
    # Assertions
    [[ "$status" -eq 0 ]]           # Exit code
    [[ "$output" =~ "expected" ]]    # Output contains
    [[ -f "file.txt" ]]              # File exists
}
```

### Common BATS Assertions

```bash
# Exit codes
[[ "$status" -eq 0 ]]      # Success
[[ "$status" -eq 1 ]]      # Failure

# Output
[[ "$output" == "exact" ]]          # Exact match
[[ "$output" =~ "pattern" ]]        # Regex match
[[ -z "$output" ]]                  # Empty output
[[ -n "$output" ]]                  # Non-empty output

# Files
[[ -f "file" ]]           # File exists
[[ -d "dir" ]]            # Directory exists
[[ -x "script" ]]         # Executable
[[ ! -f "file" ]]         # File doesn't exist

# Strings
[[ "str1" == "str2" ]]    # Equal
[[ "str1" != "str2" ]]    # Not equal
```

### Adding a New Test

1. Open `tests/gitstart.bats`
2. Add test at the end:

```bash
@test "your new test description" {
    run "$GITSTART_SCRIPT" -d test --your-flag
    [[ "$status" -eq 0 ]]
    [[ "$output" =~ "expected output" ]]
}
```

3. Run tests to verify:
```bash
bats tests/gitstart.bats
```

## Continuous Integration

### GitHub Actions Example

Create `.github/workflows/tests.yml`:

```yaml
name: Tests

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Install dependencies
        run: |
          sudo apt-get update
          sudo apt-get install -y shellcheck bats
      
      - name: Run tests
        run: ./tests/run-tests.sh
```

### Pre-commit Hook

Add to `.git/hooks/pre-commit`:

```bash
#!/bin/bash
echo "Running tests before commit..."
./tests/run-tests.sh --shellcheck-only || exit 1
```

Make it executable:
```bash
chmod +x .git/hooks/pre-commit
```

## Troubleshooting

### "bats: command not found"

Install bats:
```bash
brew install bats-core  # macOS
sudo apt install bats   # Ubuntu
```

### "shellcheck: command not found"

Install shellcheck:
```bash
brew install shellcheck       # macOS
sudo apt install shellcheck   # Ubuntu
```

### Tests fail with "GitHub not authenticated"

Login to GitHub CLI:
```bash
gh auth login
```

### Integration tests creating repositories

Integration tests are skipped by default. They would create real GitHub repositories if enabled. Always clean up after running:

```bash
# List test repos
gh repo list | grep gitstart-test

# Delete test repo
gh repo delete username/gitstart-test-12345 --yes
```

## Test Maintenance

### When Adding New Features

1. Add unit tests in `gitstart.bats`
2. Update shellcheck exclusions if needed
3. Run full test suite
4. Update this README

### When Fixing Bugs

1. Write a test that reproduces the bug
2. Fix the bug
3. Verify test passes
4. Ensure all other tests still pass

### Regular Maintenance

```bash
# Run tests regularly
./tests/run-tests.sh

# Check coverage
# Count tests
grep -c "^@test" tests/gitstart.bats

# Review shellcheck warnings
./tests/shellcheck.sh
```

## Best Practices

1. **Run tests before committing**
   ```bash
   ./tests/run-tests.sh
   ```

2. **Write tests for new features**
   - Add test case before implementing feature (TDD)
   - Verify test fails initially
   - Implement feature
   - Verify test passes

3. **Keep tests fast**
   - Unit tests should run in seconds
   - Use mocks for external dependencies
   - Separate integration tests

4. **Make tests readable**
   - Use descriptive test names
   - One assertion per concept
   - Clear setup and teardown

5. **Test edge cases**
   - Empty strings
   - Special characters
   - Boundary conditions
   - Error conditions

## Future Improvements

- [ ] Increase test coverage to >90%
- [ ] Add mocking for external commands (gh, git)
- [ ] Create safe integration test environment
- [ ] Add performance benchmarks
- [ ] Generate coverage reports
- [ ] Add mutation testing
- [ ] Test on multiple OS versions

## Resources

- [BATS Documentation](https://bats-core.readthedocs.io/)
- [ShellCheck Wiki](https://github.com/koalaman/shellcheck/wiki)
- [Bash Testing Best Practices](https://github.com/sstephenson/bats/wiki/Best-practices)

## Support

For issues with tests:
1. Ensure dependencies are installed
2. Check test output for specific failures
3. Review test code for understanding
4. Open an issue with test output

---

**Happy Testing! 🧪**
