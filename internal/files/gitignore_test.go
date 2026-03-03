package files

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFetchGitignore(t *testing.T) {
	if os.Getenv("INTEGRATION") == "" {
		t.Skip("skipping network test; set INTEGRATION=1 to run")
	}

	dir := t.TempDir()
	file := filepath.Join(dir, ".gitignore_test")

	err := FetchGitignore("Go", file)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	info, err := os.Stat(file)
	if err != nil {
		t.Fatalf("expected .gitignore file to exist, got error: %v", err)
	}
	if info.Size() == 0 {
		t.Fatalf("expected .gitignore file to have content")
	}
}
