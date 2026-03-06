package repo

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCommitAndPush(t *testing.T) {
	dir := t.TempDir()

	// Initialize git repo
	if err := InitGitRepo(dir, "main"); err != nil {
		t.Fatalf("failed to init git repo: %v", err)
	}

	// Create a file to commit
	file := filepath.Join(dir, "test.txt")
	if err := os.WriteFile(file, []byte("test"), 0644); err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	// Without a remote configured, CommitAndPush must fail on push.
	err := CommitAndPush(dir, "main", "test commit")
	if err == nil {
		t.Fatalf("expected CommitAndPush to fail without an origin remote")
	}
}
