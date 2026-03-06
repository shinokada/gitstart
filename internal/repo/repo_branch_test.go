package repo

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectCurrentBranch(t *testing.T) {
	tests := []struct {
		name     string
		head     string // content of .git/HEAD; empty means don't create the file
		expected string
	}{
		{
			name:     "standard branch",
			head:     "ref: refs/heads/main\n",
			expected: "main",
		},
		{
			name:     "non-default branch",
			head:     "ref: refs/heads/develop\n",
			expected: "develop",
		},
		{
			name:     "branch with slashes",
			head:     "ref: refs/heads/feat/my-feature\n",
			expected: "feat/my-feature",
		},
		{
			name:     "detached HEAD returns empty string",
			head:     "abc123def456abc123def456abc123def456abc12\n",
			expected: "",
		},
		{
			name:     "no .git directory returns empty string",
			head:     "", // will not create .git/HEAD
			expected: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dir := t.TempDir()

			if tc.head != "" {
				gitDir := filepath.Join(dir, ".git")
				if err := os.MkdirAll(gitDir, 0755); err != nil {
					t.Fatalf("could not create .git dir: %v", err)
				}
				if err := os.WriteFile(filepath.Join(gitDir, "HEAD"), []byte(tc.head), 0644); err != nil {
					t.Fatalf("could not write .git/HEAD: %v", err)
				}
			}

			got := DetectCurrentBranch(dir)
			if got != tc.expected {
				t.Errorf("DetectCurrentBranch() = %q, want %q", got, tc.expected)
			}
		})
	}
}
