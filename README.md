# Gitstart for macOS/Linux

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


- [gh](https://cli.github.com/)
- [jq](https://stedolan.github.io/jq/)
- [GitHub CLI](https://cli.github.com/manual/).
- [yq](https://github.com/mikefarah/yq)

**Login to Github using `gh`.**

## Installation

### macOS using Homebrew

If you have Homebrew on your macOS, your can run:

```sh
brew tap shinokada/gitstart && brew install gitstart
```

### Linux

On Ubuntu install `yq`:

```sh
sudo snap install yq
```

I keep `gitstart` in the `/home/your-username/awesome` directory:

```sh
mkdir /home/your-username/awesome
cd /home/your-username/awesome
git clone https://github.com/shinokada/manop.git
```

Create the `~/bin` directory:

```sh
mkdir ~/bin
```

Check if `/home/your-username/bin` in the PATH variable:

```sh
echo $PATH
/home/your-username/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin
```

If not add the `/home/your-username/bin` directory to the ~/.bashrc file.

```sh
export PATH="/home/your-username/bin:$PATH"
```

Source the `~/.bashrc` file and check it again:

```sh
source ~/.bashrc
echo $PATH
```

Add a symlink:

```sh
ln -s /home/your-username/awesome/giststart/gitstart ~/bin/gitstart
```

Check if the symlink is working:

```
which gitstart
/home/your-username/bin/gitstart
gitstart -h

Script name: gitstart

Description:
...
```

## Usage

- Login github using `gh auth login`.
- Choose SSH or HTTPS for the default git protocol when you login.

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
