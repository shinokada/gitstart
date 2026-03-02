package utils

import (
	"errors"
	"testing"
)

func TestErrorHandler(t *testing.T) {
	err := errors.New("test error")
	result := ErrorHandler(err, "test message")
	if result == nil {
		t.Fatalf("expected error to be returned")
	}

	nilResult := ErrorHandler(nil, "should not print")
	if nilResult != nil {
		t.Fatalf("expected nil to be returned")
	}
}
