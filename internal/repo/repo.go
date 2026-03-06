package repo

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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

// InitGitRepo initializes a git repository in the given directory with the
// given initial branch name. Passing an explicit branch makes the result
// deterministic regardless of the system-level init.defaultBranch setting.
func InitGitRepo(dir, branch string) error {
	return runCmd(dir, localCmdTimeout, "git", "init", "-b", branch)
}

// resolveGitDir returns the path to the real .git directory for dir.
// When .git is a file (worktrees, submodules), it contains a "gitdir: <path>"
// pointer that must be followed to find the actual git directory.
func resolveGitDir(dir string) string {
	gitPath := filepath.Join(dir, ".git")
	info, err := os.Stat(gitPath)
	if err != nil {
		return ""
	}
	if info.IsDir() {
		return gitPath
	}
	// .git is a file — read the "gitdir: <path>" pointer.
	data, err := os.ReadFile(gitPath)
	if err != nil {
		return ""
	}
	line := strings.TrimSpace(string(data))
	const filePrefix = "gitdir: "
	if !strings.HasPrefix(line, filePrefix) {
		return ""
	}
	pointed := strings.TrimPrefix(line, filePrefix)
	if !filepath.IsAbs(pointed) {
		pointed = filepath.Join(dir, pointed)
	}
	return filepath.Clean(pointed)
}

// DetectCurrentBranch reads the active branch from an existing git repo in dir.
// It resolves .git file indirection (worktrees, submodules) before reading HEAD.
// Returns "" if no repo is found, HEAD is unreadable, or HEAD is detached.
func DetectCurrentBranch(dir string) string {
	gitDir := resolveGitDir(dir)
	if gitDir == "" {
		return ""
	}
	data, err := os.ReadFile(filepath.Join(gitDir, "HEAD"))
	if err != nil {
		return ""
	}
	// HEAD contains "ref: refs/heads/<branch>\n" when on a named branch.
	line := strings.TrimSpace(string(data))
	const prefix = "ref: refs/heads/"
	if !strings.HasPrefix(line, prefix) {
		return ""
	}
	return strings.TrimPrefix(line, prefix)
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
