package files

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// languageAliases maps common lowercase inputs to the exact filename used in
// github.com/github/gitignore (without the .gitignore extension).
var languageAliases = map[string]string{
	"javascript": "Node",
	"js":         "Node",
	"typescript": "Node",
	"ts":         "Node",
	"golang":     "Go",
	"py":         "Python",
	"rb":         "Ruby",
	"rs":         "Rust",
	"cs":         "CSharp",
	"csharp":     "CSharp",
	"c++":        "C++",
	"cpp":        "C++",
	"sh":         "Shell",
	"bash":       "Shell",
}

// NormalizeLanguage converts a user-supplied language string to the exact
// casing used by the github/gitignore repository.
func NormalizeLanguage(lang string) string {
	lower := strings.ToLower(strings.TrimSpace(lang))
	if alias, ok := languageAliases[lower]; ok {
		return alias
	}
	// Title-case as a best-effort for everything else (e.g. "python" → "Python")
	if len(lower) > 0 {
		return strings.ToUpper(lower[:1]) + lower[1:]
	}
	return ""
}

// FetchGitignore downloads a language-specific .gitignore template from the
// github/gitignore repository. Returns an error if the template is not found
// rather than silently falling back to a minimal file.
func FetchGitignore(language, dest string) error {
	normalized := NormalizeLanguage(language)
	if normalized == "" {
		return fmt.Errorf("language cannot be empty")
	}
	url := fmt.Sprintf("https://raw.githubusercontent.com/github/gitignore/master/%s.gitignore", normalized)
	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("could not fetch .gitignore template for %q: %w", language, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("no .gitignore template found for language %q (tried %q)", language, normalized)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status %d fetching .gitignore template for %q", resp.StatusCode, language)
	}

	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()
	_, err = io.Copy(f, resp.Body)
	return err
}
