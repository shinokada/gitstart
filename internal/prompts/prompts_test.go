package prompts

import (
	"io"
	"os"
	"testing"
)

func TestDryRunPrompt(t *testing.T) {
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create stdout pipe: %v", err)
	}
	os.Stdout = w
	t.Cleanup(func() { os.Stdout = old })

	DryRunPrompt("Create project directory")

	if err := w.Close(); err != nil {
		t.Fatalf("failed to close pipe writer: %v", err)
	}
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("failed to read captured stdout: %v", err)
	}

	want := "[DRY RUN] Would perform: Create project directory\n"
	if string(out) != want {
		t.Fatalf("unexpected output: got %q, want %q", string(out), want)
	}
}

func TestPromptUser(t *testing.T) {
	t.Skip("interactive prompt: requires mocked stdin/stdout to test")
}

func TestPromptSelect(t *testing.T) {
	t.Skip("interactive prompt: requires mocked stdin/stdout to test")
}
