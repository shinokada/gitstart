# Gitstart

> Gitstart creates, adds, and pushes with one line.

This script automates creating a git repo. The script will:

- Create .gitignore if you provide a language.
- Create a new repo at GitHub.com.
- Create a README.md file with the repo name.
- Add README.md and commit with a message.
- Add the remote and push the file.

The script reads your GitHub username from ~/.config/gh/hosts.yml and uses the directory name as a GitHub repo name.

## Requirement

- Install [GitHub CLI](https://cli.github.com/manual/).
- Install [yq](https://github.com/mikefarah/yq)
- Login github using `gh auth login`.
- Choose SSH as the default git protocol when you login.

## How to use it

Download the gitstart file or cron this repo.
Make the file executable.

```bash
$ chmod 755 gitstart
```

Create a directory and then run `gitstart`. Follow the instruction. It will ask "Visibility" and "This will create 'your_repo' in your current directory. Continue?".

See the example below:

```bash
$ mkdir my_new_repo
$ cd my_new_repo
$ gitstart python
>>> Your github username is shinokada.
>>> Creating .gitignore for Python...
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  2035  100  2035    0     0  29926      0 --:--:-- --:--:-- --:--:-- 29926
>>> .gitignore created.
>>> Creating READMR.md.
>>> Running git init.
Initialized empty Git repository in /Users/shinokada/testdir/my_new_repo/.git/
>>> Adding README.md and .gitignore.
>>> Commiting with a message 'first commit'.
[master (root-commit) 685f369] first commit
 2 files changed, 139 insertions(+)
 create mode 100644 .gitignore
 create mode 100644 README.md
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
Enumerating objects: 4, done.
Counting objects: 100% (4/4), done.
Delta compression using up to 4 threads
Compressing objects: 100% (3/3), done.
Writing objects: 100% (4/4), 1.33 KiB | 1.33 MiB/s, done.
Total 4 (delta 0), reused 0 (delta 0)
To github.com:shinokada/my_new_repo.git
 * [new branch]      main -> main
Branch 'main' set up to track remote branch 'main' from 'origin'.
>>> You have created a github repo at https://github.com/shinokada/my_new_repo
```

If you prefer not to add `.gitignore`.

```terminal
$ gitstart
```
