# Gitstart Quick Reference

A one-page reference for gitstart commands and options.

## Installation

```bash
# Homebrew (macOS)
brew tap shinokada/gitstart && brew install gitstart

# Awesome package manager
awesome install shinokada/gitstart

# Debian/Ubuntu
sudo apt install ./gitstart_x.x.x_all.deb
```

## Basic Syntax

```bash
gitstart [OPTIONS]
```

## Required Option

| Option | Description | Example |
|--------|-------------|---------|
| `-d, --dir DIR` | Directory path (use `.` for current) | `gitstart -d myapp` |

## Optional Flags

| Flag | Description | Default | Example |
|------|-------------|---------|---------|
| `-l, --lang LANG` | Language for .gitignore | none | `-l python` |
| `-p, --private` | Create private repo | public | `-p` |
| `-b, --branch NAME` | Branch name | main | `-b develop` |
| `-m, --message MSG` | Commit message | "Initial commit" | `-m "v1.0"` |
| `-desc, --description` | Repo description | none | `-desc "My app"` |
| `--dry-run` | Preview without executing | - | `--dry-run` |
| `-q, --quiet` | Minimal output | - | `-q` |
| `-h, --help` | Show help | - | `-h` |
| `-v, --version` | Show version | - | `-v` |

## Common Commands

```bash
# Create new public repo
gitstart -d myapp

# Create with Python .gitignore
gitstart -d myapp -l python

# Create private repo
gitstart -d myapp -p

# Use current directory
cd existing-project
gitstart -d .

# Preview before creating
gitstart -d myapp --dry-run

# Full configuration
gitstart -d myapp \
  -l python \
  -p \
  -b develop \
  -m "Initial release" \
  -desc "My Python application"
```

## Supported Languages

Python • Java • JavaScript • Node • Go • Rust • Ruby • PHP • C • C++ • C# • Swift • Kotlin • TypeScript • Scala • Perl • R • Dart • Elixir • Haskell • Julia • Lua • and more...

See: https://github.com/github/gitignore

## License Options

When prompted, choose:
1. MIT - Simple and permissive
2. Apache 2.0 - Community-focused
3. GNU GPLv3 - Share improvements
4. None - No license file

## Configuration

- **Location**: `~/.config/gitstart/config`
- **Content**: Your GitHub username
- **Edit**: `nano ~/.config/gitstart/config`

## Workflow Examples

### New Project
```bash
gitstart -d new-project -l python
cd new-project
# Start coding!
```

### Existing Project
```bash
cd existing-project
gitstart -d . -l javascript
# Your files are preserved and committed
```

### Team Project
```bash
gitstart -d team-app \
  -l go \
  -b develop \
  -desc "Team collaboration project"
```

### Private Experiment
```bash
gitstart -d experiment -p --dry-run
# Review output
gitstart -d experiment -p
```

## Directory Behaviors

| Situation | Behavior |
|-----------|----------|
| New directory | Creates directory and repo |
| Empty existing directory | Creates repo normally |
| Directory with files | Warns, asks confirmation, preserves files |
| Directory with .git | Adds remote, preserves history |
| Existing LICENSE | Skips creating LICENSE |
| Existing README.md | Skips creating README |
| Existing .gitignore | Offers to append |

## Error Handling

| Issue | What Happens |
|-------|-------------|
| Repo creation fails | Auto-cleanup of remote repo |
| Not logged in to GitHub | Error with instructions |
| Missing dependencies | Lists what to install |
| Repository already exists | Error message |
| Network issues | Descriptive error |

## Prerequisites

- ✓ GitHub CLI (`gh`) installed and authenticated
- ✓ `jq` installed
- ✓ Bash shell
- ✓ Git installed

```bash
# Check prerequisites
gh auth status
command -v jq
command -v git
```

## Tips & Tricks

### Always Test First
```bash
gitstart -d myapp --dry-run
```

### Create Shell Function
```bash
# Add to ~/.bashrc or ~/.zshrc
newrepo() {
  gitstart -d "$1" -l "${2:-python}" -b develop
}

# Use it
newrepo myapp javascript
```

### Batch Creation
```bash
for repo in api frontend mobile; do
  gitstart -d "$repo" -q
done
```

### With Description
```bash
gitstart -d myapp -desc "$(cat description.txt)"
```

## Troubleshooting

| Problem | Solution |
|---------|----------|
| "Not logged in" | Run `gh auth login` |
| "jq not found" | Install jq: `brew install jq` or `apt install jq` |
| "Repo already exists" | Use different name or delete existing repo |
| Config not found | Will prompt for username on first run |
| Permission denied | Check GitHub token permissions |

## Getting Help

```bash
# Built-in help
gitstart -h

# Version
gitstart -v

# Check GitHub auth
gh auth status

# View created repo
gh repo view username/reponame
```

## Quick Uninstall

```bash
curl -s https://raw.githubusercontent.com/shinokada/gitstart/main/uninstall.sh | bash
```

## Resources

- **GitHub**: https://github.com/shinokada/gitstart
- **Documentation**: See README.md
- **Examples**: See EXAMPLES.md
- **Issues**: GitHub Issues
- **Author**: Shinichi Okada (@shinokada)

## Version

Current: **v0.4.0** (January 2026)

---

**Print this page for quick reference!**
