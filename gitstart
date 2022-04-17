#!/usr/bin/env bash

########################
# Author: Shinichi Okada
# Date: 2021-12-18
###########################

unset github_username dir repo license_url

script_name=$(basename "$0")
dir=""
version="0.3.0"
gitstart_cofig=$HOME/.gitstart_config

usage() {
    cat <<EOF

Script name: $script_name

Description:
============

This script automate creating git repo.
This will create a dir, add a license, README file, run git init and push to Github. 

If a language is added, it will try to add the .gitignore.
Choose SSH as the default git protocol.

List of .gitignore supported by Github
======================================

See https://github.com/github/gitignore

Usage:
======

$script_name [ -l | --lang programming_language ] [ -d | --dir directory ] [ -h | --help | -v | --version]

Examples:
=========

    $script_name -d newrepo
    $script_name -d newrepo -l python 
    # in a directory
    $script_name -d .
    # Show help
    $script_name -h
    # Show version
    $script_name -v
EOF
}

while (($# > 0)); do # while arguments count>0
    case "$1" in
    -l | --lang)
        prog_lang=$2
        shift
        ;;
    -d | --dir)
        dir=$2
        shift 2
        ;;
    -h | --help)
        usage
        exit
        ;;
    -v | --version)
        echo ${version}
        exit
        ;;
    *) # unknown flag/switch
        POSITIONAL+=("$1")
        shift
        ;;
    esac
done

##### check
if [ ! "$(command -v gh)" ]; then
    echo "Please install gh from https://github.com/cli/cli#installation."
    exit 1
fi

if [ ! "$(command -v jq)" ]; then
    echo "Please install jq from https://stedolan.github.io/jq/."
    exit 1
fi

############# Main body ############
# check if you are logged in github
if [[ $(gh auth status) -eq 1 ]]; then
    # not logged-in
    echo ">>> You must logged in. Use 'gh auth login'"
    exit 1
fi

# Directory path. If dir is . then use pwd
if [[ ${dir} = "." ]]; then
    dir=$(pwd)
else
    echo ">>> Creating ${dir}."
    dir="$(pwd)/$dir"
    # mkdir -p "$dir" || exit
    # cd "$dir" || exit
fi

# don't allow to create a git repo in the ~ (HOME)
if [[ (${dir} = "$HOME") ]]; then
    echo "This script doesn't allow to create a git repo in the home directory."
    echo "Use another directory."
    exit 1
fi

gitname() {
    printf "Please type your Github username. "
    read -r github_username
    echo "$github_username" >"$gitstart_cofig"
}

if [ ! -s "$gitstart_cofig" ]; then
    gitname
else
    github_username=$(cat "$gitstart_cofig")
    echo "Is it correct your GitHub username is $github_username. y/yes/n/no"
    read -r ANS
    ans=$(echo "$ANS" | cut -c 1-1 | tr "[:lower:]" "[:upper:]")
    if [[ $ans = "Y" ]]; then
        :
    else
        gitname
    fi
fi

repo=$(basename "${dir}")
license_url="mit"
echo ">>> Your github username is ${github_username}."
echo ">>> Your new repo name is ${repo}."

# lisence
PS3='Your license: '
licenses=("MIT: I want it simple and permissive." "Apache License 2.0: I need to work in a community." "GNU GPLv3: I care about sharing improvements." "None" "Quit")

echo "Select a license: "
select license in "${licenses[@]}"; do
    case ${license} in
    "MIT: I want it simple and permissive.")
        echo "MIT"
        break
        ;;
    "Apache License 2.0: I need to work in a community.")
        echo "Apache"
        license_url="apache-2.0"
        break
        ;;
    "GNU GPLv3: I care about sharing improvements.")
        echo "GNU"
        license_url="lgpl-3.0"
        break
        ;;
    "None")
        echo "License None"
        license_url=false
        break
        ;;
    "Quit")
        echo "User requested exit."
        exit
        ;;
    *) echo "Invalid option $REPLY" ;;
    esac
done

if [ ! -d "$dir" ]; then
    mkdir -p "$dir"
fi

# creating a remote github repo using gh
printf "Creating a public remote repo %s" "$dir"
gh repo create "$repo" --public --clone || {
    echo "Not able to create a remote repo."
    exit 1
}

cd "$dir" || exit

if [[ ${license_url} != false ]]; then
    touch "${dir}"/LICENSE
    curl -s "https://api.github.com/licenses/${license_url}" | jq -r '.body' >"${dir}"/LICENSE
    echo ">>> LICENSE is created."
fi

if [[ ${prog_lang} ]]; then
    # github gitignore url
    url=https://raw.githubusercontent.com/github/gitignore/master/"${prog_lang^}".gitignore
    # get the status http code, 200, 404 etc.
    http_status=$(curl --write-out '%{http_code}' --silent --output /dev/null "${url}")

    if [[ $http_status -eq 200 ]]; then
        # get argument e.g python, go etc, capitalize it
        echo ">>> Creating .gitignore for ${1^}..."
        # create .gitignore
        curl "${url}" -o .gitignore
        echo ">>> .gitignore created."
    else
        echo ">>> Not able to find ${1^} gitignore at https://github.com/github/gitignore."
        echo ">>> Adding .gitignore with minimum contents."
        touch "${dir}/.gitignore"
        echo ".DS_Store" >"${dir}/.gitignore"
    fi
else
    echo ">>> Adding .gitignore with minimum contents."
    touch "${dir}/.gitignore"
    echo ".DS_Store" >"${dir}/.gitignore"
fi

# README
echo ">>> Creating README.md."
printf "# %s \n
## Overview\n\n
## Requirement\n\n
## Usage\n\n
## Features\n\n
## Reference\n\n
## Author\n\n
## License

Please see license.txt.\n" "${repo}" >README.md

# git commands

echo ">>> Adding README.md and .gitignore."
git add . || {
    echo "Not able to add."
    exit 1
}
echo ">>> Commiting with a message 'first commit'."
git commit -m "first commit" || {
    echo "Not able to commit."
    exit 1
}
git branch -M main || {
    echo "Not able to add the branch"
    exit 1
}
git push -u origin main || {
    echo "Not able to push to a remote repo."
    exit 1
}

echo ">>> You have created a github repo at https://github.com/${github_username}/${repo}"

exit 0
