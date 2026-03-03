package utils

import (
	"errors"
	"testing"
)

func TestErrorHandler(t *testing.T) {
	err := errors.New("test error")
	result := ErrorHandler(err, "test message")
	if result != err {
		t.Fatalf("expected original error to be returned, got: %v", result)
	}

	nilResult := ErrorHandler(nil, "should not print")
	if nilResult != nil {
		t.Fatalf("expected nil to be returned")
	}
}
