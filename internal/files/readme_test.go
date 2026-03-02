package files

import (
	"os"
	"testing"
)

func TestCreateReadme(t *testing.T) {
	file := "README_test.md"
	defer func() {
		if err := os.Remove(file); err != nil {
			t.Fatalf("failed to remove test README file: %v", err)
		}
	}()

	err := CreateReadme("TestProject", "Test description.", file)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	info, err := os.Stat(file)
	if err != nil {
		t.Fatalf("expected README file to exist, got error: %v", err)
	}
	if info.Size() == 0 {
		t.Fatalf("expected README file to have content")
	}
}
