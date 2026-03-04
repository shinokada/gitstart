package files

import (
	"os"
)

// CreateProjectDir creates a new project directory if it doesn't exist.
func CreateProjectDir(path string) error {
	return os.MkdirAll(path, 0755)
}
