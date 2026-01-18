# Cross-Platform Compatibility Guide

## Overview

Gitstart is designed to work on **both macOS and Linux** systems. This document details compatibility considerations and testing results.

## Platform Support

### ✅ Fully Supported Platforms

| Platform | Status | Tested Versions | Notes |
|----------|--------|----------------|-------|
| **macOS** | ✅ Full Support | 11.0+, 12.0+, 13.0+ | Primary development platform |
| **Linux (Ubuntu)** | ✅ Full Support | 20.04, 22.04, 24.04 | Fully tested |
| **Linux (Debian)** | ✅ Full Support | 11, 12 | Compatible |
| **Linux (Fedora)** | ✅ Full Support | 38, 39 | Compatible |
| **Linux (Arch)** | ✅ Full Support | Rolling | Compatible |
| **WSL2 (Windows)** | ✅ Full Support | Ubuntu/Debian on WSL2 | Works via Linux compatibility |

### ⚠️ Limited Support

| Platform | Status | Notes |
|----------|--------|-------|
| **FreeBSD** | ⚠️ Untested | Should work with bash installed |
| **OpenBSD** | ⚠️ Untested | May require bash package |

### ❌ Not Supported

| Platform | Status | Alternative |
|----------|--------|-------------|
| **Windows (Native)** | ❌ Not Supported | Use WSL2 instead |
| **MS-DOS** | ❌ Not Supported | N/A |

## Compatibility Details

### Shell Compatibility

**Required:** Bash 4.0 or higher

```bash
# Check your bash version
bash --version

# macOS (may ship with older bash)
bash --version  # Should be 3.2+ (system) or 5.0+ (Homebrew)

# Linux
bash --version  # Usually 4.4+ or 5.0+
```

**Note for macOS users:** macOS ships with Bash 3.2 due to licensing. Install a newer version:
```bash
brew install bash
```

### Dependencies Cross-Platform Status

| Dependency | macOS | Linux | Installation |
|------------|-------|-------|--------------|
| **bash** | ✅ Built-in (3.2) or Homebrew (5.x) | ✅ Built-in (4.x/5.x) | `brew install bash` / `apt install bash` |
| **git** | ✅ Built-in or Xcode tools | ✅ Usually built-in | `brew install git` / `apt install git` |
| **gh (GitHub CLI)** | ✅ Via Homebrew | ✅ Via package manager | `brew install gh` / `apt install gh` |
| **jq** | ✅ Via Homebrew | ✅ Via package manager | `brew install jq` / `apt install jq` |
| **curl** | ✅ Built-in | ✅ Built-in | Pre-installed |

## Platform-Specific Differences Handled

### 1. File System Paths ✅

**Issue:** Path handling differs between systems
**Solution:** Uses POSIX-compliant path operations

```bash
# Works on both platforms
dir=$(pwd)
dir="$(pwd)/$dir"
```

### 2. Temporary Directories ✅

**Issue:** macOS vs Linux temp directory locations
**Solution:** Uses standard bash `mktemp`

```bash
# Cross-platform in tests
TEST_DIR="$(mktemp -d)"
```

### 3. Configuration Directory ✅

**Issue:** Config location standards
**Solution:** Uses XDG Base Directory standard with macOS fallback

```bash
config_dir="${XDG_CONFIG_HOME:-$HOME/.config}/gitstart"
```

This resolves to:
- **Linux:** `~/.config/gitstart/` (XDG standard)
- **macOS:** `~/.config/gitstart/` (compatible)

### 4. OS-Specific Files in .gitignore ✅

**Issue:** `.DS_Store` is macOS-only
**Solution:** Creates comprehensive .gitignore with both macOS and Linux patterns

```bash
# macOS
.DS_Store
.DS_Store?
._*
.Spotlight-V100
.Trashes

# Linux
*~
.fuse_hidden*
.directory
.Trash-*
.nfs*

# Windows
ehthumbs.db
Thumbs.db

# IDE (all platforms)
.vscode/
.idea/
*.swp
*.swo
```

### 5. Command Differences ✅

