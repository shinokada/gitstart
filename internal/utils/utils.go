package utils

import (
	"fmt"
	"os"
)

// ErrorHandler prints an error message to stderr and returns the error.
func ErrorHandler(err error, msg string) error {
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s: %v\n", msg, err)
	}
	return err
}
