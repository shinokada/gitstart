package files

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateReadme(t *testing.T) {
	file := filepath.Join(t.TempDir(), "README_test.md")

	err := CreateReadme("TestProject", "Test description.", file)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	content, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("expected README file to exist: %v", err)
	}
	expected := "# TestProject\n\nTest description.\n"
	if string(content) != expected {
		t.Fatalf("unexpected content: got %q, want %q", string(content), expected)
	}
}
