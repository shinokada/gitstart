package files

import (
	"os"
	"testing"
)

func TestCreateProjectDir(t *testing.T) {
	tmpDir := "test_tmp_dir"
	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			t.Fatalf("failed to remove temp dir: %v", err)
		}
	}()

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
