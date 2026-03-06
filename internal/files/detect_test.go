package files

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectLanguage(t *testing.T) {
	tests := []struct {
		name     string
		markers  []string // files to create in the temp dir
		expected string
	}{
		{
			name:     "Go project",
			markers:  []string{"go.mod"},
			expected: "Go",
		},
		{
			name:     "Node project via package.json",
			markers:  []string{"package.json"},
			expected: "Node",
		},
		{
			name:     "Python project via requirements.txt",
			markers:  []string{"requirements.txt"},
			expected: "Python",
		},
		{
			name:     "Python project via pyproject.toml",
			markers:  []string{"pyproject.toml"},
			expected: "Python",
		},
		{
			name:     "Rust project",
			markers:  []string{"Cargo.toml"},
			expected: "Rust",
		},
		{
			name:     "Composer (PHP) project",
			markers:  []string{"composer.json"},
			expected: "Composer",
		},
		{
			name:     "Ruby project",
			markers:  []string{"Gemfile"},
			expected: "Ruby",
		},
		{
			name:     "Dart project",
			markers:  []string{"pubspec.yaml"},
			expected: "Dart",
		},
		{
			name:     "Java project via pom.xml",
			markers:  []string{"pom.xml"},
			expected: "Java",
		},
		{
			name:     "Java project via build.gradle",
			markers:  []string{"build.gradle"},
			expected: "Java",
		},
		{
			name:     "empty directory returns empty string",
			markers:  []string{},
			expected: "",
		},
		{
			name:     "unknown markers return empty string",
			markers:  []string{"CMakeLists.txt", "main.cpp"},
			expected: "",
		},
		{
			// Go takes priority over Node in the marker list ordering.
			name:     "Go wins over Node when both markers present",
			markers:  []string{"go.mod", "package.json"},
			expected: "Go",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dir := t.TempDir()
			for _, f := range tc.markers {
				if err := os.WriteFile(filepath.Join(dir, f), []byte(""), 0644); err != nil {
					t.Fatalf("could not create marker file %s: %v", f, err)
				}
			}
			got := DetectLanguage(dir)
			if got != tc.expected {
				t.Errorf("DetectLanguage() = %q, want %q", got, tc.expected)
			}
		})
	}
}

func TestDetectLanguage_NonExistentDir(t *testing.T) {
	got := DetectLanguage("/non/existent/directory/xyz123")
	if got != "" {
		t.Errorf("expected empty string for non-existent dir, got %q", got)
	}
}
