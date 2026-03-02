#!/usr/bin/env bats

# Integration tests that require actual GitHub interaction
# These tests are separated because they:
# 1. Create actual GitHub repositories
# 2. Require network connectivity
# 3. Require GitHub CLI authentication
# 4. Take longer to run
# 5. Need cleanup of remote resources

# Run these tests with: bats tests/integration.bats
# Make sure to clean up after running!

setup() {
    export TEST_DIR="$(mktemp -d)"
    export GITSTART_SCRIPT="${BATS_TEST_DIRNAME}/../gitstart"
    export TEST_CONFIG_DIR="${TEST_DIR}/.config/gitstart"
    export XDG_CONFIG_HOME="${TEST_DIR}/.config"
    export TEST_REPO_PREFIX="gitstart-test-$$"  # Use PID for uniqueness
    
    mkdir -p "$TEST_CONFIG_DIR"
    
    # Get actual GitHub username
    export GH_USERNAME=$(gh api user --jq .login 2>/dev/null)
    if [[ -z "$GH_USERNAME" ]]; then
        skip "Not logged in to GitHub CLI"
    fi
    
    echo "$GH_USERNAME" > "$TEST_CONFIG_DIR/config"
    cd "$TEST_DIR" || {
        echo "Failed to change to test directory"
        return 1
    }
}

teardown() {
    # Clean up local test directory
    if [[ -d "$TEST_DIR" ]]; then
        cd /
        rm -rf "$TEST_DIR"
    fi
}

# Helper function to clean up GitHub repos
cleanup_github_repo() {
    local repo_name="$1"
    gh repo delete "${GH_USERNAME}/${repo_name}" --yes 2>/dev/null || true
}

# Test: Check GitHub authentication
@test "GitHub CLI is authenticated" {
    run gh auth status
    [[ "$status" -eq 0 ]]
}

# Test: Create simple public repository
@test "create simple public repository" {
    local repo_name="${TEST_REPO_PREFIX}-simple"
    
    # Create repo (non-interactive)
    # Note: This will prompt for license - would need expect/automation
    skip "Requires non-interactive license selection"
    
    # Cleanup
    cleanup_github_repo "$repo_name"
}

# Test: Create repository with Python .gitignore
@test "create repository with Python .gitignore" {
    skip "Requires full GitHub integration and cleanup"
}

# Test: Create private repository
@test "create private repository" {
    skip "Requires full GitHub integration and cleanup"
}

# Test: Create repository in existing directory
@test "initialize existing directory with files" {
    skip "Requires full GitHub integration and cleanup"
}

# Test: Create repository with custom branch
@test "create repository with custom branch name" {
    skip "Requires full GitHub integration and cleanup"
}

# Test: Verify remote repository exists
@test "created repository exists on GitHub" {
    skip "Requires full GitHub integration"
}

# Test: Verify files were pushed
@test "files are present in remote repository" {
    skip "Requires full GitHub integration"
}

# Test: Verify branch name
@test "custom branch name is used" {
    skip "Requires full GitHub integration"
}

# Test: Verify repository visibility
@test "repository has correct visibility" {
    skip "Requires full GitHub integration"
}

# Test: Verify .gitignore content
@test ".gitignore contains language-specific rules" {
    skip "Requires full GitHub integration"
}

# Test: Error handling - duplicate repository
@test "handles duplicate repository name gracefully" {
    skip "Requires full GitHub integration"
}

# Test: Cleanup on failure
@test "cleans up remote repo on failure" {
    skip "Requires full GitHub integration and error simulation"
}

# Note: Integration tests are marked as skip by default
# To enable them, remove the skip commands and run with proper cleanup
