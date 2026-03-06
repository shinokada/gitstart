package repo

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInitGitRepo(t *testing.T) {
	tmpDir := t.TempDir()

	err := InitGitRepo(tmpDir, "main")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if _, err := os.Stat(filepath.Join(tmpDir, ".git")); err != nil {
		t.Fatalf("expected .git directory to exist, got error: %v", err)
	}
}
