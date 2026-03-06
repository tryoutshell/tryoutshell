package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tryoutshell",
	Short: "Interactive terminal-based learning tool for DevSecOps and developer topics",
	Long: `TryOutShell — Interactive Learning in Your Terminal

Explore security, containers, CI/CD, and developer tools with hands-on
lessons, quizzes, and an AI-powered blog reader — all from your terminal.

Commands:
  start         Start an interactive learning session
  list          Browse available organizations and lessons
  present       Present a markdown file as terminal slides
  quiz          Launch quiz mode for a lesson
  read          Read a blog/article in a split-pane TUI with AI chat
  saved         List and open saved articles
  progress      Show your learning progress
  update        Check for and download lesson updates
  completion    Generate shell completion scripts`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
