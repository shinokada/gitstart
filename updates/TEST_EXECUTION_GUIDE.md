# Test Execution Guide

## Quick Test Commands

### Run All Tests
```bash
cd /Users/shinichiokada/Bash/gitstart
./tests/run-tests.sh
```

### Run Individual Test Suites

**1. ShellCheck (Static Analysis)**
```bash
./tests/shellcheck.sh
```

**2. BATS Unit Tests**
```bash
bats ./tests/gitstart.bats
```

**3. Validation Tests**
```bash
./tests/test-validation.sh
```

**4. Path Handling Tests**
```bash
./tests/test-path-handling.sh
```

**5. Dry-Run Tests**
```bash
./tests/test-dry-run.sh
./tests/test-dry-run-simple.sh
```

**6. Quick Smoke Test**
```bash
./tests/quick-test.sh
```

## Expected Results After Fixes

### ✓ Fixed Issues

1. **Empty Commit Message Validation**
   - Before: "Option requires an argument" (wrong error)
   - After: "Commit message cannot be empty" (correct error)

2. **ShellCheck Arithmetic Error**
   - Before: `arithmetic syntax error in expression (error token is "0")`
   - After: Clean execution, proper counting

3. **BATS HOME Directory Test**
   - Before: Test failed (HOME not exported)
   - After: Test passes (HOME properly exported)

4. **Test Pipeline Failures**
   - Before: Tests exit early due to pipefail
   - After: Proper output capture with `|| true`

### Test Output Examples

**test-validation.sh:**
```bash
Testing Validation Improvements
===============================

Test 1: Empty commit message (should fail)
-------------------------------------------
✓ Empty commit message correctly rejected

Test 2: Valid commit message (should pass)
------------------------------------------
✓ Valid commit message accepted

Test 3: Missing commit message argument (should fail)
-----------------------------------------------------
✓ Missing argument correctly detected

...

===============================
All validation tests passed! ✓
```

**shellcheck.sh:**
```bash
Running shellcheck on gitstart script...
========================================

Checking for errors and warnings...
Issues found:
...
========================================
Summary:
  Errors:   0
  Warnings: 0
  Notes:    2
========================================
✓ Only style suggestions found
```

**gitstart.bats:**
```bash
gitstart.bats
 ✓ gitstart script exists and is executable
 ✓ gitstart -v returns version
 ✓ gitstart -h shows help
 ...
 ✓ gitstart handles empty commit message
 ✓ gitstart refuses to create repo in home directory
 ...

50 tests, 0 failures
```

## Troubleshooting

### If Tests Still Fail

**1. Check File Permissions**
```bash
./tests/fix-permissions.sh
chmod +x gitstart
chmod +x tests/*.sh
```

**2. Verify Changes Were Applied**
```bash
# Check the gitstart script has the fix
grep -A 2 "^    -m | --message)" gitstart

# Should show:
#     -m | --message)
#         [[ $# -ge 2 ]] || error "Option ${1} requires an argument"
#         [[ -n "${2:-}" ]] || error "Commit message cannot be empty"
```

**3. Check ShellCheck Fix**
```bash
# Check shellcheck.sh has the fix
grep -A 5 "error_count=" tests/shellcheck.sh

# Should show:
#     error_count=$(echo "$shellcheck_output" | grep -c "error:" || true)
#     warning_count=$(echo "$shellcheck_output" | grep -c "warning:" || true)
#     note_count=$(echo "$shellcheck_output" | grep -c "note:" || true)
#     
#     # Ensure counts are numeric
#     error_count=${error_count:-0}
```

**4. Check BATS Fix**
```bash
# Check gitstart.bats has the fix
grep -A 4 "refuses to create repo in home" tests/gitstart.bats

# Should show:
# @test "gitstart refuses to create repo in home directory" {
#     export HOME="$TEST_DIR"
#     cd "$HOME"
#     run "$GITSTART_SCRIPT" -d .
```

**5. Check Pipeline Fix**
```bash
# Check test-validation.sh has the fix
grep -A 3 "Test 1: Empty commit" tests/test-validation.sh

# Should show:
# echo "Test 1: Empty commit message (should fail)"
# echo "-------------------------------------------"
# output=$("${GITSTART}" -d test-repo -m "" --dry-run 2>&1 || true)
# if grep -Fq "Commit message cannot be empty" <<<"$output"; then
```

## Verification Checklist

Run each command and verify the output:

- [ ] `./tests/shellcheck.sh` - No arithmetic errors
- [ ] `./tests/test-validation.sh` - All tests pass
- [ ] `bats ./tests/gitstart.bats` - All 50+ tests pass
- [ ] `./tests/test-path-handling.sh` - All path tests pass
- [ ] `./tests/quick-test.sh` - Quick smoke test passes

## CI/CD Integration

The tests are now properly organized and will work in CI/CD:

**GitHub Actions Example:**
```yaml
- name: Run tests
  run: |
    chmod +x gitstart tests/*.sh
    ./tests/run-tests.sh
```

## Directory Organization

After cleanup, the project structure is:

```
gitstart/
├── gitstart                 # Main script (fixed)
├── tests/                   # All tests (organized)
│   ├── run-tests.sh        # Main test runner
│   ├── shellcheck.sh       # Static analysis (fixed)
│   ├── gitstart.bats       # Unit tests (fixed)
│   ├── test-validation.sh  # Validation tests (fixed, moved)
│   ├── test-path-handling.sh # Path tests (fixed, moved)
│   └── ...                 # Other test files
└── updates/                # Documentation
    └── PROJECT_CLEANUP_AND_FIXES.md
```

## Success Criteria

All these should pass:

1. ✓ Empty commit message correctly rejected
2. ✓ No ShellCheck arithmetic errors
3. ✓ All BATS tests pass
4. ✓ HOME directory test passes
5. ✓ Pipeline tests don't exit early
6. ✓ Root directory clean and organized

## Next Steps

After verifying all tests pass:

1. Commit the changes
2. Push to GitHub
3. Verify CI/CD pipeline passes
4. Update documentation if needed
5. Consider adding more test coverage
