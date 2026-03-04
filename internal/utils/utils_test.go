package utils

import (
	"errors"
	"io"
	"os"
	"strings"
	"testing"
)

func captureStderr(t *testing.T, fn func()) string {
	t.Helper()
	orig := os.Stderr
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	os.Stderr = w
	fn()
	_ = w.Close()
	os.Stderr = orig
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("failed to read stderr: %v", err)
	}
	return string(out)
}

func TestErrorHandler(t *testing.T) {
	err := errors.New("test error")
	stderrOut := captureStderr(t, func() {
		result := ErrorHandler(err, "test message")
		if result != err {
			t.Fatalf("expected original error to be returned, got: %v", result)
		}
	})
	if !strings.Contains(stderrOut, "ERROR: test message: test error") {
		t.Fatalf("expected stderr to contain formatted error, got: %q", stderrOut)
	}

	nilStderrOut := captureStderr(t, func() {
		nilResult := ErrorHandler(nil, "should not print")
		if nilResult != nil {
			t.Fatalf("expected nil to be returned")
		}
	})
	if nilStderrOut != "" {
		t.Fatalf("expected no stderr output for nil error, got: %q", nilStderrOut)
	}
}
