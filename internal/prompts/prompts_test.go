package prompts

import (
	"testing"
)

func TestDryRunPrompt(t *testing.T) {
	// This test just ensures the function runs without error.
	DryRunPrompt("Create project directory")

}

func TestPromptUser(t *testing.T) {
	t.Skip("interactive prompt: requires mocked stdin/stdout to test")
}

func TestPromptSelect(t *testing.T) {
	t.Skip("interactive prompt: requires mocked stdin/stdout to test")
}
