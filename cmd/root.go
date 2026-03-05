package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/shinokada/gitstart/internal/files"
	"github.com/shinokada/gitstart/internal/prompts"
	"github.com/shinokada/gitstart/internal/repo"
	"github.com/spf13/cobra"
)

var version = "dev"

func getVersion() string {
	if version != "dev" {
		return version
	}
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
		return info.Main.Version
	}
	return version
}

var rootCmd = &cobra.Command{
	Use:   "gitstart",
	Short: "Automate GitHub repository creation",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if private && public {
			return fmt.Errorf("flags --private and --public are mutually exclusive")
		}
		branch = strings.TrimSpace(branch)
		if branch == "" {
			return fmt.Errorf("flag --branch cannot be empty")
		}
		return nil
	},
	Long: `gitstart automates project setup, git init, and GitHub repo creation.

Basic Usage:
	gitstart -d repo-name
	cd existing_project && gitstart -d .

More examples:
	gitstart -d my-project
	gitstart -d my-python-app -l python
	gitstart -d secret-project -p
	gitstart -d my-app -m "First release" -b develop
	gitstart -d awesome-tool --description "An amazing CLI tool for developers"
	gitstart -d test-repo --dry-run
	gitstart -d automated-repo -q
	cd my-existing-project && gitstart -d . -l javascript --description "My existing JavaScript project"
`,
	Run: func(cmd *cobra.Command, args []string) {
		if directory == "" {
			_ = cmd.Help()
			return
		}
		if dryRun {
			prompts.DryRunPrompt("[OPTIONS]")
			prompts.DryRunPrompt("  Directory: " + directory)
			prompts.DryRunPrompt("  Language: " + language)
			prompts.DryRunPrompt("  Branch: " + branch)
			prompts.DryRunPrompt("  Commit message: " + message)
			prompts.DryRunPrompt("  Private: " + strconv.FormatBool(private))
			prompts.DryRunPrompt("  Public: " + strconv.FormatBool(public))
			prompts.DryRunPrompt("  Description: " + description)
			prompts.DryRunPrompt("  Quiet: " + strconv.FormatBool(quiet))
			prompts.DryRunPrompt("  Dry-run: true")
			prompts.DryRunPrompt("[ACTIONS]")
			prompts.DryRunPrompt("Would create project directory if needed")
			if strings.TrimSpace(language) != "" {
				prompts.DryRunPrompt("Would create .gitignore for language: " + language)
			} else {
				prompts.DryRunPrompt("Would skip .gitignore (no language specified)")
			}
			if quiet {
				prompts.DryRunPrompt("Would skip LICENSE creation (quiet mode)")
			} else {
				prompts.DryRunPrompt("Would prompt for and create LICENSE file")
			}
			prompts.DryRunPrompt("Would create README.md with project name and description")
			prompts.DryRunPrompt("Would initialize git repository if not present")
			prompts.DryRunPrompt("Would add all files and commit with message")
			prompts.DryRunPrompt("Would create GitHub repository (public/private as specified)")
			prompts.DryRunPrompt("Would add remote origin and push to branch")
			prompts.DryRunPrompt("No actions will be performed in dry-run mode.")
			return
		}
		if err := run(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	},
}

var dryRun bool
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show gitstart version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gitstart version", getVersion())
	},
}
var directory string
var language string
var branch string
var message string
var private bool
var public bool
var description string
var quiet bool

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().BoolVarP(&dryRun, "dry-run", "n", false, "Preview actions without making changes")
	rootCmd.PersistentFlags().StringVarP(&directory, "directory", "d", "", "Project directory name (use . for current directory)")
	rootCmd.PersistentFlags().StringVarP(&language, "language", "l", "", "Programming language for .gitignore")
	rootCmd.PersistentFlags().StringVarP(&branch, "branch", "b", "main", "Branch name (default: main)")
	rootCmd.PersistentFlags().StringVarP(&message, "message", "m", "Initial commit", "Commit message")
	rootCmd.PersistentFlags().BoolVarP(&private, "private", "p", false, "Create a private repository (default: public)")
	rootCmd.PersistentFlags().BoolVarP(&public, "public", "P", false, "Create a public repository")
	rootCmd.PersistentFlags().StringVar(&description, "description", "", "Repository description")
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "Minimal output")
}

