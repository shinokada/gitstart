package config

import (
	"testing"
)

func TestSaveAndLoadUsername(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", tmp)
	t.Setenv("APPDATA", tmp) // Windows fallback

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

}
