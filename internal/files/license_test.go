package files

import (
	"os"
	"testing"
)

func TestFetchLicenseText(t *testing.T) {
	file := "LICENSE_test"
	defer func() {
		if err := os.Remove(file); err != nil {
			t.Fatalf("failed to remove test license file: %v", err)
		}
	}()

	err := FetchLicenseText("mit", file)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	info, err := os.Stat(file)
	if err != nil {
		t.Fatalf("expected LICENSE file to exist, got error: %v", err)
	}
	if info.Size() == 0 {
		t.Fatalf("expected LICENSE file to have content")
	}
}