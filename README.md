# Gitstart

> Automate GitHub repository creation with one command

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Version](https://img.shields.io/badge/version-0.4.0-blue.svg)](https://github.com/shinokada/gitstart/releases)
[![Tests](https://img.shields.io/badge/tests-passing-brightgreen.svg)](https://github.com/shinokada/gitstart/actions)
[![ShellCheck](https://img.shields.io/badge/shellcheck-passing-brightgreen.svg)](https://github.com/koalaman/shellcheck)

## What is Gitstart?

Gitstart is a powerful Bash script that automates the entire process of creating a new Git repository and pushing it to GitHub. Instead of running multiple commands, you can initialize a complete project with proper structure in one line.

**Before Gitstart:**
```bash
mkdir my-project
cd my-project
git init
touch README.md LICENSE .gitignore
# ... manually create files ...
git add .
git commit -m "Initial commit"
gh repo create my-project --public
git remote add origin git@github.com:username/my-project.git
git push -u origin main
```

**With Gitstart:**
```bash
gitstart -d my-project -l python
```

## Features

✨ **Complete Automation**
- Creates directory structure
- Initializes Git repository
- Creates GitHub repository
- Generates LICENSE, README.md, .gitignore
- Commits and pushes to GitHub

🎯 **Smart File Generation**
- Language-specific .gitignore files (Python, JavaScript, Go, Rust, etc.)
- Multiple license options (MIT, Apache 2.0, GNU GPLv3)
- Professional README.md template

🔧 **Flexible Configuration**
- Public or private repositories
- Custom branch names
- Custom commit messages
- Repository descriptions
- Dry-run mode to preview changes

🛡️ **Safe and Reliable**
- Detects existing files and prompts before overwriting
- Works with existing Git repositories
- Automatic cleanup on errors
- Comprehensive error handling

🧪 **Well Tested**
- Automated test suite with shellcheck and bats
- CI/CD pipeline with GitHub Actions
- Unit and integration tests
- Cross-platform compatibility (Linux, macOS)

🌍 **Cross-Platform**
- Works on macOS, Linux (Ubuntu, Debian, Fedora, Arch)
- WSL2 support for Windows users
- Tested on both platforms via CI/CD
- OS-aware .gitignore generation

## Quick Start

```bash
# Install (macOS)
brew tap shinokada/gitstart && brew install gitstart

# Login to GitHub
gh auth login

# Create your first repository
gitstart -d my-awesome-project

# That's it! Your project is now on GitHub
```

## Installation

### macOS (Homebrew)
```bash
brew tap shinokada/gitstart && brew install gitstart
```

### Linux (Awesome Package Manager)
```bash
awesome install shinokada/gitstart
```

### Debian/Ubuntu
```bash
# Download from releases page
wget https://github.com/shinokada/gitstart/releases/download/v0.4.0/gitstart_0.4.0_all.deb
sudo apt install ./gitstart_0.4.0_all.deb
```

### Fedora/RHEL/CentOS
```bash
# Manual installation (no rpm package yet)
curl -o gitstart https://raw.githubusercontent.com/shinokada/gitstart/main/gitstart
chmod +x gitstart
sudo mv gitstart /usr/local/bin/
```

### Arch Linux
```bash
# Manual installation (AUR package coming soon)
curl -o gitstart https://raw.githubusercontent.com/shinokada/gitstart/main/gitstart
chmod +x gitstart
sudo mv gitstart /usr/local/bin/
```

### WSL2 (Windows)
```bash
# Use Ubuntu/Debian instructions
wget https://github.com/shinokada/gitstart/releases/download/v0.4.0/gitstart_0.4.0_all.deb
sudo apt install ./gitstart_0.4.0_all.deb
```

### Manual Installation (All Platforms)
```bash
curl -o gitstart https://raw.githubusercontent.com/shinokada/gitstart/main/gitstart
chmod +x gitstart
sudo mv gitstart /usr/local/bin/

# Or install to user directory (no sudo needed)
mkdir -p ~/.local/bin
mv gitstart ~/.local/bin/
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

## Requirements

- **GitHub CLI** (`gh`) - [Installation guide](https://cli.github.com/manual/installation)
- **jq** - JSON processor
- **Git** - Version control
- **Bash** - Shell (pre-installed on macOS/Linux)

```bash
# macOS
brew install gh jq

# Ubuntu/Debian
sudo apt install gh jq

# Fedora
sudo dnf install gh jq

# Arch Linux
sudo pacman -S github-cli jq

# Verify installation
gh --version
jq --version
```

## Usage

### Basic Usage

```bash
# Create a new repository
gitstart -d project-name

# With programming language
gitstart -d my-app -l python

# Private repository
gitstart -d secret-project -p

# In current directory
cd existing-project
gitstart -d .
```

### All Options

```bash
gitstart [OPTIONS]

Options:
  -d, --dir DIR              Directory name or path (required)
  -l, --lang LANGUAGE        Programming language for .gitignore
  -p, --private              Create private repository (default: public)
  -b, --branch NAME          Branch name (default: main)
  -m, --message MESSAGE      Initial commit message (default: "Initial commit")
  -desc, --description TEXT  Repository description
  --dry-run                  Preview without executing
  -q, --quiet                Minimal output
  -h, --help                 Show help
  -v, --version              Show version
```

### Examples

**Python Project**
```bash
gitstart -d my-python-app -l python
```

**Private React App**
```bash
gitstart -d react-app -l javascript -p -desc "My React application"
```

**Custom Everything**
```bash
gitstart -d full-config \
  -l go \
  -p \
  -b develop \
  -m "Project initialization v1.0" \
  -desc "A fully configured Go project"
```

**Preview Before Creating**
```bash
gitstart -d test-project --dry-run
```

## Supported Languages

Python • JavaScript • Node • TypeScript • Go • Rust • Java • C • C++ • C# • Ruby • PHP • Swift • Kotlin • Scala • Dart • Elixir • Haskell • Perl • R • Lua • and more!

Full list: https://github.com/github/gitignore

## Documentation

📖 **Comprehensive Guides:**
- [Quick Reference](QUICK_REFERENCE.md) - One-page command reference
- [Examples](EXAMPLES.md) - Real-world usage examples
- [Testing Guide](TESTING.md) - Test procedures and validation
- [Migration Guide](MIGRATION.md) - Upgrade from v0.3.0
- [Changelog](CHANGELOG.md) - Version history

## What's New in v0.4.0

🎉 **Major Update!**

- ✅ Private repository support
- ✅ Custom commit messages
- ✅ Custom branch names
- ✅ Repository descriptions
- ✅ Dry-run mode
- ✅ Quiet mode for automation
- ✅ Full support for existing directories
- ✅ Automatic error cleanup
- ✅ Better file conflict handling
- ✅ XDG-compliant config location

See [CHANGELOG.md](CHANGELOG.md) for complete details.

## Common Use Cases

### 1. Quick Project Setup
```bash
gitstart -d new-idea -l python
cd new-idea
# Start coding immediately!
```

### 2. Team Project
```bash
gitstart -d team-project \
  -l javascript \
  -b develop \
  -desc "Team collaboration project"
```

### 3. Initialize Existing Project
```bash
cd ~/Downloads/client-project
gitstart -d . -l go
```

### 4. Microservices
```bash
for service in user-svc payment-svc notification-svc; do
  gitstart -d "$service" -l go -p -q
done
```

### 5. Automated Scripts
```bash
#!/bin/bash
create_repo() {
  gitstart -d "$1" -l "${2:-python}" -q
}

create_repo "api-service" "go"
create_repo "web-frontend" "javascript"
```

## How It Works

1. **Validates** - Checks GitHub authentication and dependencies
2. **Creates** - Makes directory (if needed) and GitHub repository
3. **Initializes** - Sets up Git with proper configuration
4. **Generates** - Creates LICENSE, README.md, and .gitignore
5. **Commits** - Stages and commits all files
6. **Pushes** - Uploads to GitHub with proper remote setup

## Configuration

Gitstart stores your GitHub username in:
```
~/.config/gitstart/config
```

On first run, you'll be prompted to enter it. This follows XDG Base Directory standards.

## Troubleshooting

**"Not logged in to GitHub"**
```bash
gh auth login
```

**"jq not found"**
```bash
# macOS
brew install jq

# Ubuntu/Debian
sudo apt install jq
```

**"Repository already exists"**
```bash
# Delete existing repo
gh repo delete username/repo-name --yes

# Or use different name
gitstart -d repo-name-v2
```

**Config not found**
- Will prompt for username on first run
- Manually set: `echo "your-username" > ~/.config/gitstart/config`

## Uninstall

```bash
curl -s https://raw.githubusercontent.com/shinokada/gitstart/main/uninstall.sh | bash
```

Or manually:
```bash
# Remove script
sudo rm /usr/local/bin/gitstart  # or your installation path

# Remove config
rm -rf ~/.config/gitstart
```

## Testing

Gitstart has a comprehensive test suite using **shellcheck** (static analysis) and **bats** (functional testing).

### Quick Testing

```bash
# Run all tests
make test

# Or use the test runner directly
./tests/run-tests.sh

# Run specific tests
make test-shellcheck    # Static analysis only
make test-unit          # Unit tests only
```

### Test Dependencies

```bash
# Install test dependencies
make install-deps

# Or manually:
brew install shellcheck bats-core  # macOS
sudo apt install shellcheck bats   # Ubuntu/Debian
```

### Running Tests

**All tests:**
```bash
./tests/run-tests.sh
```

**Shellcheck only:**
```bash
./tests/shellcheck.sh
```

**BATS unit tests:**
```bash
bats tests/gitstart.bats
```

**Specific test:**
```bash
bats tests/gitstart.bats --filter "version"
```

### CI/CD

Tests automatically run on push/PR via GitHub Actions:
- ✅ ShellCheck static analysis
- ✅ Unit tests on Ubuntu and macOS
- ✅ Bash compatibility checks
- ✅ Security scanning

See [tests/README.md](tests/README.md) for detailed testing documentation.

## Contributing

Contributions are welcome! Please:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. **Run tests:** `make test` or `./tests/run-tests.sh`
5. Ensure all tests pass
6. Submit a pull request

For detailed testing guidelines, see [tests/README.md](tests/README.md).

## Support

- 🐛 **Bug Reports**: [GitHub Issues](https://github.com/shinokada/gitstart/issues)
- 💬 **Questions**: [GitHub Discussions](https://github.com/shinokada/gitstart/discussions)
- 📧 **Contact**: [@shinokada](https://twitter.com/shinokada)

## License

MIT License - see [LICENSE](License) file for details.

Copyright (c) 2021-2026 Shinichi Okada

## Author

**Shinichi Okada** ([@shinokada](https://github.com/shinokada))

- [Medium](https://shinichiokada.medium.com/)
- [Twitter](https://twitter.com/shinokada)
- [More about Gitstart](https://towardsdatascience.com/automate-creating-a-new-github-repository-with-gitstart-1ae961b99866)

## Acknowledgments

Thanks to all users who have provided feedback and contributed to making Gitstart better!

---

<p align="center">
  <strong>Star ⭐ the repo if you find it useful!</strong>
</p>

<p align="center">
  <a href='https://ko-fi.com/Z8Z2CHALG' target='_blank'>
    <img height='42' src='https://storage.ko-fi.com/cdn/kofi3.png?v=3' alt='Buy Me a Coffee' />
  </a>
</p>
