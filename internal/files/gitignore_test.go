package files

import (
	"os"
	"testing"
)

func TestFetchGitignore(t *testing.T) {
	file := ".gitignore_test"
	defer func() {
		if err := os.Remove(file); err != nil {
			t.Fatalf("failed to remove test gitignore file: %v", err)
		}
	}()

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
