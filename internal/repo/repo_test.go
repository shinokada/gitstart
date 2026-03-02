package repo

import (
	"os"
	"testing"
)

func TestInitGitRepo(t *testing.T) {
	tmpDir := "test_git_repo"
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			t.Fatalf("failed to remove temp dir: %v", err)
		}
	}()

	err := InitGitRepo(tmpDir)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if _, err := os.Stat(tmpDir + "/.git"); err != nil {
		t.Fatalf("expected .git directory to exist, got error: %v", err)
	}
}
