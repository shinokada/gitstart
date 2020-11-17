# gitstart

This script automates creating a git repo. The script will:

- Read your GitHub username from ~/.config/gh/hosts.yml.
- Uses the directory name as a GitHub repo name.
- Create a README.md file with the repo name.
- Add README.md and commit with a message.
- Create a new repo at GitHub.com
- Add the remote and push the file.

## Requirement

- You must install [GitHub CLI](https://cli.github.com/manual/).
- Install [yq](https://github.com/mikefarah/yq)
- When you login GitHub, choose SSH as the default git protocol.
- Login github using `gh auth login`.

## How to use it

Download the gitstart file or cron this repo.
Make the file executable.

```bash
$ chmod 755 gitstart
```

Create a directory and then run `gitstart`.

```bash
$ mkdir my_new_repo
$ cd my_new_repo
$ gitstart
Initialized empty Git repository in /Users/shinokada/gitstart/.git/
[master (root-commit) 41bc594] first commit
 1 file changed, 1 insertion(+)
 create mode 100644 README.md
github.com
  ✓ Logged in to github.com as shinokada (~/.config/gh/hosts.yml)
  ✓ Git operations for github.com configured to use ssh protocol.
  ✓ Token: *******************

? Visibility Public
? This will create 'gitstart' in your current directory. Continue?  Yes
✓ Created repository shinokada/gitstart on GitHub
✓ Added remote git@github.com:shinokada/gitstart.git
fatal: remote origin already exists.
Enumerating objects: 3, done.
Counting objects: 100% (3/3), done.
Writing objects: 100% (3/3), 226 bytes | 226.00 KiB/s, done.
Total 3 (delta 0), reused 0 (delta 0)
To github.com:shinokada/gitstart.git
 * [new branch]      main -> main
Branch 'main' set up to track remote branch 'main' from 'origin'.
You have created a github repo at https://github.com/shinokada/gitstart
```
