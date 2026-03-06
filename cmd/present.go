package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tryoutshell/tryoutshell/internal/ui"
)

var presentCmd = &cobra.Command{
	Use:   "present <file.md>",
	Short: "Present a markdown file as terminal slides",
	Long: `Present a markdown file as a full-screen terminal presentation.

Slides are separated by '---' (a horizontal rule on its own line).
Each slide supports standard markdown: headings, lists, code blocks,
blockquotes, bold/italic text, and links.

Navigation:
  space / → / ↓ / enter / n / j / l    next slide
  ← / ↑ / p / h / k / N               previous slide
  gg                                   first slide
  G                                    last slide
  <number> G                           jump to slide <number>
  /                                    search slides (Enter to jump, ctrl+n for next result)
  ctrl+u / ctrl+d                      scroll up / down within a slide
  q / ctrl+c                           quit

Example:
  tryoutshell present my-blog-post.md
  tryoutshell present docs/intro.md`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := ui.LaunchPresentation(args[0]); err != nil {
			log.Fatalf("Error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(presentCmd)
}
