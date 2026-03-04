package files

import (
	"os"
	"fmt"
)

// CreateReadme generates a README.md file with project name and description.
func CreateReadme(projectName, description, dest string) error {
	content := fmt.Sprintf("# %s\n\n%s\n", projectName, description)
	return os.WriteFile(dest, []byte(content), 0644)
}
