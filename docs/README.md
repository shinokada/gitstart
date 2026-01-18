<p align="center">
<img width="400" src="https://raw.githubusercontent.com/shinokada/gitstart/main/images/gitstart.png" />
</p>

<p align="center">
<a href='https://ko-fi.com/Z8Z2CHALG' target='_blank'><img height='42' style='border:0px;height:42px;' src='https://storage.ko-fi.com/cdn/kofi3.png?v=3' alt='Buy Me a Coffee at ko-fi.com' /></a>
</p>

<h1  align="center">Gitstart</h1>

[More details about Gitstart.](https://towardsdatascience.com/automate-creating-a-new-github-repository-with-gitstart-1ae961b99866)

## Overview

> Gitstart creates, adds, and pushes with one line.

This script automates creating a git repository. The script will:

- Create .gitignore if you provide a language
- Create a license file based on your choice
- Create a new repository at GitHub.com (public or private)
- Create a README.md file with the repository name
- Initialize git repository (if needed)
- Add README.md and commit with a custom message
- Add the remote and push the files
- Support for existing directories and projects

The script reads your GitHub username from configuration and uses the directory name as the GitHub repository name.

## Requirements

- UNIX-like system (Tested on Ubuntu and macOS)
- [GitHub CLI](https://cli.github.com/manual/), >v2.3.0
- [jq](https://stedolan.github.io/jq/)

Linux users can download gh CLI from the [Releases page](https://github.com/cli/cli/releases), then run:

```sh
sudo apt install ./gh_x.x.x_xxxxxxx.deb
```

## Installation

### Linux/macOS using Awesome

After installing [Awesome package manager](https://github.com/shinokada/awesome):

```sh
awesome install shinokada/gitstart
```

### macOS using Homebrew

If you have Homebrew on your macOS, you can run:

```sh
brew tap shinokada/gitstart && brew install gitstart
```

### Debian/Ubuntu

Download the latest version from [releases page](https://github.com/shinokada/gitstart/releases) and run:

```sh
sudo apt install ./gitstart_version_all.deb
```

## Uninstallation

If you installed Gitstart using Awesome package manager/Homebrew/Debian package, run:

```sh
curl -s https://raw.githubusercontent.com/shinokada/gitstart/main/uninstall.sh > tmp1 && bash tmp1 && rm tmp1
```

## Usage

### Basic Usage

Login to GitHub and start using gitstart:

```sh
# Login to GitHub
gh auth login

# Create a new repository
gitstart -d repo-name

# Create in current directory
cd existing_project
gitstart -d .
```

### Options

```
-d, --dir DIRECTORY          Directory name or path (use . for current directory)
-l, --lang LANGUAGE          Programming language for .gitignore
-p, --private                Create a private repository (default: public)
-b, --branch BRANCH          Branch name (default: main)
-m, --message MESSAGE        Initial commit message (default: "Initial commit")
-desc, --description DESC    Repository description
--dry-run                    Show what would happen without executing
-q, --quiet                  Minimal output
-h, --help                   Show help message
-v, --version                Show version
```

### Examples

**Create a new repository:**
```sh
gitstart -d my-project
```

**Create with specific programming language:**
```sh
gitstart -d my-python-app -l python
```

**Create a private repository:**
```sh
gitstart -d secret-project -p
```

**Use custom commit message and branch:**
```sh
gitstart -d my-app -m "First release" -b develop
```

**Add repository description:**
```sh
gitstart -d awesome-tool -desc "An amazing CLI tool for developers"
```

**Preview changes without executing (dry run):**
```sh
gitstart -d test-repo --dry-run
```

**Quiet mode for scripts:**
```sh
gitstart -d automated-repo -q
```

**Initialize existing project:**
```sh
cd my-existing-project
gitstart -d . -l javascript -desc "My existing JavaScript project"
```

### Working with Existing Directories

Gitstart now fully supports existing directories with the following behaviors:

**Empty directory:** Creates repository normally

**Directory with files but no git:** 
- Warns about existing files
- Asks for confirmation
- Preserves existing files
- Adds them to the initial commit

**Directory with existing git repository:**
- Detects existing `.git` folder
- Adds remote to existing repository
- Preserves git history

**Existing LICENSE, README.md, or .gitignore:**
- Detects existing files
- Offers to append or skip
- Prevents accidental overwrites

### Interactive License Selection

When you run gitstart, you'll be prompted to select a license:

```
Select a license:
1) MIT: I want it simple and permissive.
2) Apache License 2.0: I need to work in a community.
3) GNU GPLv3: I care about sharing improvements.
4) None
5) Quit
```

## Configuration

Gitstart stores your GitHub username in `~/.config/gitstart/config` (follows XDG standards). On first run, you'll be prompted to enter your username, which will be saved for future use.

## Error Handling

The script includes comprehensive error handling:

- **Automatic cleanup**: If repository creation fails, the remote repository is automatically deleted
- **Validation checks**: Ensures all required tools are installed
- **Auth verification**: Confirms you're logged in to GitHub
- **File conflict detection**: Warns about existing files before overwriting
- **Detailed error messages**: Clear information about what went wrong and how to fix it

## About Licensing

Read more about [Licensing](https://docs.github.com/en/free-pro-team@latest/rest/reference/licenses).

## Changelog

### Version 0.4.0 (2026-01-18)

**New Features:**
- Private repository support with `-p/--private` flag
- Custom commit messages with `-m/--message` flag
- Custom branch names with `-b/--branch` flag
- Repository description with `-desc/--description` flag
- Dry run mode with `--dry-run` flag
- Quiet mode with `-q/--quiet` flag
- Full support for existing directories and files
- Automatic rollback on errors
- Detection and handling of existing git repositories

**Improvements:**
- Fixed GitHub auth check (now uses proper exit code checking)
- XDG-compliant config directory (`~/.config/gitstart/config`)
- Better error messages with context
- File conflict detection and user prompts
- Smarter handling of existing LICENSE, README, and .gitignore files
- Improved overall code quality and error handling

**Bug Fixes:**
- Fixed issue with `gh repo create --clone` in existing directories
- Fixed auth status check that was comparing string to integer
- Proper handling of existing files to prevent data loss

### Version 0.3.0
- Initial public release

## Author

Shinichi Okada

- [Medium](https://shinichiokada.medium.com/)
- [Twitter](https://twitter.com/shinokada)

## License

Copyright (c) 2021-2026 Shinichi Okada (@shinokada)
This software is released under the MIT License, see LICENSE.
