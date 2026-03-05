<p align="center">
<img width="400" src="https://raw.githubusercontent.com/shinokada/gitstart/main/images/gitstart.png" />
</p>

<p align="center">
<a href='https://ko-fi.com/Z8Z2CHALG' target='_blank'><img height='42' style='border:0px;height:42px;' src='https://storage.ko-fi.com/cdn/kofi3.png?v=3' alt='Buy Me a Coffee at ko-fi.com' /></a>
</p>

<h1  align="center">Gitstart</h1>

## Overview

> Gitstart creates, adds, and pushes with one line.

Gitstart automates creating a GitHub repository. It will:

- Create `.gitignore` if you provide a language
- Create a license file based on your choice
- Create a new repository at GitHub.com (public or private)
- Create a `README.md` file with the repository name
- Initialize a git repository (if needed)
- Add files and commit with a custom message
- Add the remote and push
- Support existing directories and projects

## Requirements

- [GitHub CLI](https://cli.github.com/manual/) (`gh`), authenticated
- `git` installed
- macOS, Linux, or Windows

## Installation

### Homebrew (macOS/Linux)

```sh
brew install shinokada/gitstart/gitstart
```

### Windows Scoop

```sh
scoop bucket add shinokada https://github.com/shinokada/scoop-bucket
scoop install gitstart
```

### Debian/Ubuntu

Download the latest `.deb` from the [releases page](https://github.com/shinokada/gitstart/releases) and run:

```sh
sudo apt install ./gitstart_x.x.x_linux_amd64.deb
```

### Fedora/RHEL

Download the latest `.rpm` from the [releases page](https://github.com/shinokada/gitstart/releases) and run:

```sh
sudo rpm -i gitstart_x.x.x_linux_amd64.rpm
```

### Go install

```sh
go install github.com/shinokada/gitstart@latest
```

This places the `gitstart` binary in your `$GOPATH/bin` or `$GOBIN` directory. Make sure that directory is in your `PATH`.

## Usage

### Basic Usage

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
-d, --directory DIRECTORY    Directory name or path (use . for current directory)
-l, --language LANGUAGE      Programming language for .gitignore
-p, --private                Create a private repository (default: public)
-P, --public                 Create a public repository
-b, --branch BRANCH          Branch name (default: main)
-m, --message MESSAGE        Initial commit message (default: "Initial commit")
    --description DESC       Repository description
-n, --dry-run                Show what would happen without executing
-q, --quiet                  Minimal output
-h, --help                   Show help message
    version                  Show version
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
gitstart -d awesome-tool --description "An amazing CLI tool for developers"
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
gitstart -d . -l javascript --description "My existing JavaScript project"
```

**Show version:**
```sh
gitstart version
```

### Shell Completion

Gitstart supports Tab completion in your shell. Once set up, pressing Tab after typing part of a flag or subcommand will either complete it automatically or show you the available options. For example:

```sh
gitstart --di[TAB]        # completes to --directory
gitstart --[TAB]          # lists all flags
gitstart [TAB]            # lists all subcommands: completion, help, version
```

Run the setup command for your shell once, then open a new terminal (or source your config file):

**Bash**
```sh
gitstart completion bash >> ~/.bashrc
source ~/.bashrc
```

**Zsh**
```sh
gitstart completion zsh > "${fpath[1]}/_gitstart"
source ~/.zshrc
```

**Fish**
```sh
gitstart completion fish > ~/.config/fish/completions/gitstart.fish
```

**PowerShell**
```sh
gitstart completion powershell >> $PROFILE
```

### Working with Existing Directories

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

## Error Handling

- **Automatic cleanup**: If repository creation fails, the remote repository is automatically deleted
- **Validation checks**: Ensures all required tools are installed
- **Auth verification**: Confirms you're logged in to GitHub
- **File conflict detection**: Warns about existing files before overwriting
- **Detailed error messages**: Clear information about what went wrong and how to fix it

## About Licensing

Read more about [Licensing](https://docs.github.com/en/free-pro-team@latest/rest/reference/licenses).

## Changelog

### Version 1.0.0 (2026)

Gitstart is now rewritten in Go with full cross-platform support (macOS, Linux, Windows).

### Version 0.4.0 (2026-01-18)

**New Features:**
- Private repository support with `-p/--private` flag
- Custom commit messages with `-m/--message` flag
- Custom branch names with `-b/--branch` flag
- Repository description with `--description` flag
- Dry run mode with `--dry-run` flag
- Quiet mode with `-q/--quiet` flag
- Full support for existing directories and files
- Automatic rollback on errors
- Detection and handling of existing git repositories

**Improvements:**
- XDG-compliant config directory (`~/.config/gitstart/config`)
- Better error messages with context
- File conflict detection and user prompts
- Smarter handling of existing LICENSE, README, and .gitignore files

**Bug Fixes:**
- Fixed issue with `gh repo create --clone` in existing directories
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
