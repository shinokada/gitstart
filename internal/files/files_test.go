package files

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateProjectDir(t *testing.T) {
	baseDir := t.TempDir()
	tmpDir := filepath.Join(baseDir, "test_project")

	err := CreateProjectDir(tmpDir)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	info, err := os.Stat(tmpDir)
	if err != nil {
		t.Fatalf("expected directory to exist, got error: %v", err)
	}
	if !info.IsDir() {
		t.Fatalf("expected a directory, got something else")
	}
}
