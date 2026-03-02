package repo

import (
	"os/exec"
)

// InitGitRepo initializes a git repository in the given directory.
func InitGitRepo(dir string) error {
	cmd := exec.Command("git", "init")
	cmd.Dir = dir
	return cmd.Run()
}

// CommitAndPush stages all files, commits, and pushes to the remote repository.
func CommitAndPush(dir, branch, message string) error {
	cmds := [][]string{
		{"git", "add", "."},
		{"git", "commit", "-m", message},
		{"git", "branch", "-M", branch},
		{"git", "push", "-u", "origin", branch},
	}
	for _, args := range cmds {
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Dir = dir
		if err := cmd.Run(); err != nil {
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
	cmd := exec.Command("gh", args...)
	cmd.Dir = dir
	return cmd.Run()
}
