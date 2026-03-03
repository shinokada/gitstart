package repo

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"
)

const (
	localCmdTimeout  = 30 * time.Second
	remoteCmdTimeout = 2 * time.Minute
)

// InitGitRepo initializes a git repository in the given directory.
func InitGitRepo(dir string) error {
	ctx, cancel := context.WithTimeout(context.Background(), localCmdTimeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, "git", "init")
	cmd.Dir = dir
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git init: %w: %s", err, stderr.String())
	}
	return nil
}

// CommitAndPush stages all files, commits, and pushes to the remote repository.
func CommitAndPush(dir, branch, message string) error {
	type commandSpec struct {
		args    []string
		timeout time.Duration
	}
	cmds := []commandSpec{
		{args: []string{"git", "add", "."}, timeout: localCmdTimeout},
		// --allow-empty ensures the commit always succeeds even if there are no
		// staged changes, which is the expected behaviour for an automation CLI.
		{args: []string{"git", "commit", "--allow-empty", "-m", message}, timeout: localCmdTimeout},
		{args: []string{"git", "branch", "-M", branch}, timeout: localCmdTimeout},
		{args: []string{"git", "push", "-u", "origin", branch}, timeout: remoteCmdTimeout},
	}
	for _, c := range cmds {
		ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
		cmd := exec.CommandContext(ctx, c.args[0], c.args[1:]...)
		cmd.Dir = dir
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		err := cmd.Run()
		cancel()
		if err != nil {
			return fmt.Errorf("%s: %w: %s", c.args[0], err, stderr.String())
		}
	}
	return nil
}

// CreateGitHubRepo creates a GitHub repository using the gh CLI.
func CreateGitHubRepo(dir, repoName, visibility, description string) error {
	args := []string{"repo", "create", repoName, "--source=.", "--remote=origin", "--push"}
	switch visibility {
	case "":
	case "public", "private", "internal":
		args = append(args, "--"+visibility)
	default:
		return fmt.Errorf("invalid visibility %q: expected public, private, or internal", visibility)
	}
	if description != "" {
		args = append(args, "--description", description)
	}
	ctx, cancel := context.WithTimeout(context.Background(), remoteCmdTimeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, "gh", args...)
	cmd.Dir = dir
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("gh repo create: %w: %s", err, stderr.String())
	}
	return nil
}