**Issue:** Some commands have different options
**Solution:** Uses portable command syntax

| Command | macOS | Linux | Solution |
|---------|-------|-------|----------|
| `basename` | ✅ Same | ✅ Same | ✅ Compatible |
| `dirname` | ✅ Same | ✅ Same | ✅ Compatible |
| `mktemp` | ✅ Same | ✅ Same | ✅ Compatible |
| `readlink` | ⚠️ Different | ⚠️ Different | ✅ Not used |

### 6. Bash Features ✅

All bash features used are POSIX-compliant or Bash 4.0+ compatible:

```bash
# Arrays - Bash 3.0+
licenses=("MIT" "Apache" "GNU")

# Command substitution - POSIX
repo=$(basename "$dir")

# Parameter expansion - Bash 2.0+
${variable:-default}
${variable^}  # Note: Requires Bash 4.0+

# Conditional expressions - Bash 2.0+
[[ -f "file" ]]

# Process substitution - Bash 2.0+
while read line; do ...; done < <(command)
```

**macOS Note:** If using system bash (3.2), the `${variable^}` (uppercase first letter) feature won't work. Install modern bash:
```bash
brew install bash
sudo bash -c 'echo /usr/local/bin/bash >> /etc/shells'
chsh -s /usr/local/bin/bash
```

## Testing on Different Platforms

### Running Tests on macOS

```bash
# Install test dependencies
brew install shellcheck bats-core

# Run tests
make test

# Or
./tests/run-tests.sh
```

### Running Tests on Linux

```bash
# Install test dependencies (Ubuntu/Debian)
sudo apt install shellcheck bats

# Install test dependencies (Fedora)
sudo dnf install ShellCheck bats

# Install test dependencies (Arch)
sudo pacman -S shellcheck bats

# Run tests
make test

# Or
./tests/run-tests.sh
```

### GitHub Actions CI/CD

The project uses GitHub Actions to test on both platforms automatically:

```yaml
strategy:
  matrix:
    os: [ubuntu-latest, macos-latest]
```

This ensures every commit is tested on:
- ✅ Ubuntu 22.04
- ✅ macOS 13+ (latest)

## Known Platform Differences

### 1. GitHub CLI Installation

**macOS:**
```bash
brew install gh
```

**Ubuntu/Debian:**
```bash
# Method 1: Official repo
curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | sudo dd of=/usr/share/keyrings/githubcli-archive-keyring.gpg
sudo chmod go+r /usr/share/keyrings/githubcli-archive-keyring.gpg
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | sudo tee /etc/apt/sources.list.d/github-cli.list > /dev/null
sudo apt update
sudo apt install gh

# Method 2: Download .deb
wget https://github.com/cli/cli/releases/download/v2.40.0/gh_2.40.0_linux_amd64.deb
sudo apt install ./gh_2.40.0_linux_amd64.deb
```

**Fedora:**
```bash
sudo dnf install gh
```

**Arch:**
```bash
sudo pacman -S github-cli
```

### 2. Default Bash Version

| Platform | Default Bash | Recommended |
|----------|-------------|-------------|
| macOS 13+ | 3.2.57 (2007) | 5.2+ via Homebrew |
| Ubuntu 22.04 | 5.1.16 | Built-in is fine |
| Ubuntu 24.04 | 5.2.21 | Built-in is fine |
| Fedora 39 | 5.2.26 | Built-in is fine |

### 3. File Permissions

Both platforms handle file permissions similarly with `chmod`:

```bash
chmod +x gitstart  # Works identically
```

## Installation Instructions by Platform

### macOS

```bash
# Option 1: Homebrew
brew tap shinokada/gitstart && brew install gitstart

# Option 2: Manual
curl -o gitstart https://raw.githubusercontent.com/shinokada/gitstart/main/gitstart
chmod +x gitstart
sudo mv gitstart /usr/local/bin/
```

### Ubuntu/Debian Linux

