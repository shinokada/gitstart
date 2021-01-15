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

## Requirement

- Install [GitHub CLI](https://cli.github.com/manual/).
- Install [yq@3](https://github.com/mikefarah/yq)

```bash
# intall yq
$ brew install yq
```

- Login github using `gh auth login`.
- Choose SSH as the default git protocol when you login.

## Usage

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

Create a project directory and then run `gitstart`. Follow the instruction. It will ask "Visibility" and "This will create 'your_repo' in your current directory. Continue?".

You can add your programming language to add `.gitignore` file.

```bash
# For Python
$ gitstart python
# For Go
$ gitstart go
```

See the example below:

```bash
$ mkdir my_new_repo
$ cd my_new_repo
$ mkdir my_new_repo
$ cd my_new_repo
❯ gitstart python
>>> Your github username is shinokada.
Select a license:
1) MIT: I want it simple and permissive.
2) Apache License 2.0: I need to work in a community.
3) GNU GPLv3: I care about sharing improvements.
4) Quit
Your lisence: 2
Apache
>>> Creating .gitignore for Python...
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:100  2035  100  2035    0     0  26776      0 --:--:-- --:--:-- --:--:-- 27133
>>> .gitignore created.
>>> Creating READMR.md.
>>> Running git init.
Initialized empty Git repository in /Users/shinokada/Downloads/my_new_repo/.git/
>>> Adding README.md and .gitignore.
>>> Commiting with a message 'first commit'.
[master (root-commit) beafcfa] first commit
 12 files changed, 363 insertions(+)
 create mode 100644 .gitignore
 create mode 100644 README.md
 create mode 100644 license.txt
>>> Creating the main branch.
github.com
  ✓ Logged in to github.com as shinokada (~/.config/gh/hosts.yml)
  ✓ Git operations for github.com configured to use ssh protocol.
  ✓ Token: *******************

>>> You are logged in. Creating your my_new_repo in remote.
? Visibility Public
? This will create 'my_new_repo' in your current directory. Continue?  Yes
✓ Created repository shinokada/my_new_repo on GitHub
✓ Added remote git@github.com:shinokada/my_new_repo.git
fatal: remote origin already exists.
>>> Pushing local repo to the remote.
Enumerating objects: 7, done.
Counting objects: 100% (7/7), done.
Delta compression using up to 4 threads
Compressing objects: 100% (6/6), done.
Writing objects: 100% (7/7), 5.44 KiB | 2.72 MiB/s, done.
Total 7 (delta 0), reused 0 (delta 0)
To github.com:shinokada/my_new_repo.git
 * [new branch]      main -> main
Branch 'main' set up to track remote branch 'main' from 'origin'.
>>> You have created a github repo at https://github.com/shinokada/my_new_repo
```

If you prefer not to add `.gitignore`.

```bash
# start it
$ gitstart
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
