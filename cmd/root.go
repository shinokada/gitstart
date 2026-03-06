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
		message = strings.TrimSpace(message)
		if message == "" {
			return fmt.Errorf("flag --message cannot be empty")
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

After a framework starter:
	npx sv create my-app && cd my-app && gitstart -d . --post-framework
	npm create vite@latest my-app && cd my-app && gitstart -d . --post-framework
	composer create-project laravel/laravel my-app && cd my-app && gitstart -d . --post-framework
`,
	Run: func(cmd *cobra.Command, args []string) {
		if directory == "" {
			_ = cmd.Help()
			return
		}

		// Derive effective skip flags without mutating the Cobra-bound vars.
		// Mutating noLicense/noReadme directly would corrupt state across
		// multiple in-process Execute() calls (e.g. in tests).
		effNoLicense := noLicense || postFramework
		effNoReadme := noReadme || postFramework

		if dryRun {
			// Resolve directory for dry-run display.
			dir := resolveDir(directory)
			repoName := filepath.Base(dir)

			detectedLang := ""
			if language == "" {
				detectedLang = files.DetectLanguage(dir)
			}
			effectiveLang := language
			if effectiveLang == "" {
				effectiveLang = detectedLang
			}

			branchWasSetDry := cmd.Flags().Changed("branch")
			detectedBranch := ""
			if !branchWasSetDry {
				detectedBranch = repo.DetectCurrentBranch(dir)
			}
			effectiveBranch := branch
			if detectedBranch != "" {
				effectiveBranch = detectedBranch
			}

			prompts.DryRunPrompt("[OPTIONS]")
			prompts.DryRunPrompt("  Directory: " + dir)
			prompts.DryRunPrompt("  Repo name: " + repoName)
			if language != "" {
				prompts.DryRunPrompt("  Language (explicit): " + language)
			} else if detectedLang != "" {
				prompts.DryRunPrompt("  Language (auto-detected): " + detectedLang)
			} else {
				prompts.DryRunPrompt("  Language: (none)")
			}
			prompts.DryRunPrompt("  Branch: " + effectiveBranch)
			prompts.DryRunPrompt("  Commit message: " + message)
			prompts.DryRunPrompt("  Private: " + strconv.FormatBool(private))
			prompts.DryRunPrompt("  Public: " + strconv.FormatBool(public))
			prompts.DryRunPrompt("  Description: " + description)
			prompts.DryRunPrompt("  Quiet: " + strconv.FormatBool(quiet))
			prompts.DryRunPrompt("  No-license: " + strconv.FormatBool(effNoLicense))
			prompts.DryRunPrompt("  No-readme: " + strconv.FormatBool(effNoReadme))
			prompts.DryRunPrompt("  Post-framework: " + strconv.FormatBool(postFramework))
			prompts.DryRunPrompt("  Dry-run: true")
			prompts.DryRunPrompt("[ACTIONS]")
			prompts.DryRunPrompt("Would create project directory if needed")
			if effectiveLang != "" {
				prompts.DryRunPrompt("Would create .gitignore for language: " + effectiveLang)
			} else {
				prompts.DryRunPrompt("Would skip .gitignore (no language specified or detected)")
			}
			if effNoLicense {
				prompts.DryRunPrompt("Would skip LICENSE creation (--no-license / --post-framework)")
			} else if quiet {
				prompts.DryRunPrompt("Would skip LICENSE creation (quiet mode)")
			} else {
				prompts.DryRunPrompt("Would prompt for and create LICENSE file")
			}
			if effNoReadme {
				prompts.DryRunPrompt("Would skip README.md creation (--no-readme / --post-framework)")
			} else {
				prompts.DryRunPrompt("Would create README.md with project name and description")
			}
			prompts.DryRunPrompt("Would initialize git repository if not present")
			prompts.DryRunPrompt("Would add all files and commit with message")
			prompts.DryRunPrompt("Would create GitHub repository (public/private as specified)")
			prompts.DryRunPrompt("Would add remote origin and push to branch: " + effectiveBranch)
			prompts.DryRunPrompt("No actions will be performed in dry-run mode.")
			return
		}

		// Derive whether --branch was explicitly set locally so the value
		// is never stale across repeated in-process Execute() calls.
		branchWasSet := cmd.Flags().Changed("branch")

		if err := run(effNoLicense, effNoReadme, branchWasSet); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	},
}

var (
	dryRun        bool
	directory     string
	language      string
	branch        string
	message       string
	private       bool
	public        bool
	description   string
	quiet         bool
	noLicense     bool
	noReadme      bool
	postFramework bool
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show gitstart version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gitstart version", getVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().BoolVarP(&dryRun, "dry-run", "n", false, "Preview actions without making changes")
	rootCmd.PersistentFlags().StringVarP(&directory, "directory", "d", "", "Project directory name or path (use . for current directory)")
	rootCmd.PersistentFlags().StringVarP(&language, "language", "l", "", "Programming language for .gitignore (auto-detected if omitted)")
	rootCmd.PersistentFlags().StringVarP(&branch, "branch", "b", "main", "Branch name (default: main; auto-detected from existing repo if not set)")
	rootCmd.PersistentFlags().StringVarP(&message, "message", "m", "Initial commit", "Commit message")
	rootCmd.PersistentFlags().BoolVarP(&private, "private", "p", false, "Create a private repository (default: public)")
	rootCmd.PersistentFlags().BoolVarP(&public, "public", "P", false, "Create a public repository")
	rootCmd.PersistentFlags().StringVar(&description, "description", "", "Repository description")
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "Minimal output")
	rootCmd.PersistentFlags().BoolVar(&noLicense, "no-license", false, "Skip LICENSE file creation")
	rootCmd.PersistentFlags().BoolVar(&noReadme, "no-readme", false, "Skip README.md creation")
	rootCmd.PersistentFlags().BoolVar(&postFramework, "post-framework", false, "Optimised for use after a framework starter (implies --no-license --no-readme)")
}

// resolveDir converts the directory flag value to an absolute, clean path.
func resolveDir(dir string) string {
	if filepath.IsAbs(dir) {
		return filepath.Clean(dir)
	}
	wd, err := os.Getwd()
	if err != nil {
		return filepath.Clean(dir)
	}
	return filepath.Clean(filepath.Join(wd, dir))
}

func run(effNoLicense, effNoReadme, branchWasSet bool) error {
	dir := resolveDir(directory)
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
	if err := ensureLicense(dir, effNoLicense); err != nil {
		return err
	}
	if err := ensureReadme(dir, repoName, effNoReadme); err != nil {
		return err
	}
	// resolvedBranch is determined before git init so that a newly-created
	// repo always gets the exact branch the user expects (or the default).
	// This keeps dry-run output consistent with the actual run.
	resolvedBranch := effectiveBranch(dir, branchWasSet)
	if err := ensureGitRepo(dir, resolvedBranch); err != nil {
		return err
	}
	if err := createRemoteAndPush(dir, repoName, resolvedBranch); err != nil {
		return err
	}
	return nil
}

func ensureGitignore(dir string) error {
	p := filepath.Join(dir, ".gitignore")

	// If .gitignore already exists, always skip — even if we would have
	// auto-detected a language — to avoid overwriting framework-generated files.
	if _, err := os.Stat(p); err == nil {
		if !quiet {
			fmt.Println(".gitignore already exists, skipping.")
		}
		return nil
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("could not access %s: %w", p, err)
	}

	lang := strings.TrimSpace(language)

	// Auto-detect if no language was explicitly provided.
	if lang == "" {
		detected := files.DetectLanguage(dir)
		if detected != "" {
			if !quiet {
				fmt.Printf("Auto-detected language: %s. Creating .gitignore...\n", detected)
			}
			lang = detected
		}
	}

	if lang == "" {
		if !quiet {
			fmt.Println("No language specified or detected, skipping .gitignore creation.")
		}
		return nil
	}

	if !quiet {
		fmt.Println("Creating .gitignore...")
	}
	return files.FetchGitignore(lang, p)
}

func ensureLicense(dir string, effNoLicense bool) error {
	// Respect --no-license (also set by --post-framework).
	if effNoLicense {
		if !quiet {
			fmt.Println("Skipping LICENSE creation (--no-license).")
		}
		return nil
	}

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
		// Skip interactive prompt in quiet/non-interactive mode.
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

func ensureReadme(dir, repoName string, effNoReadme bool) error {
	// Respect --no-readme (also set by --post-framework).
	if effNoReadme {
		if !quiet {
			fmt.Println("Skipping README.md creation (--no-readme).")
		}
		return nil
	}

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

func ensureGitRepo(dir, branch string) error {
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
	if err := repo.InitGitRepo(dir, branch); err != nil {
		return fmt.Errorf("could not initialize git repository: %w", err)
	}
	return nil
}

// effectiveBranch returns the branch to push to. branchWasSet reports whether
// the user explicitly passed --branch on this invocation. If not, and a
// pre-existing repo is found, we read the active branch from git so we honour
// whatever the framework starter set. For newly-created repos there is no HEAD
// yet, so we fall back to the branch flag default and pass it to git init.
func effectiveBranch(dir string, branchWasSet bool) string {
	if !branchWasSet {
		if detected := repo.DetectCurrentBranch(dir); detected != "" {
			return detected
		}
	}
	return branch
}

func createRemoteAndPush(dir, repoName, pushBranch string) error {
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
		fmt.Printf("Committing and pushing to branch %q...\n", pushBranch)
	}
	if err := repo.CommitAndPush(dir, pushBranch, message); err != nil {
		// Clean up the orphaned remote repo so the user can retry cleanly.
		if cleanupErr := repo.DeleteGitHubRepo(repoName); cleanupErr != nil {
			fmt.Fprintf(os.Stderr, "warning: could not delete orphaned repository %q: %v\n", repoName, cleanupErr)
		} else if !quiet {
			fmt.Fprintf(os.Stderr, "note: deleted orphaned repository %q after push failure\n", repoName)
		}
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

// ghAuthenticatedUser returns the GitHub username of the currently
// authenticated gh CLI user. A 5-second timeout guards against the CLI
// hanging on network or auth issues.
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
