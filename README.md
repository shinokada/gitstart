# Gitstart

[Please read more details about Gitstart.](https://towardsdatascience.com/automate-creating-a-new-github-repository-with-gitstart-1ae961b99866)

[Find the update at Github.](https://github.com/shinokada/gitstart)

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

- [jq](https://stedolan.github.io/jq/)
- [GitHub CLI](https://cli.github.com/manual/).
- [yq@3](https://github.com/mikefarah/yq)

## Installation

- Using homebrew:

```sh
brew tap shinokada/gitstart && brew install gitstart
```

- Manually

Download the gitstart file or cron this repo.
Make the file executable.

```bash
# change permission
$ chmod 755 gitstart
```

Create a `~/bin` directory and add the path to your terminal config file, `~/.zshrc` or `~/.bashrc`.

```.zshrc
export PATH="~/bin:$PATH"
```

## Usage

- Login github using `gh auth login`.
- Choose SSH as the default git protocol when you login.

```sh
# define a dir path
gitstart -d ./path
# in a current dir
cd new_repo
gitstart -d .
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

### Adding programming language .gitignore

By adding a programming language, it will add `.gitignore` file.

```bash
# For Python
$ gitstart -d ./repo_name -l python
# For Go
$ gitstart -d ./repo_name -l go
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
