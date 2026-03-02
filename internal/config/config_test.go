package config

import (
	"os"
	"testing"
)

func TestSaveAndLoadUsername(t *testing.T) {
	username := "testuser"
	if err := SaveUsername(username); err != nil {
		t.Fatalf("failed to save username: %v", err)
	}

	loaded, err := LoadUsername()
	if err != nil {
		t.Fatalf("failed to load username: %v", err)
	}
	if loaded != username {
		t.Fatalf("expected %s, got %s", username, loaded)
	}

	// Cleanup
	path, err := ConfigFilePath()
	if err == nil {
		_ = os.Remove(path)
	}
}
