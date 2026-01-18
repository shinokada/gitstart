# Gitstart Usage Examples

This document provides real-world examples and use cases for gitstart.

## Table of Contents

- [Quick Start Examples](#quick-start-examples)
- [Programming Language Examples](#programming-language-examples)
- [Repository Configuration Examples](#repository-configuration-examples)
- [Existing Project Examples](#existing-project-examples)
- [Team Workflow Examples](#team-workflow-examples)
- [Advanced Usage Examples](#advanced-usage-examples)

## Quick Start Examples

### Create a Simple Project

```bash
# Create a basic repository
gitstart -d my-first-repo

# The script will:
# 1. Ask for license (choose MIT)
# 2. Create directory
# 3. Initialize git
# 4. Create LICENSE, README.md, .gitignore
# 5. Push to GitHub
```

### Preview Before Creating

```bash
# See what will happen without making changes
gitstart -d my-project --dry-run

# Review the output, then create for real
gitstart -d my-project
```

### Create a Private Repository

```bash
# For projects you don't want public
gitstart -d secret-project -p
```

## Programming Language Examples

### Python Project

```bash
gitstart -d my-python-app -l python

# Creates with Python-specific .gitignore:
# - __pycache__/
# - *.py[cod]
# - venv/
# - .pytest_cache/
# etc.
```

### JavaScript/Node.js Project

```bash
gitstart -d my-node-app -l node

# Creates with Node.js .gitignore:
# - node_modules/
# - npm-debug.log
# - .env
# etc.
```

### Go Project

```bash
gitstart -d my-go-app -l go -b main

# Creates with Go-specific .gitignore
```

### Rust Project

```bash
gitstart -d my-rust-app -l rust

# Creates with Rust-specific .gitignore:
# - target/
# - Cargo.lock (for binaries)
# etc.
```

### Java Project

```bash
gitstart -d my-java-app -l java

# Creates with Java-specific .gitignore:
# - *.class
# - target/
# - .gradle/
# etc.
```

### Multiple File Types

```bash
# For projects with multiple languages
gitstart -d fullstack-app -l python

# Then manually add Node.js ignores
cd fullstack-app
curl -s https://raw.githubusercontent.com/github/gitignore/master/Node.gitignore >> .gitignore
git add .gitignore
git commit -m "Add Node.js to .gitignore"
git push
```

## Repository Configuration Examples

### With Custom Commit Message

```bash
# Professional initial commit
gitstart -d production-app -m "Initial production release v1.0.0"

# Feature branch initialization
gitstart -d feature-auth -m "Initialize authentication feature"

# Documentation project
gitstart -d api-docs -m "Initialize API documentation"
```

### With Custom Branch Name

```bash
# Using 'develop' branch
gitstart -d my-app -b develop

# Using 'master' (legacy projects)
gitstart -d legacy-project -b master

# Feature branch
gitstart -d new-feature -b feature/user-dashboard
```

### With Description

```bash
# Clear description for discoverability
gitstart -d awesome-cli \
    -desc "A powerful command-line tool for developers"

# Descriptive project
gitstart -d data-analysis \
    -desc "Python scripts for data analysis and visualization"

# Library project
gitstart -d my-lib \
    -desc "A reusable library for web development"
```

### Complete Configuration

```bash
# All options together
gitstart -d complete-project \
    -l python \
    -p \
    -b develop \
    -m "Project initialization with complete setup" \
    -desc "A comprehensive example of gitstart usage"
```

## Existing Project Examples

### Initialize Current Directory

```bash
# You're in a project directory
cd ~/projects/existing-app

# Initialize git repository here
gitstart -d .

# The script will:
# - Detect existing files
# - Ask for confirmation
# - Preserve all files
# - Add them to initial commit
```

### Project with Existing Files

```bash
# You have an existing project
cd ~/Downloads/client-project
ls
# main.py  requirements.txt  README.txt

# Initialize git
gitstart -d . -l python

# Answer 'y' when prompted about existing files
# All files will be included in initial commit
```

### Project with Existing Git Repository

```bash
# You already have local git history
cd ~/code/local-project
git log  # Shows existing commits

# Add GitHub remote and push
gitstart -d .

# The script will:
# - Detect existing .git
# - Create remote repository
# - Add remote origin
# - Push existing commits
```

### Existing Project with Custom Files

```bash
# Project with custom LICENSE and README
mkdir my-project
cd my-project
echo "Custom License Terms" > LICENSE
echo "# My Custom README" > README.md
echo "print('hello')" > main.py

# Initialize git without overwriting
gitstart -d . -l python

# When prompted:
# - LICENSE: Script will skip (file exists)
# - README: Script will skip (file exists)  
# - .gitignore: Will append Python rules
```

## Team Workflow Examples

### Setting Up Team Repository

```bash
# Team lead initializes repo
gitstart -d team-project \
    -l javascript \
    -b develop \
    -m "Initialize team project" \
    -desc "Collaborative project for Team Alpha"

# Share with team
echo "Team: Clone the repo with:"
echo "git clone git@github.com:username/team-project.git"
```

### Feature Branch Workflow

```bash
# Developer creates feature branch
git clone git@github.com:team/main-project.git
cd main-project
git checkout -b feature/new-api

# Or initialize separate feature repo
gitstart -d new-api-feature \
    -b feature/new-api \
    -m "Start new API feature"
```

### Microservices Setup

```bash
# Service 1: User Service
gitstart -d user-service \
    -l go \
    -p \
    -desc "User authentication and management service"

# Service 2: Payment Service
gitstart -d payment-service \
    -l python \
    -p \
    -desc "Payment processing service"

# Service 3: Notification Service
gitstart -d notification-service \
    -l node \
    -p \
    -desc "Email and SMS notification service"
```

## Advanced Usage Examples

### Automated Repository Creation

```bash
#!/bin/bash
# Script to create multiple related repositories

projects=(
    "frontend:javascript:React application"
    "backend:python:Django REST API"
    "mobile:swift:iOS application"
    "docs:markdown:Project documentation"
)

for project in "${projects[@]}"; do
    IFS=':' read -r name lang desc <<< "$project"
    gitstart -d "$name" \
        -l "$lang" \
        -desc "$desc" \
        -m "Initialize $name" \
        -q
    sleep 2  # Rate limiting
done

echo "Created ${#projects[@]} repositories"
```

### Monorepo Initialization

```bash
# Create main monorepo
gitstart -d my-monorepo \
    -m "Initialize monorepo structure" \
    -desc "Monorepo for all company projects"

cd my-monorepo

# Create package structure
mkdir -p packages/{web,mobile,shared}

# Add .gitignore for each
echo "node_modules/" > packages/web/.gitignore
echo "node_modules/" > packages/mobile/.gitignore
echo "*.js" > packages/shared/.gitignore

git add .
git commit -m "Add package structure"
git push
```

### Template Repository

```bash
# Create a template
gitstart -d project-template \
    -l python \
    -desc "Template for new Python projects"

cd project-template

# Add template structure
mkdir -p src tests docs
echo "from setuptools import setup, find_packages" > setup.py
echo "pytest" > requirements-dev.txt

# Add to repository
git add .
git commit -m "Add project template structure"
git push

# On GitHub: Settings → Template repository (check box)
```

### Documentation Site

```bash
# Initialize docs repository
gitstart -d docs-site \
    -b gh-pages \
    -m "Initialize documentation site" \
    -desc "Project documentation and guides"

cd docs-site

# Create docs structure
mkdir -p docs/{guides,api,tutorials}
echo "# Documentation" > docs/index.md

git add .
git commit -m "Add docs structure"
git push
```

### Experiment/POC Repository

```bash
# Quick experiment
gitstart -d experiment-neural-net \
    -l python \
    -p \
    -m "POC: Neural network approach" \
    -desc "Experimental implementation of neural network"

# If successful, make it public later:
# gh repo edit experiment-neural-net --visibility public
```

### Client Project Setup

```bash
# For client work
CLIENT="acme-corp"
PROJECT="website-redesign"

gitstart -d "${CLIENT}-${PROJECT}" \
    -l javascript \
    -p \
    -m "Initialize ${CLIENT} ${PROJECT}" \
    -desc "Website redesign project for ${CLIENT}"

cd "${CLIENT}-${PROJECT}"

# Add client-specific structure
mkdir -p {assets,components,pages}
echo "# ${CLIENT} - ${PROJECT}" > README.md
git add .
git commit -m "Add project structure"
git push
```

### Migration from Other VCS

```bash
# You have an existing SVN/Mercurial project
cd existing-svn-project

# Export/convert to git first (varies by VCS)
# Then initialize GitHub repo

gitstart -d . \
    -m "Migrate from SVN to Git" \
    -desc "Historical project migrated from SVN"
```

### Multi-Language Project

```bash
# Full-stack application
gitstart -d fullstack-app -l javascript

cd fullstack-app

# Add Python backend ignores
echo "
# Python
__pycache__/
*.py[cod]
venv/
" >> .gitignore

# Add Go API ignores  
echo "
# Go
vendor/
*.exe
" >> .gitignore

git add .gitignore
git commit -m "Add multi-language .gitignore"
git push
```

### Quiet Mode for Scripts

```bash
#!/bin/bash
# Automated project creation

create_project() {
    local name=$1
    local lang=$2
    
    gitstart -d "$name" -l "$lang" -q || {
        echo "Failed to create $name" >&2
        return 1
    }
    
    echo "Created: $name"
}

# Create multiple projects silently
create_project "api-gateway" "go"
create_project "user-service" "python"
create_project "frontend" "javascript"
```

## Tips and Tricks

### Testing Before Creating

```bash
# Always test with dry-run first
gitstart -d new-idea --dry-run

# Review output, adjust parameters
gitstart -d new-idea -l python -p
```

### Consistent Project Structure

```bash
# Create a shell function for your team's standard setup
mk-project() {
    local name=$1
    local type=${2:-python}
    
    gitstart -d "$name" \
        -l "$type" \
        -b develop \
        -m "Initialize $name" \
        -desc "Project: $name"
    
    cd "$name"
    mkdir -p {src,tests,docs}
    git add .
    git commit -m "Add project structure"
    git push
    cd ..
}

# Use it
mk-project "my-awesome-app" "javascript"
```

### Combining with Other Tools

```bash
# Create repo and set up virtual environment
gitstart -d python-project -l python
cd python-project
python -m venv venv
source venv/bin/activate
pip install -r requirements.txt

# Create repo and initialize npm
gitstart -d node-project -l node
cd node-project
npm init -y
npm install

# Create repo and initialize cargo
gitstart -d rust-project -l rust
cd rust-project
cargo init
```

## Troubleshooting Examples

### Repository Name Conflicts

```bash
# Check if name exists first
gh repo view username/my-repo &>/dev/null && echo "Exists" || echo "Available"

# If exists, use different name
gitstart -d my-repo-v2
```

### Network Issues

```bash
# Test connection first
gh auth status
gh repo list --limit 1

# Then create
gitstart -d my-project
```

### Permission Issues

```bash
# Ensure you have write access
gh repo view username/org-repo

# Create in personal account
gitstart -d personal-project

# For organization repos, use gh directly
gh repo create org-name/repo-name
```

These examples cover most common use cases. Adapt them to your specific needs!
