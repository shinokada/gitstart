# Migration Guide: v0.3.0 to v0.4.0

This guide helps users upgrade from gitstart v0.3.0 to v0.4.0.

## Breaking Changes

### Configuration File Location

**Old location:** `~/.gitstart_config`  
**New location:** `~/.config/gitstart/config`

The configuration file has been moved to follow XDG Base Directory standards.

#### Migration Steps

1. The script will automatically prompt for your GitHub username on first run
2. If you want to migrate your existing configuration:

```bash
# Create new config directory
mkdir -p ~/.config/gitstart

# Copy old config to new location
if [ -f ~/.gitstart_config ]; then
    cp ~/.gitstart_config ~/.config/gitstart/config
    echo "Configuration migrated successfully"
fi

# Optionally remove old config
rm ~/.gitstart_config
```

## New Features

### 1. Private Repository Support

You can now create private repositories:

```bash
# Old way (always public)
gitstart -d my-repo

# New way (private)
gitstart -d my-repo -p
# or
gitstart -d my-repo --private
```

### 2. Custom Commit Messages

Customize your initial commit message:

```bash
# Old way (hardcoded "first commit")
gitstart -d my-repo

# New way (custom message)
gitstart -d my-repo -m "Initial release v1.0"
```

### 3. Custom Branch Names

Choose your default branch name:

```bash
# Old way (always "main")
gitstart -d my-repo

# New way (custom branch)
gitstart -d my-repo -b develop
gitstart -d my-repo -b master
```

### 4. Repository Description

Add a description when creating the repository:

```bash
gitstart -d my-tool -desc "A powerful CLI tool for developers"
```

### 5. Dry Run Mode

Preview what the script will do without making changes:

```bash
gitstart -d my-repo --dry-run
```

This is useful for:
- Testing configurations
- Understanding what files will be created
- Checking if a directory is suitable for initialization

### 6. Quiet Mode

For use in scripts or when you want minimal output:

```bash
gitstart -d my-repo -q
# or
gitstart -d my-repo --quiet
```

### 7. Better Handling of Existing Directories

The script now properly handles existing directories:

**Before v0.4.0:**
- Would fail or overwrite files unexpectedly
- Could cause data loss

**In v0.4.0:**
- Detects existing files and prompts for confirmation
- Detects existing git repositories
- Offers to append to existing .gitignore
- Skips creating LICENSE or README if they exist
- Preserves existing files in commits

Example workflow:
```bash
cd existing-project
gitstart -d . -l python

# You'll be prompted:
# "WARNING: Directory already contains files"
# "Continue anyway? This may overwrite existing files. (y/n):"
```

## Improved Error Handling

### Automatic Cleanup

If something goes wrong during repository creation, v0.4.0 automatically cleans up:

```bash
gitstart -d my-repo

# If creation fails after remote repo is created:
# ">>> Error occurred. Cleaning up remote repository..."
# The remote repo is automatically deleted
```

### Better Error Messages

**Before:**
```
Not able to create a remote repo.
```

**After:**
```
ERROR: Failed to create remote repository
Please check:
  1. Repository name 'my-repo' may already exist
  2. GitHub CLI authentication status
  3. Network connectivity
```

## Updated Command-Line Interface

### New Help Output

Run `gitstart -h` to see the comprehensive help:

```
Options:
========

    -d, --dir DIRECTORY          Directory name or path (use . for current directory)
    -l, --lang LANGUAGE          Programming language for .gitignore
    -p, --private                Create a private repository (default: public)
    -b, --branch BRANCH          Branch name (default: main)
    -m, --message MESSAGE        Initial commit message (default: "Initial commit")
    -desc, --description DESC    Repository description
    --dry-run                    Show what would happen without executing
    -q, --quiet                  Minimal output
    -h, --help                   Show this help message
    -v, --version                Show version
```

## Recommendations

### For Scripts and Automation

If you're using gitstart in scripts, consider:

1. **Use quiet mode** to reduce output:
   ```bash
   gitstart -d "$repo_name" -q
   ```

2. **Use dry-run first** to validate:
   ```bash
   if gitstart -d "$repo_name" --dry-run; then
       gitstart -d "$repo_name"
   fi
   ```

3. **Set all options explicitly**:
   ```bash
   gitstart -d my-repo \
       -l python \
       -p \
       -b main \
       -m "Automated initialization" \
       -desc "Auto-generated repository"
   ```

### For Interactive Use

1. **Try dry-run first** with new configurations
2. **Use descriptive commit messages** for better git history
3. **Add descriptions** to make repositories more discoverable
4. **Review prompts carefully** when working with existing directories

## Troubleshooting

### "Repository already exists" Error

If you get this error, the repository may already exist on GitHub:

```bash
# Check if repo exists
gh repo view username/repo-name

# Delete if needed
gh repo delete username/repo-name

# Try again
gitstart -d repo-name
```

### Configuration Not Found

If the script asks for your username again:

```bash
# Check config location
cat ~/.config/gitstart/config

# Manually set it
echo "your-github-username" > ~/.config/gitstart/config
```

### Existing Directory Issues

If working with an existing directory:

```bash
# Use dry-run to see what will happen
gitstart -d . --dry-run

# Review the output before proceeding
gitstart -d .
```

## Getting Help

If you encounter issues:

1. Check the [README](README.md) for updated documentation
2. Run `gitstart -h` for command help
3. Try `--dry-run` to understand what the script will do
4. Open an issue on GitHub with details about your problem

## Feedback

We welcome feedback on the new features! Please:
- Report bugs as GitHub issues
- Suggest improvements
- Share your use cases

Thank you for using gitstart!