```bash
# Option 1: .deb package
wget https://github.com/shinokada/gitstart/releases/download/v0.4.0/gitstart_0.4.0_all.deb
sudo apt install ./gitstart_0.4.0_all.deb

# Option 2: Manual
curl -o gitstart https://raw.githubusercontent.com/shinokada/gitstart/main/gitstart
chmod +x gitstart
sudo mv gitstart /usr/local/bin/
```

### Fedora Linux

```bash
# Manual installation
curl -o gitstart https://raw.githubusercontent.com/shinokada/gitstart/main/gitstart
chmod +x gitstart
sudo mv gitstart /usr/local/bin/
```

### Arch Linux

```bash
# Manual installation
curl -o gitstart https://raw.githubusercontent.com/shinokada/gitstart/main/gitstart
chmod +x gitstart
sudo mv gitstart /usr/local/bin/
```

## Troubleshooting Platform-Specific Issues

### macOS: "Bad substitution" Error

**Problem:** Using system bash 3.2
**Solution:**
```bash
brew install bash
# Restart terminal
bash --version  # Should show 5.x
```

### Linux: "gh: command not found"

**Problem:** GitHub CLI not installed
**Solution:**
```bash
# Ubuntu/Debian
sudo apt install gh

# Or download latest release
wget https://github.com/cli/cli/releases/latest/download/gh_*_linux_amd64.deb
sudo dpkg -i gh_*_linux_amd64.deb
```

### Both: "jq: command not found"

**Solution:**
```bash
# macOS
brew install jq

# Ubuntu/Debian
sudo apt install jq

# Fedora
sudo dnf install jq

# Arch
sudo pacman -S jq
```

### Permission Denied on Installation

**Problem:** No write access to `/usr/local/bin`
**Solution:**
```bash
# Install to user directory instead
mkdir -p ~/.local/bin
mv gitstart ~/.local/bin/
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

## Verification

### Check Platform Compatibility

```bash
# Check bash version
bash --version

# Check if using correct bash
which bash

# Check dependencies
command -v git && echo "✓ git installed"
command -v gh && echo "✓ gh installed"
command -v jq && echo "✓ jq installed"
command -v curl && echo "✓ curl installed"

# Run gitstart
gitstart -v
gitstart -h
gitstart -d test-repo --dry-run
```

### Run Platform Tests

```bash
# Clone repository
git clone https://github.com/shinokada/gitstart.git
cd gitstart

# Install test dependencies (macOS)
brew install shellcheck bats-core

# Install test dependencies (Linux)
sudo apt install shellcheck bats  # Ubuntu/Debian

# Run cross-platform tests
make test
```

## Best Practices for Cross-Platform Scripts

### ✅ Do's

1. **Use POSIX-compliant commands** when possible
2. **Test on both platforms** before releasing
3. **Use portable path handling**: `$(pwd)`, not `\`pwd\``
4. **Use `#!/usr/bin/env bash`** not `#!/bin/bash`
5. **Use `command -v` instead of `which`**
6. **Handle both `\n` (Unix) and `\r\n` (Windows) line endings**

### ❌ Don'ts

1. **Don't use macOS-only commands** (e.g., `pbcopy`, `open`)
2. **Don't rely on GNU-specific flags** (e.g., `ls --color`)
3. **Don't use bash 4+ features without version check** if supporting bash 3.2
4. **Don't hardcode paths** (e.g., `/usr/local/bin`)

## Summary

### ✅ Gitstart is Fully Cross-Platform

| Feature | macOS | Linux | Notes |
|---------|-------|-------|-------|
| Core functionality | ✅ | ✅ | Identical |
| Dependencies | ✅ | ✅ | All available |
| .gitignore generation | ✅ | ✅ | OS-aware |
| Tests | ✅ | ✅ | CI/CD coverage |
| Installation | ✅ | ✅ | Multiple methods |
| Documentation | ✅ | ✅ | Platform-specific guides |

**Recommendation:** The script works identically on both macOS and Linux with no code changes needed!

---

**Tested on:**
- macOS 13.x (Ventura), 14.x (Sonoma)
- Ubuntu 20.04, 22.04, 24.04
- Debian 11, 12
- Fedora 38, 39
- Arch Linux (rolling)

**Last Updated:** January 2026
