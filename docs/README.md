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

- Auto-detect project language and create `.gitignore` (or use `-l` to specify)
- Create a license file based on your choice
- Create a new repository at GitHub.com (public or private)
- Create a `README.md` file with the repository name
- Initialize a git repository (if needed)
- Auto-detect the active branch from an existing repo
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

### After a Framework Starter

gitstart works seamlessly after scaffolding tools like `npx sv create`, `npm create vite@latest`, or `composer create-project`. Use `--post-framework` to skip prompts for files the framework already created, while still auto-detecting the language and branch:

```sh
npx sv create my-app && cd my-app && gitstart -d . --post-framework
npm create vite@latest my-app && cd my-app && gitstart -d . --post-framework
npm create vite@latest my-app -- --template react && cd my-app && gitstart -d . --post-framework
composer create-project laravel/laravel my-app && cd my-app && gitstart -d . --post-framework
npx nuxi@latest init my-app && cd my-app && gitstart -d . --post-framework
```

Or use quiet mode for a minimal one-liner:

```sh
npx sv create my-app && cd my-app && gitstart -d . -q
```

### Options

```text
-d, --directory DIRECTORY    Directory name or path (use . for current directory)
-l, --language LANGUAGE      Programming language for .gitignore (auto-detected if omitted)
-p, --private                Create a private repository (default: public)
-P, --public                 Create a public repository
-b, --branch BRANCH          Branch name (auto-detected from existing repo; default: main)
-m, --message MESSAGE        Initial commit message (default: "Initial commit")
    --description DESC       Repository description
    --no-license             Skip LICENSE file creation
    --no-readme              Skip README.md creation
    --post-framework         Optimised for use after a framework starter
                             (implies --no-license --no-readme)
-n, --dry-run                Show what would happen without executing
-q, --quiet                  Minimal output
-h, --help                   Show help message
    version                  Show version
```

### Language Auto-detection

When `-l` is not provided and no `.gitignore` exists, gitstart inspects the project directory for well-known marker files and infers the language automatically:

| Marker file(s) | Detected language |
|---|---|
| `go.mod` | Go |
| `Cargo.toml` | Rust |
| `pubspec.yaml` | Dart |
| `composer.json` | Composer (PHP) |
| `Gemfile` | Ruby |
| `pom.xml`, `build.gradle`, `build.gradle.kts` | Java |
| `requirements.txt`, `pyproject.toml`, `setup.py`, `setup.cfg` | Python |
| `package.json` | Node |

If multiple markers are present the first match in the table above wins. You can always override auto-detection with `-l`.

### Branch Auto-detection

When `--branch` is not explicitly set and a `.git` directory already exists (e.g. created by a framework starter), gitstart reads the active branch from `.git/HEAD` and pushes to that branch instead of defaulting to `main`. Passing `--branch` explicitly always takes precedence.

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

**Skip LICENSE and README (e.g. framework already created them):**
```sh
cd my-existing-project
gitstart -d . --no-license --no-readme
```

**Use --post-framework after a Svelte scaffold:**
```sh
npx sv create my-app
cd my-app
gitstart -d . --post-framework
```

**Use --post-framework after a React scaffold:**
```sh
npm create vite@latest my-app -- --template react
cd my-app
gitstart -d . --post-framework
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

**Empty directory:** Creates repository normally.

**Directory with files but no git:**
- Warns about existing files
- Asks for confirmation
- Preserves existing files
- Adds them to the initial commit

**Directory with existing git repository:**
- Detects existing `.git` folder
- Auto-detects the active branch from `.git/HEAD`
- Adds remote to existing repository
- Preserves git history

**Existing LICENSE, README.md, or .gitignore:**
- Detects existing files and skips them
- Use `--no-license` or `--no-readme` to explicitly suppress creation
- Use `--post-framework` to suppress both at once

### Interactive License Selection

When you run gitstart without `--no-license`, `--post-framework`, or `-q`, you'll be prompted to select a license:

```text
Select a license:
1) mit: Simple and permissive
2) apache-2.0: Community-friendly
3) gpl-3.0: Share improvements
4) None
```

## Error Handling

- **Automatic cleanup**: If repository creation fails, the remote repository is automatically deleted
- **Validation checks**: Ensures all required tools are installed
- **Auth verification**: Confirms you're logged in to GitHub
- **File conflict detection**: Detects existing files and skips safely
- **Detailed error messages**: Clear information about what went wrong and how to fix it

## About Licensing

Read more about [Licensing](https://docs.github.com/en/free-pro-team@latest/rest/reference/licenses).

## Changelog

### Version 1.2.2

**Bug Fixes:**
- Add v1.2.1 changelog and README entries (missed before tagging)

### Version 1.2.1

**Bug Fixes:**
- Fixed goreleaser CI failure: corrected `License` to `LICENSE` in
  `.goreleaser.yaml` (archives files list and nfpms contents)

### Version 1.2.0

**New Features:**
- Auto-detect project language from marker files (`go.mod`, `package.json`, `Cargo.toml`, etc.) when `-l` is not provided
- `--no-license` flag to skip LICENSE creation without suppressing all output
- `--no-readme` flag to skip README.md creation without suppressing all output
- `--post-framework` flag: optimised mode for use after framework starters — implies `--no-license --no-readme`
- Auto-detect active branch from `.git/HEAD` when `--branch` is not explicitly set

**Bug Fixes:**
- Fixed dry-run language detection to always auto-detect (not only when `--post-framework` is set)
- Fixed `composer.json` marker mapping from `PHP` to `Composer` (PHP.gitignore does not exist in GitHub/gitignore)
- Renamed internal `resolvDir` to `resolveDir` (typo fix)

### Version 1.1.0

- Added shell completion support (bash, zsh, fish, PowerShell)

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
