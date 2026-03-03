package repo

import (
	"context"
	"os/exec"
	"time"
)

const cmdTimeout = 30 * time.Second

// InitGitRepo initializes a git repository in the given directory.
func InitGitRepo(dir string) error {
	ctx, cancel := context.WithTimeout(context.Background(), cmdTimeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, "git", "init")
	cmd.Dir = dir
	return cmd.Run()
}

// CommitAndPush stages all files, commits, and pushes to the remote repository.
func CommitAndPush(dir, branch, message string) error {
	cmds := [][]string{
		{"git", "add", "."},
		// --allow-empty ensures the commit always succeeds even if there are no
		// staged changes, which is the expected behaviour for an automation CLI.
		{"git", "commit", "--allow-empty", "-m", message},
		{"git", "branch", "-M", branch},
		{"git", "push", "-u", "origin", branch},
	}
	for _, args := range cmds {
		ctx, cancel := context.WithTimeout(context.Background(), cmdTimeout)
		cmd := exec.CommandContext(ctx, args[0], args[1:]...)
		cmd.Dir = dir
		err := cmd.Run()
		cancel()
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateGitHubRepo creates a GitHub repository using the gh CLI.
func CreateGitHubRepo(dir, repoName, visibility, description string) error {
	args := []string{"repo", "create", repoName, "--source=.", "--remote=origin", "--push"}
	if visibility != "" {
		args = append(args, "--"+visibility)
	}
	if description != "" {
		args = append(args, "--description", description)
	}
	ctx, cancel := context.WithTimeout(context.Background(), cmdTimeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, "gh", args...)
	cmd.Dir = dir
	return cmd.Run()
}
