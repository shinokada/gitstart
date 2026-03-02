package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitstart",
	Short: "Automate GitHub repository creation",
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
}

var dryRun bool
var version = "v1.0.0" // Update as needed
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
