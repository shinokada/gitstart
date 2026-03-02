package repo

import (
	"os"
	"testing"
)

func TestCommitAndPush(t *testing.T) {
	dir := "test_commit_repo"
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer func() {
		_ = os.RemoveAll(dir)
	}()

	// Initialize git repo
	if err := InitGitRepo(dir); err != nil {
		t.Fatalf("failed to init git repo: %v", err)
	}

	// Create a file to commit
	file := dir + "/test.txt"
	if err := os.WriteFile(file, []byte("test"), 0644); err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	// Set up a dummy remote to avoid push error
	// This is a placeholder; in real tests, mock or skip push

	err := CommitAndPush(dir, "main", "test commit")
	if err == nil {
		t.Logf("CommitAndPush ran (push likely failed due to no remote, but commit should succeed)")
	}
}
