# ShellCheck SC2312 Fixes - 2026-01-18

## Issue Description

ShellCheck was reporting two SC2312 warnings about command substitutions that could mask return values.

### SC2312 Warning
> Consider invoking this command separately to avoid masking its return value (or use '|| true' to ignore).

This warning appears when a command substitution is used directly in a test condition, which can mask the command's exit status.

## Issues Found

### Issue 1: Line 241 - Directory Content Check
**Location:** `gitstart` line 241

**Original Code:**
```bash
# Check for files (including hidden files)
if [[ -n "$(ls -A "${dir}" 2>/dev/null)" ]]; then
    has_files=true
fi
```

**Problem:**
- If `ls` fails, the error is masked by the command substitution
- ShellCheck warns that we should handle the return value explicitly

**Fixed Code:**
```bash
# Check for files (including hidden files)
local dir_contents
dir_contents="$(ls -A "${dir}" 2>/dev/null || true)"
if [[ -n "${dir_contents}" ]]; then
    has_files=true
fi
```

**What Changed:**
1. Store command output in a variable first
2. Add `|| true` to explicitly ignore failures
3. Test the variable instead of inline command substitution

**Why This Is Better:**
- Explicit handling of potential failures
- Clearer separation of command execution and testing
- ShellCheck approved pattern

---

### Issue 2: Line 344 - Git Branch Detection
**Location:** `gitstart` line 344 (now ~420)

**Original Code:**
```bash
current_branch="$(git branch --show-current 2>/dev/null)"
current_branch="${current_branch:-}"
```

**Problem:**
- If `git branch --show-current` fails (e.g., in detached HEAD), error is masked
- ShellCheck wants explicit error handling

**Fixed Code:**
```bash
current_branch="$(git branch --show-current 2>/dev/null || true)"
current_branch="${current_branch:-}"
```

**What Changed:**
1. Added `|| true` to explicitly ignore failures
2. Makes it clear we expect this command might fail sometimes

**Why This Is Better:**
- Explicit that failures are acceptable here
- Works correctly in edge cases (detached HEAD, no branches)
- ShellCheck approved pattern

---

## Technical Explanation

### Why SC2312 Matters

When you write:
```bash
if [[ -n "$(some_command)" ]]; then
```

If `some_command` fails (exits with non-zero), the failure is masked because:
1. The command substitution captures output, not exit code
2. The test only checks if the output is non-empty
3. With `set -e`, this could cause unexpected behavior

### The Fix Pattern

The recommended pattern is:
```bash
# Option 1: Store first, then test
variable="$(command || true)"
if [[ -n "${variable}" ]]; then

# Option 2: Test separately
if command; then
    variable="$(command)"
```

We used Option 1 because we need the command output regardless of exit status.

---

## Testing

### Before Fixes
```bash
$ ./tests/shellcheck.sh
Issues found:
/Users/.../gitstart:241:17: note: Consider invoking this command separately... [SC2312]
/Users/.../gitstart:344:21: note: Consider invoking this command separately... [SC2312]

Summary:
  Errors:   0
  Warnings: 0
  Notes:    2
```

### After Fixes
```bash
$ ./tests/shellcheck.sh
✓ No issues found!

The gitstart script passes all shellcheck checks.
```

---

## Impact

### Code Quality
- ✅ All ShellCheck warnings resolved
- ✅ More explicit error handling
- ✅ Better code documentation

### Functionality
- ✅ No change in behavior
- ✅ More robust error handling
- ✅ Better edge case handling

### Testing
- ✅ All existing tests still pass
- ✅ No new issues introduced
- ✅ ShellCheck now passes cleanly

---

## Related ShellCheck Rules

### SC2312
**Title:** Consider invoking this command separately to avoid masking its return value

**Severity:** Note (informational)

**When it triggers:**
- Command substitution used directly in test condition
- Potential for masking command failures
- Better patterns exist

**How to fix:**
1. Store result in variable first
2. Add `|| true` if failures are acceptable
3. Test variable separately

**Example:**
```bash
# Bad
if [[ -n "$(command)" ]]; then

# Good
result="$(command || true)"
if [[ -n "${result}" ]]; then
```

---

## Best Practices Applied

1. **Explicit Error Handling**
   - Use `|| true` when failures are acceptable
   - Makes intent clear to readers
   - Satisfies ShellCheck requirements

2. **Variable Storage**
   - Store command output before testing
   - Separates execution from testing
   - More debuggable

3. **Clear Intent**
   - Code explicitly shows we expect possible failures
   - Future maintainers understand the logic
   - Better documentation

---

## Files Modified

### gitstart (2 changes)
- Line ~241: Fixed directory content check
- Line ~420: Fixed git branch detection

### No Test Changes Required
- All tests continue to pass
- No behavior changes
- Only code quality improvements

---

## Verification Steps

1. **Run ShellCheck:**
   ```bash
   ./tests/shellcheck.sh
   # Should show: ✓ No issues found!
   ```

2. **Run All Tests:**
   ```bash
   ./tests/run-tests.sh
   # Should pass all tests
   ```

3. **Test Edge Cases:**
   ```bash
   # Test with non-existent directory
   ./gitstart -d /tmp/test-$(date +%s) --dry-run
   
   # Test with existing directory
   mkdir /tmp/test-dir
   ./gitstart -d /tmp/test-dir --dry-run
   rm -rf /tmp/test-dir
   ```

---

## Summary

**Changes:** 2 lines modified  
**Issues Fixed:** 2 ShellCheck warnings (SC2312)  
**Impact:** Code quality improvement, no functional changes  
**Status:** ✅ Complete and verified

Both SC2312 warnings have been resolved by:
1. Separating command execution from testing
2. Adding explicit error handling with `|| true`
3. Using local variables to store intermediate results

The script now passes all ShellCheck checks with zero errors, zero warnings, and zero notes.

---

## References

- [ShellCheck SC2312](https://www.shellcheck.net/wiki/SC2312)
- [Bash Command Substitution](https://www.gnu.org/software/bash/manual/html_node/Command-Substitution.html)
- [Bash set -e behavior](https://www.gnu.org/software/bash/manual/html_node/The-Set-Builtin.html)
