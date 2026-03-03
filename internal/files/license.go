package files

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// FetchLicenseText fetches license text from GitHub API and writes to LICENSE file.
func FetchLicenseText(license string, dest string) error {
	url := fmt.Sprintf("https://api.github.com/licenses/%s", license)
	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		msg, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
		return fmt.Errorf("failed to fetch license %q: status %d: %s", license, resp.StatusCode, strings.TrimSpace(string(msg)))
	}

	var data struct {
		Body string `json:"body"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}
	if data.Body == "" {
		return fmt.Errorf("license %q returned empty body", license)
	}

	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	_, err = io.WriteString(f, data.Body)
	return err
}
