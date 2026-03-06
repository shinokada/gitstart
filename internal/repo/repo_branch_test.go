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

func TestDetectCurrentBranch_WorktreeFile(t *testing.T) {
	// Simulate a worktree/submodule where .git is a file containing a
	// "gitdir: <path>" pointer rather than being a directory.
	worktreeDir := t.TempDir()
	realGitDir := t.TempDir()

	// Write the real HEAD into the separate git directory.
	if err := os.WriteFile(filepath.Join(realGitDir, "HEAD"), []byte("ref: refs/heads/feature\n"), 0644); err != nil {
		t.Fatalf("could not write HEAD: %v", err)
	}

	// Write the .git file pointer in the worktree directory (absolute path).
	gitFile := filepath.Join(worktreeDir, ".git")
	if err := os.WriteFile(gitFile, []byte("gitdir: "+realGitDir+"\n"), 0644); err != nil {
		t.Fatalf("could not write .git file: %v", err)
	}

	if got := DetectCurrentBranch(worktreeDir); got != "feature" {
		t.Errorf("DetectCurrentBranch() = %q, want %q", got, "feature")
	}
}
