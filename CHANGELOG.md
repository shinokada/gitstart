# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.4.0] - 2026-01-18

### Added
- Private repository support with `-p/--private` flag
- Custom commit messages with `-m/--message` flag
- Custom branch names with `-b/--branch` flag
- Repository description with `-desc/--description` flag
- Dry run mode with `--dry-run` flag to preview changes without executing
- Quiet mode with `-q/--quiet` flag for minimal output
- Full support for existing directories with files
- Automatic cleanup/rollback if repository creation fails
- Detection and smart handling of existing git repositories
- Detection and handling of existing LICENSE, README.md, and .gitignore files
- User prompts for confirmation when working with existing files
- Comprehensive error handling with descriptive messages
- XDG-compliant configuration directory (`~/.config/gitstart/config`)

### Changed
- Improved GitHub authentication check (now properly uses exit codes)
- Configuration file location moved from `~/.gitstart_config` to `~/.config/gitstart/config`
- Better README.md template with more structured sections
- Enhanced .gitignore handling with append mode for existing files
- Improved user prompts and confirmations
- More informative success messages with repository details

### Fixed
- Fixed issue with `gh repo create --clone` failing in existing directories
- Fixed auth status check that incorrectly compared command output to integer
- Fixed potential data loss when running in directories with existing files
- Proper handling of directories that already contain a git repository
- Better error messages throughout the script

### Security
- Added error trap for automatic cleanup on failures
- Validation of directory paths to prevent running in HOME directory
- Better handling of edge cases to prevent unintended data loss

## [0.3.0] - 2021-12-18

### Added
- Initial public release
- Basic repository creation functionality
- GitHub CLI integration
- License selection (MIT, Apache 2.0, GNU GPLv3)
- .gitignore support for various programming languages
- README.md template generation
- Automatic git initialization and push

### Features
- Interactive license selection
- Programming language-specific .gitignore files
- GitHub username configuration storage
- Support for creating repositories in new directories
- Support for using current directory with `-d .`

[0.4.0]: https://github.com/shinokada/gitstart/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/shinokada/gitstart/releases/tag/v0.3.0
