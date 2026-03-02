package utils

import (
	"fmt"
)

// ErrorHandler prints and returns an error.
func ErrorHandler(err error, msg string) error {
	if err != nil {
		fmt.Printf("ERROR: %s: %v\n", msg, err)
	}
	return err
}
