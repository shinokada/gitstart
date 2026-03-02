package config

import (
	"os"
	"path/filepath"
)

// ConfigFilePath returns the path to the config file in the user's config directory.
func ConfigFilePath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "gitstart", "config"), nil
}

// SaveUsername stores the GitHub username in the config file.
func SaveUsername(username string) error {
	path, err := ConfigFilePath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(username), 0644)
}

// LoadUsername retrieves the GitHub username from the config file.
func LoadUsername() (string, error) {
	path, err := ConfigFilePath()
	if err != nil {
		return "", err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
