# Test Fix: Empty Commit Message Validation

## The Problem

The test `gitstart handles empty commit message` was failing because:

1. **We added validation** that rejects empty commit messages (Round 3 fix)
2. **The test expected success** (`status -eq 0`) when passing `-m ""`
3. **The behavior changed** from accepting to rejecting empty messages

## Root Cause

The test was written before we added the empty message validation:

```bash
# Original test (WRONG):
@test "gitstart handles empty commit message" {
    run "$GITSTART_SCRIPT" -d test -m "" --dry-run
    [[ "$status" -eq 0 ]]  # ← Expected success, but now fails!
}
```

This test was checking that the script **accepts** empty messages, but our new validation **rejects** them (which is correct behavior).

## The Fix

Updated the test to expect failure and check for the error message:

```bash
# Fixed test (CORRECT):
@test "gitstart handles empty commit message" {
    run "$GITSTART_SCRIPT" -d test -m "" --dry-run
    [[ "$status" -eq 1 ]]  # ← Now expects failure
    [[ "$output" =~ "Commit message cannot be empty" ]] || [[ "$output" =~ "empty" ]]
}
```

## Why This is Correct

### Before Our Validation Fix
```bash
$ ./gitstart -d test -m ""
# Would proceed to git commit stage
# Then fail with: "error: empty message given, aborting commit"
# Test expected: status 0 (because validation wasn't failing)
```

### After Our Validation Fix
```bash
$ ./gitstart -d test -m ""
ERROR: Commit message cannot be empty
# Fails immediately at validation
# Test should expect: status 1 (validation correctly rejects)
```

## Test Expectations

The test name "handles empty commit message" is still accurate - it **handles** it by:
- ✓ Detecting the empty message
- ✓ Rejecting it with clear error
- ✓ Exiting with error code 1

## All Validation Test Cases

| Input | Expected Result | Why |
|-------|----------------|-----|
| `-m ""` | Fail (status 1) | Empty string not allowed |
| `-m` (no arg) | Fail (status 1) | Argument required |
| `-m "text"` | Pass (status 0) | Valid message |
| `-m "  "` | Pass (status 0) | Non-empty (git will handle) |
| `-m "Long..."` | Pass (status 0) | Length is okay |

## Testing

Run the comprehensive validation test:

```bash
chmod +x test-validation.sh
./test-validation.sh
```

Expected output:
```
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

## Run BATS Tests

```bash
bats tests/gitstart.bats -f "empty commit"
```

Should now show:
```
✓ gitstart handles empty commit message
```

## Summary

- **Test was outdated** after we added validation
- **Test now expects rejection** of empty messages (correct)
- **Behavior is correct**: Fail fast with clear error
- **Test name still accurate**: Script "handles" empty messages by rejecting them

**Status**: ✅ FIXED
