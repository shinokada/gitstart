package files

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// FetchGitignore downloads a language-specific .gitignore template from GitHub or creates a minimal one.
func FetchGitignore(language, dest string) error {
	if language != "" {
		url := fmt.Sprintf("https://raw.githubusercontent.com/github/gitignore/master/%s.gitignore", language)
		client := &http.Client{Timeout: 15 * time.Second}
		resp, err := client.Get(url)
		if err == nil {
			defer func() { _ = resp.Body.Close() }()
		}
		if err == nil && resp.StatusCode == http.StatusOK {
			f, err := os.Create(dest)
			if err != nil {
				return err
			}
			defer func() { _ = f.Close() }()
			_, err = io.Copy(f, resp.Body)
			return err
		}
	}
	// Fallback to minimal .gitignore
	minimal := ".DS_Store\n.idea/\n.vscode/\n*.swp\n"
	return os.WriteFile(dest, []byte(minimal), 0644)
}
