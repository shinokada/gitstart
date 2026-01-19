# Testing Guide for Gitstart

This document provides test cases and scenarios for validating gitstart functionality.

## Prerequisites

Before testing, ensure you have:
- GitHub CLI (`gh`) installed and authenticated
- `jq` installed
- Bash shell available
- Test GitHub account (don't use production account for destructive tests)

## Test Environment Setup

```bash
# Create a test directory
mkdir -p ~/gitstart-tests
cd ~/gitstart-tests

# Ensure you're logged in
gh auth status

# Note: Some tests will create repositories on GitHub
# You may want to delete them after testing
```

## Test Categories

### 1. Basic Functionality Tests

#### Test 1.1: Create New Repository (Public)
```bash
gitstart -d test-repo-public
# Expected: Creates public repo, initializes git, pushes to GitHub
# Verify: Check https://github.com/YOUR_USERNAME/test-repo-public
```

#### Test 1.2: Create New Repository (Private)
```bash
gitstart -d test-repo-private -p
# Expected: Creates private repo
# Verify: Check repo visibility on GitHub
```

#### Test 1.3: With Programming Language
```bash
gitstart -d test-python -l python
# Expected: Downloads Python .gitignore from GitHub
# Verify: cat test-python/.gitignore | grep "*.pyc"
```

#### Test 1.4: Custom Commit Message
```bash
gitstart -d test-commit -m "Initial release v1.0"
# Expected: First commit has custom message
# Verify: cd test-commit && git log --oneline
```

#### Test 1.5: Custom Branch Name
```bash
gitstart -d test-branch -b develop
# Expected: Creates 'develop' branch instead of 'main'
# Verify: cd test-branch && git branch
```

#### Test 1.6: With Description
```bash
gitstart -d test-desc -desc "My test repository"
# Expected: Repository has description on GitHub
# Verify: gh repo view YOUR_USERNAME/test-desc | grep "description"
```

### 2. Existing Directory Tests

#### Test 2.1: Empty Existing Directory
```bash
mkdir test-empty-dir
gitstart -d test-empty-dir
# Expected: Works normally
```

#### Test 2.2: Directory with Files (No Git)
```bash
mkdir test-with-files
cd test-with-files
echo "existing file" > existing.txt
cd ..
gitstart -d test-with-files
# Expected: Warns about existing files, asks for confirmation
# Verify: existing.txt should be in the repo
```

#### Test 2.3: Directory with Existing Git Repo
```bash
mkdir test-existing-git
cd test-existing-git
git init
echo "# Test" > README.md
git add README.md
git commit -m "Local commit"
cd ..
gitstart -d test-existing-git
# Expected: Detects existing git, adds remote, preserves history
# Verify: git log should show "Local commit"
```

#### Test 2.4: Current Directory
```bash
mkdir test-current-dir && cd test-current-dir
gitstart -d .
# Expected: Initializes repo in current directory
cd ..
```

#### Test 2.5: Existing LICENSE File
```bash
mkdir test-existing-license
cd test-existing-license
echo "Custom License" > LICENSE
cd ..
gitstart -d test-existing-license
# Expected: Skips creating LICENSE
# Verify: LICENSE should contain "Custom License"
```

#### Test 2.6: Existing README.md
```bash
mkdir test-existing-readme
cd test-existing-readme
echo "# My Project" > README.md
cd ..
gitstart -d test-existing-readme
# Expected: Skips creating README.md
# Verify: README should contain "# My Project"
```

#### Test 2.7: Existing .gitignore
```bash
mkdir test-existing-gitignore
cd test-existing-gitignore
echo "custom_ignore.txt" > .gitignore
cd ..
gitstart -d test-existing-gitignore -l python
# Expected: Asks to append to existing .gitignore
# Manual: Choose 'y' to append
# Verify: .gitignore should contain both custom and Python ignores
```

### 3. Dry Run Tests

#### Test 3.1: Basic Dry Run
```bash
gitstart -d test-dry-run --dry-run
# Expected: Shows what would happen, no changes made
# Verify: Directory test-dry-run should NOT exist
```

#### Test 3.2: Dry Run with All Options
```bash
gitstart -d test-dry-full --dry-run -l python -p -b develop -m "Test" -desc "Description"
# Expected: Shows complete configuration, no changes made
```

### 4. Error Handling Tests

#### Test 4.1: No Directory Specified
```bash
gitstart
# Expected: ERROR message about missing directory argument
# Exit code: 1
```

#### Test 4.2: Invalid Language
```bash
gitstart -d test-invalid-lang -l nonexistentlanguage
# Expected: Falls back to minimal .gitignore
# Verify: .gitignore contains only .DS_Store
```

#### Test 4.3: Home Directory Protection
```bash
cd ~
gitstart -d .
# Expected: ERROR about not allowing repo in home directory
# Exit code: 1
```

#### Test 4.4: Repository Already Exists
```bash
gitstart -d test-duplicate
gitstart -d test-duplicate
# Expected: Error from gh about existing repository
# Verify: Automatic cleanup should occur
```

#### Test 4.5: Not Logged into GitHub
```bash
gh auth logout
gitstart -d test-no-auth
# Expected: ERROR about needing to login
# Cleanup: gh auth login
```

### 5. Configuration Tests

#### Test 5.1: First Run (No Config)
```bash
# Remove config if exists
rm -rf ~/.config/gitstart
gitstart -d test-first-run
# Expected: Prompts for GitHub username
# Expected: Saves to ~/.config/gitstart/config
```

#### Test 5.2: Config File Exists
```bash
# Ensure config exists
mkdir -p ~/.config/gitstart
echo "testuser" > ~/.config/gitstart/config
gitstart -d test-with-config
# Expected: Reads username from config, asks for confirmation
```

#### Test 5.3: Wrong Username in Config
```bash
echo "wronguser" > ~/.config/gitstart/config
gitstart -d test-wrong-user
# Expected: Asks if username is correct, allows changing
# Manual: Choose 'n' and provide correct username
```

### 6. License Selection Tests

#### Test 6.1: MIT License
```bash
gitstart -d test-mit
# Expected: Prompts for license selection
# Manual: Choose option 1 (MIT)
# Verify: cat test-mit/LICENSE | grep "MIT License"
```

#### Test 6.2: Apache License
```bash
gitstart -d test-apache
# Manual: Choose option 2 (Apache 2.0)
# Verify: cat test-apache/LICENSE | grep "Apache License"
```

#### Test 6.3: GNU GPLv3 License
```bash
gitstart -d test-gnu
# Manual: Choose option 3 (GNU GPLv3)
# Verify: cat test-gnu/LICENSE | grep "GNU"
```

#### Test 6.4: No License
```bash
gitstart -d test-no-license
# Manual: Choose option 4 (None)
# Verify: LICENSE file should not exist
```

### 7. Quiet Mode Tests

#### Test 7.1: Quiet Mode
```bash
gitstart -d test-quiet -q
# Expected: Minimal output, no prompts
# Note: Uses default values for license etc.
```

### 8. Integration Tests

#### Test 8.1: Complete Workflow
```bash
gitstart -d complete-test \
    -l javascript \
    -p \
    -b main \
    -m "Production ready v1.0" \
    -desc "A complete test of all features"
# Expected: Creates private repo with all specified options
# Verify all aspects on GitHub
```

#### Test 8.2: Multiple Languages
```bash
# Test different language .gitignores
for lang in python java go rust javascript; do
    gitstart -d test-${lang} -l ${lang}
    # Verify each .gitignore is language-specific
done
```

### 9. Edge Cases

#### Test 9.1: Very Long Repository Name
```bash
gitstart -d this-is-a-very-long-repository-name-to-test-limits
# Expected: Should work if under GitHub's limit (100 chars)
```

#### Test 9.2: Repository Name with Hyphens
```bash
gitstart -d test-with-many-hyphens-in-name
# Expected: Works normally
```

#### Test 9.3: Special Characters in Description
```bash
gitstart -d test-special -desc "Test with 'quotes' and \"double quotes\""
# Expected: Handles special characters correctly
```

#### Test 9.4: Empty Commit Message
```bash
gitstart -d test-empty-msg -m ""
# Expected: Should use empty string or default
```

### 10. Cleanup Tests

#### Test 10.1: Failed Repository Creation
```bash
# Create repo manually first
gh repo create test-cleanup --public
# Then try to create again
gitstart -d test-cleanup
# Expected: Should fail but cleanup the remote repo
# Verify: Repo should not exist after error
```

## Automated Test Script

Save this as `run-tests.sh`:

```bash
#!/bin/bash

TESTS_PASSED=0
TESTS_FAILED=0

run_test() {
    local test_name="$1"
    local test_command="$2"
    
    echo "Running: $test_name"
    if eval "$test_command"; then
        echo "✓ PASSED: $test_name"
        ((TESTS_PASSED++))
    else
        echo "✗ FAILED: $test_name"
        ((TESTS_FAILED++))
    fi
    echo ""
}

# Basic tests
run_test "Version Check" "gitstart -v"
run_test "Help Display" "gitstart -h"
run_test "Dry Run" "gitstart -d test-dry --dry-run"

echo "================================"
echo "Tests Passed: $TESTS_PASSED"
echo "Tests Failed: $TESTS_FAILED"
echo "================================"

# Cleanup
read -p "Delete test repositories from GitHub? (y/n): " cleanup
if [[ $cleanup == "y" ]]; then
    # Add cleanup commands
    echo "Cleaning up..."
fi
```

## Manual Verification Checklist

After running tests, verify:

- [ ] Repositories visible on GitHub
- [ ] Correct visibility (public/private)
- [ ] Files present in repository
- [ ] Commit messages correct
- [ ] Branch names correct
- [ ] .gitignore appropriate for language
- [ ] LICENSE file correct
- [ ] README.md generated
- [ ] Repository description set
- [ ] Remote URL correct

## Cleanup After Testing

```bash
# List all test repositories
gh repo list | grep test-

# Delete individual repo
gh repo delete USERNAME/REPO_NAME --yes

# Delete all test repos (careful!)
gh repo list --limit 1000 | grep "test-" | awk '{print $1}' | xargs -I {} gh repo delete {} --yes

# Clean local directories
cd ~/gitstart-tests
rm -rf test-*

# Reset configuration if needed
rm -rf ~/.config/gitstart
```

## Reporting Issues

When reporting bugs, include:
1. Test case that failed
2. Expected behavior
3. Actual behavior
4. Error messages
5. Output of `gitstart -v`
6. Output of `gh --version`
7. Operating system and version

## Continuous Testing

For ongoing development:
1. Run dry-run tests before each commit
2. Test with real repositories weekly
3. Verify on both macOS and Linux
4. Test with different GitHub accounts
5. Keep test repositories for regression testing
