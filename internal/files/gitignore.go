package files

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// languageMarkers maps known project marker filenames to a gitignore language.
// The first matching marker wins, so order within each entry does not matter
// but entries higher in the slice take priority if multiple languages match.
var languageMarkers = []struct {
	files    []string
	language string
}{
	{files: []string{"go.mod"}, language: "Go"},
	{files: []string{"Cargo.toml"}, language: "Rust"},
	{files: []string{"pubspec.yaml"}, language: "Dart"},
	{files: []string{"composer.json"}, language: "PHP"},
	{files: []string{"Gemfile"}, language: "Ruby"},
	{files: []string{"pom.xml", "build.gradle", "build.gradle.kts"}, language: "Java"},
	{files: []string{"requirements.txt", "pyproject.toml", "setup.py", "setup.cfg"}, language: "Python"},
	{files: []string{"package.json"}, language: "Node"},
}

// DetectLanguage inspects dir for well-known project marker files and returns
// the inferred gitignore language name, or "" if nothing is recognised.
func DetectLanguage(dir string) string {
	for _, entry := range languageMarkers {
		for _, marker := range entry.files {
			if _, err := os.Stat(filepath.Join(dir, marker)); err == nil {
				return entry.language
			}
		}
	}
	return ""
}

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
	url := fmt.Sprintf("https://raw.githubusercontent.com/github/gitignore/main/%s.gitignore", normalized)
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

	tmp, err := os.CreateTemp(filepath.Dir(dest), ".gitignore-*")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	defer func() { _ = os.Remove(tmpName) }()

	if _, err := io.Copy(tmp, resp.Body); err != nil {
		_ = tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmpName, dest)
}
