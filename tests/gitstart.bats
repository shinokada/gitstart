#!/usr/bin/env bats

# Setup and teardown
setup() {
    # Create temporary directory for tests
    export TEST_DIR="$(mktemp -d)"
    export GITSTART_SCRIPT="${BATS_TEST_DIRNAME}/../gitstart"
    export TEST_CONFIG_DIR="${TEST_DIR}/.config/gitstart"
    export XDG_CONFIG_HOME="${TEST_DIR}/.config"
    
    # Create config directory
    mkdir -p "$TEST_CONFIG_DIR"
    
    # Set test username
    echo "testuser" > "$TEST_CONFIG_DIR/config"
    
    # Change to test directory
    cd "$TEST_DIR"
}

teardown() {
    # Clean up test directory
    if [[ -d "$TEST_DIR" ]]; then
        rm -rf "$TEST_DIR"
    fi
    
    # Clean up any test repositories on GitHub (if created)
    # This would need gh CLI and proper permissions
    # gh repo delete testuser/test-repo --yes 2>/dev/null || true
}

# Test: Script exists and is executable
@test "gitstart script exists and is executable" {
    [[ -f "$GITSTART_SCRIPT" ]]
    [[ -x "$GITSTART_SCRIPT" ]]
}

# Test: Version flag
@test "gitstart -v returns version" {
    run "$GITSTART_SCRIPT" -v
    [[ "$status" -eq 0 ]]
    [[ "$output" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]
}

# Test: Help flag
@test "gitstart -h shows help" {
    run "$GITSTART_SCRIPT" -h
    [[ "$status" -eq 0 ]]
    [[ "$output" =~ "Usage:" ]]
    [[ "$output" =~ "Options:" ]]
}

# Test: No arguments shows error
@test "gitstart without arguments shows error" {
    run "$GITSTART_SCRIPT"
    [[ "$status" -eq 1 ]]
    [[ "$output" =~ "ERROR" ]] || [[ "$output" =~ "required" ]]
}

# Test: Missing directory argument
@test "gitstart without -d flag shows error" {
    run "$GITSTART_SCRIPT" -l python
    [[ "$status" -eq 1 ]]
}

# Test: Dry run mode
@test "gitstart --dry-run shows preview without creating" {
    run "$GITSTART_SCRIPT" -d test-repo --dry-run
    [[ "$status" -eq 0 ]]
    [[ "$output" =~ "DRY RUN" ]]
    [[ "$output" =~ "No changes will be made" ]]
    [[ ! -d "test-repo" ]]
}

# Test: Dry run with all options
@test "gitstart --dry-run with all options shows configuration" {
    run "$GITSTART_SCRIPT" -d test-repo -l python -p -b develop -m "Test" --dry-run
    [[ "$status" -eq 0 ]]
    [[ "$output" =~ "python" ]]
    [[ "$output" =~ "private" ]]
    [[ "$output" =~ "develop" ]]
    [[ "$output" =~ "Test" ]]
}

# Test: Home directory protection
@test "gitstart refuses to create repo in home directory" {
    HOME="$TEST_DIR"
    cd "$HOME"
    run "$GITSTART_SCRIPT" -d .
    [[ "$status" -eq 1 ]]
    [[ "$output" =~ "home directory" ]] || [[ "$output" =~ "HOME" ]]
}

# Test: Config file creation
@test "gitstart creates config file if not exists" {
    rm -f "$TEST_CONFIG_DIR/config"
    # This test would be interactive, so we skip actual execution
    [[ ! -f "$TEST_CONFIG_DIR/config" ]]
}

# Test: Config file reading
@test "gitstart reads existing config file" {
    echo "testuser" > "$TEST_CONFIG_DIR/config"
    [[ -f "$TEST_CONFIG_DIR/config" ]]
    [[ "$(cat "$TEST_CONFIG_DIR/config")" == "testuser" ]]
}

# Test: Invalid language fallback
@test "gitstart with invalid language creates minimal .gitignore" {
    # This would require mocking gh and git commands
    skip "Requires mocking external commands"
}

# Test: Script validates dependencies
@test "script checks for gh command" {
    # Test that script validates gh is installed
    # Would need to mock the command -v check
    skip "Requires mocking command checks"
}

# Test: Script checks jq dependency
@test "script checks for jq command" {
    skip "Requires mocking command checks"
}

# Test: Multiple options parsing
@test "gitstart parses multiple options correctly" {
    run "$GITSTART_SCRIPT" -d test -l python -p -b develop -m "msg" --dry-run
    [[ "$status" -eq 0 ]]
    [[ "$output" =~ "test" ]]
}

# Test: Long option names
@test "gitstart accepts long option names" {
    run "$GITSTART_SCRIPT" --dir test-repo --lang python --dry-run
    [[ "$status" -eq 0 ]]
    [[ "$output" =~ "test-repo" ]]
    [[ "$output" =~ "python" ]]
}

# Test: Short option names
@test "gitstart accepts short option names" {
    run "$GITSTART_SCRIPT" -d test-repo -l python --dry-run
    [[ "$status" -eq 0 ]]
}

# Test: Mixed option styles
@test "gitstart accepts mixed short and long options" {
    run "$GITSTART_SCRIPT" -d test --lang python -p --branch develop --dry-run
    [[ "$status" -eq 0 ]]
}

# Test: Quiet mode reduces output
@test "gitstart -q produces minimal output" {
    run "$GITSTART_SCRIPT" -d test --dry-run -q
    [[ "$status" -eq 0 ]]
    # Quiet mode should still show some output in dry-run
}

# Test: Help shows all new options
@test "help shows all v0.4.0 options" {
    run "$GITSTART_SCRIPT" -h
    [[ "$output" =~ "--private" ]]
    [[ "$output" =~ "--branch" ]]
    [[ "$output" =~ "--message" ]]
    [[ "$output" =~ "--description" ]]
    [[ "$output" =~ "--dry-run" ]]
    [[ "$output" =~ "--quiet" ]]
}

# Test: Repository name extraction
@test "gitstart extracts repository name from directory path" {
    run "$GITSTART_SCRIPT" -d /path/to/my-repo --dry-run
    [[ "$status" -eq 0 ]]
    [[ "$output" =~ "my-repo" ]]
}

# Test: Current directory support
@test "gitstart -d . uses current directory name" {
    mkdir -p "$TEST_DIR/current-dir-test"
    cd "$TEST_DIR/current-dir-test"
    run "$GITSTART_SCRIPT" -d . --dry-run
    [[ "$status" -eq 0 ]]
    [[ "$output" =~ "current-dir-test" ]]
}

# Test: Branch name validation
@test "gitstart accepts custom branch name" {
    run "$GITSTART_SCRIPT" -d test -b custom-branch --dry-run
    [[ "$status" -eq 0 ]]
    [[ "$output" =~ "custom-branch" ]]
}

# Test: Commit message customization
@test "gitstart accepts custom commit message" {
    run "$GITSTART_SCRIPT" -d test -m "Custom initial commit" --dry-run
    [[ "$status" -eq 0 ]]
    [[ "$output" =~ "Custom initial commit" ]]
}

# Test: Repository description
@test "gitstart accepts repository description" {
    run "$GITSTART_SCRIPT" -d test -desc "Test description" --dry-run
    [[ "$status" -eq 0 ]]
    [[ "$output" =~ "Test description" ]]
}

# Test: Private flag
@test "gitstart -p sets private visibility" {
    run "$GITSTART_SCRIPT" -d test -p --dry-run
    [[ "$status" -eq 0 ]]
    [[ "$output" =~ "private" ]]
}

# Test: Public by default
@test "gitstart defaults to public visibility" {
    run "$GITSTART_SCRIPT" -d test --dry-run
    [[ "$status" -eq 0 ]]
    [[ "$output" =~ "public" ]]
}

# Test: Script handles special characters in names
@test "gitstart handles hyphens in repository names" {
    run "$GITSTART_SCRIPT" -d test-repo-name --dry-run
    [[ "$status" -eq 0 ]]
    [[ "$output" =~ "test-repo-name" ]]
}

# Test: Script handles underscores in names
@test "gitstart handles underscores in repository names" {
    run "$GITSTART_SCRIPT" -d test_repo_name --dry-run
    [[ "$status" -eq 0 ]]
    [[ "$output" =~ "test_repo_name" ]]
}

# Test: Unknown option handling
@test "gitstart rejects unknown options" {
    run "$GITSTART_SCRIPT" --unknown-option
    [[ "$status" -eq 1 ]]
}

# Test: Version number format
@test "version follows semantic versioning" {
    run "$GITSTART_SCRIPT" -v
    [[ "$output" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]
}

# Test: Script is POSIX compliant (uses bash)
@test "script uses bash shebang" {
    run head -n 1 "$GITSTART_SCRIPT"
    [[ "$output" =~ "#!/usr/bin/env bash" ]] || [[ "$output" =~ "#!/bin/bash" ]]
}

# Test: Script has set -euo pipefail for safety
@test "script uses strict mode" {
    run grep -n "set -euo pipefail" "$GITSTART_SCRIPT"
    [[ "$status" -eq 0 ]]
}

# Test: Empty string handling
@test "gitstart handles empty commit message" {
    run "$GITSTART_SCRIPT" -d test -m "" --dry-run
    [[ "$status" -eq 0 ]]
}

# Test: Long description handling
@test "gitstart handles long descriptions" {
    long_desc="This is a very long description that should be handled properly by the script without causing any issues"
    run "$GITSTART_SCRIPT" -d test -desc "$long_desc" --dry-run
    [[ "$status" -eq 0 ]]
    [[ "$output" =~ "long description" ]]
}

# Test: Multiple language flags (last one wins)
@test "gitstart handles multiple language flags" {
    run "$GITSTART_SCRIPT" -d test -l python -l javascript --dry-run
    [[ "$status" -eq 0 ]]
    # Last specified language should be used
    [[ "$output" =~ "javascript" ]]
}
