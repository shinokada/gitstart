package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/shinichiokada/gitstart/internal/prompts"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitstart",
	Short: "Automate GitHub repository creation",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if private && public {
			return fmt.Errorf("flags --private and --public are mutually exclusive")
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
			prompts.DryRunPrompt("Would create .gitignore for language if specified")
			prompts.DryRunPrompt("Would prompt for and create LICENSE file")
			prompts.DryRunPrompt("Would create README.md with project name and description")
			prompts.DryRunPrompt("Would initialize git repository if not present")
			prompts.DryRunPrompt("Would add all files and commit with message")
			prompts.DryRunPrompt("Would create GitHub repository (public/private as specified)")
			prompts.DryRunPrompt("Would add remote origin and push to branch")
			prompts.DryRunPrompt("Would handle existing files and directories as described in documentation")
			prompts.DryRunPrompt("No actions will be performed in dry-run mode.")
			return
		}
		// ...existing code for actual execution...
	},
}

var dryRun bool
var version = "dev" // overridden at build time via -ldflags
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show gitstart version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gitstart version", version)
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