func run() error {
	// Resolve directory to an absolute, clean path
	dir := directory
	if !filepath.IsAbs(dir) {
		wd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("could not get current directory: %w", err)
		}
		dir = filepath.Join(wd, dir)
	}
	dir = filepath.Clean(dir)
	repoName := filepath.Base(dir)

	if err := files.CreateProjectDir(dir); err != nil {
		return fmt.Errorf("could not create directory %q: %w", dir, err)
	}
	if !quiet {
		fmt.Printf("Setting up project %q in %s\n", repoName, dir)
	}

	if err := ensureGitignore(dir); err != nil {
		return err
	}
	if err := ensureLicense(dir); err != nil {
		return err
	}
	if err := ensureReadme(dir, repoName); err != nil {
		return err
	}
	if err := ensureGitRepo(dir); err != nil {
		return err
	}
	if err := createRemoteAndPush(dir, repoName); err != nil {
		return err
	}
	return nil
}

func ensureGitignore(dir string) error {
	lang := strings.TrimSpace(language)
	if lang == "" {
		if !quiet {
			fmt.Println("No language specified, skipping .gitignore creation.")
		}
		return nil
	}
	p := filepath.Join(dir, ".gitignore")
	if _, err := os.Stat(p); err == nil {
		if !quiet {
			fmt.Println(".gitignore already exists, skipping.")
		}
		return nil
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("could not access %s: %w", p, err)
	}
	if !quiet {
		fmt.Println("Creating .gitignore...")
	}
	if err := files.FetchGitignore(lang, p); err != nil {
		return err
	}
	return nil
}

func ensureLicense(dir string) error {
	p := filepath.Join(dir, "LICENSE")
	if _, err := os.Stat(p); err == nil {
		if !quiet {
			fmt.Println("LICENSE already exists, skipping.")
		}
		return nil
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("could not access %s: %w", p, err)
	}
	if quiet {
		// Skip interactive prompt in quiet/non-interactive mode
		return nil
	}
	licenseOptions := []string{
		"mit: Simple and permissive",
		"apache-2.0: Community-friendly",
		"gpl-3.0: Share improvements",
		"None",
	}
	choice := prompts.PromptSelect("Select a license:", licenseOptions)
	switch choice {
	case licenseOptions[0]:
		return files.FetchLicenseText("mit", p)
	case licenseOptions[1]:
		return files.FetchLicenseText("apache-2.0", p)
	case licenseOptions[2]:
		return files.FetchLicenseText("gpl-3.0", p)
	default:
		fmt.Println("Skipping LICENSE.")
	}
	return nil
}

func ensureReadme(dir, repoName string) error {
	p := filepath.Join(dir, "README.md")
	if _, err := os.Stat(p); err == nil {
		if !quiet {
			fmt.Println("README.md already exists, skipping.")
		}
		return nil
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("could not access %s: %w", p, err)
	}
	if !quiet {
		fmt.Println("Creating README.md...")
	}
	if err := files.CreateReadme(repoName, description, p); err != nil {
		return fmt.Errorf("could not create README.md: %w", err)
	}
	return nil
}

func ensureGitRepo(dir string) error {
	gitDir := filepath.Join(dir, ".git")
	if _, err := os.Stat(gitDir); err == nil {
		if !quiet {
			fmt.Println("Git repository already exists, skipping init.")
		}
		return nil
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("could not access %s: %w", gitDir, err)
	}
	if !quiet {
		fmt.Println("Initializing git repository...")
	}
	if err := repo.InitGitRepo(dir); err != nil {
		return fmt.Errorf("could not initialize git repository: %w", err)
	}
	return nil
}

func createRemoteAndPush(dir, repoName string) error {
	visibility := "public"
	if private {
		visibility = "private"
	}
	if !quiet {
		fmt.Printf("Creating GitHub repository %s...\n", repoName)
	}
	if err := repo.CreateGitHubRepo(dir, repoName, visibility, description); err != nil {
		return fmt.Errorf("could not create GitHub repository: %w", err)
	}
	if !quiet {
		fmt.Printf("Committing and pushing to branch %q...\n", branch)
	}
	if err := repo.CommitAndPush(dir, branch, message); err != nil {
		return fmt.Errorf("could not commit and push: %w", err)
	}
	if !quiet {
		ghUser := ghAuthenticatedUser()
		if ghUser != "" {
			fmt.Printf("✓ Done! Repository created: %s/%s\n", ghUser, repoName)
		} else {
			fmt.Printf("✓ Done! Repository: %s\n", repoName)
		}
	}
	return nil
}

// ghAuthenticatedUser returns the GitHub username of the currently authenticated gh CLI user.
// A 5-second timeout guards against the CLI hanging on network or auth issues.
func ghAuthenticatedUser() string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "gh", "api", "user", "--jq", ".login")
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
