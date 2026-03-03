package files

import (
	"bytes"
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

	data, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("expected .gitignore file to exist, got error: %v", err)
	}
	if len(bytes.TrimSpace(data)) == 0 {
		t.Fatalf("expected .gitignore file to have content")
	}
	fallback := ".DS_Store\n.idea/\n.vscode/\n*.swp\n"
	if string(data) == fallback {
		t.Fatalf("expected fetched Go template, got fallback content")
	}
}
