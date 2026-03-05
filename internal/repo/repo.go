package repo

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

const (
	localCmdTimeout  = 30 * time.Second
	remoteCmdTimeout = 2 * time.Minute
)

// runCmd executes a command in dir with the given timeout, wrapping any error
// with the full command string and trimmed stderr for easier diagnosis.
func runCmd(dir string, timeout time.Duration, args ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, args[0], args[1:]...)
	cmd.Dir = dir
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(stderr.String()))
	}
	return nil
}

// InitGitRepo initializes a git repository in the given directory.
func InitGitRepo(dir string) error {
	return runCmd(dir, localCmdTimeout, "git", "init")
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
		if err := runCmd(dir, c.timeout, c.args...); err != nil {
			return err
		}
	}
	return nil
}

// DeleteGitHubRepo deletes a GitHub repository using the gh CLI.
// Used for cleanup when a subsequent step fails after repo creation.
func DeleteGitHubRepo(repoName string) error {
	return runCmd(".", remoteCmdTimeout, "gh", "repo", "delete", repoName, "--yes")
}

// CreateGitHubRepo creates a GitHub repository using the gh CLI and sets the remote origin,
// but does not push. Call CommitAndPush afterwards to stage, commit, and push.
func CreateGitHubRepo(dir, repoName, visibility, description string) error {
	args := []string{"gh", "repo", "create", repoName, "--source=.", "--remote=origin"}
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
	return runCmd(dir, remoteCmdTimeout, args...)
}
