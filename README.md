<p align="center">
<img width="400" src="https://raw.githubusercontent.com/shinokada/gitstart/main/images/gitstart.png" />
</p>
<h1  align="center">Gitstart</h1>

[More details about Gitstart.](https://towardsdatascience.com/automate-creating-a-new-github-repository-with-gitstart-1ae961b99866)

## Overview

> Gitstart creates, adds, and pushes with one line.

This script automates creating a git repo. The script will:

- Create .gitignore if you provide a language.
- Create a license.txt depends on your choice.
- Create a new repo at GitHub.com.
- Create a README.md file with the repo name.
- Add README.md and commit with a message.
- Add the remote and push the file.

The script reads your GitHub username from ~/.config/gh/hosts.yml and uses the directory name as a GitHub repo name.

## Requirements

- UNIX-lie (Tested on Ubuntu and MacOS.)
- [GitHub CLI](https://cli.github.com/manual/).
- [jq](https://stedolan.github.io/jq/).

## Installation

### Linux/macOS using Awesome

After installing [Awesome package manager](https://github.com/shinokada/awesome):

```sh
awesome install shinokada/gitstart
```

### macOS using Homebrew

If you have Homebrew on your macOS, your can run:

```sh
brew tap shinokada/gitstart && brew install gitstart
```

## Usage

- Login github using `gh auth login`.
- Choose SSH or HTTPS for the default git protocol when you login.

```sh
# define a dir path
gitstart repo-name
# in a current dir
cd new_repo
gitstart
```

- Select a license.
  
```sh
Select a license:
1) MIT: I want it simple and permissive.
2) Apache License 2.0: I need to work in a community.
3) GNU GPLv3: I care about sharing improvements.
4) None
5) Quit
Your lisence: 1
MIT
```

- Select a visibility.

```sh
>>> You are logged in. Creating your newtest in remote.
? Visibility  [Use arrows to move, type to filter]
> Public
  Private
  Internal
```

- Yes to add an origin git remote to your local repo.

```sh
? This will add an "origin" git remote to your local repository. Continue? Yes
```

## About Licensing

Read more about [Licensing](https://docs.github.com/en/free-pro-team@latest/rest/reference/licenses).

## Author

Shinichi Okada

- [Medium](https://shinichiokada.medium.com/)
- [Twitter](https://twitter.com/shinokada)

## License

Copyright (c) 2021 Shinichi Okada (@shinokada)
This software is released under the MIT License, see LICENSE.
