package files

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// FetchLicenseText fetches license text from GitHub API and writes to LICENSE file.
func FetchLicenseText(license string, dest string) error {
	url := fmt.Sprintf("https://api.github.com/licenses/%s", license)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var data struct {
		Body string `json:"body"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
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
