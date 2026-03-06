package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"github.com/tryoutshell/tryoutshell/internal/reader"
)

var saveArticle bool

var readCmd = &cobra.Command{
	Use:   "read <url>",
	Short: "Read a blog post or article in the terminal",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]

		fmt.Printf("Fetching article from %s...\n", url)
		title, markdown, err := reader.FetchArticle(url)
		if err != nil {
			log.Fatalf("Error fetching article: %v", err)
		}

		if saveArticle {
			if err := saveArticleToFile(title, markdown, url); err != nil {
				fmt.Printf("Warning: could not save article: %v\n", err)
			} else {
				fmt.Println("Article saved.")
			}
		}

		if err := reader.Launch(title, markdown, url); err != nil {
			log.Fatalf("Error launching reader: %v", err)
		}
	},
}

func saveArticleToFile(title, markdown, url string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	saveDir := filepath.Join(homeDir, ".local", "share", "tryoutshell", "saved")
	if err := os.MkdirAll(saveDir, 0755); err != nil {
		return err
	}

	slug := sanitizeFilename(title)
	if slug == "" {
		slug = fmt.Sprintf("article-%d", time.Now().Unix())
	}

	content := fmt.Sprintf("---\ntitle: %s\nurl: %s\nsaved: %s\n---\n\n%s",
		title, url, time.Now().Format(time.RFC3339), markdown)

	return os.WriteFile(filepath.Join(saveDir, slug+".md"), []byte(content), 0644)
}

func sanitizeFilename(s string) string {
	var result []byte
	for _, c := range []byte(s) {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '-' || c == '_' {
			result = append(result, c)
		} else if c == ' ' {
			result = append(result, '-')
		}
	}
	if len(result) > 80 {
		result = result[:80]
	}
	return string(result)
}

func init() {
	rootCmd.AddCommand(readCmd)
	readCmd.Flags().BoolVar(&saveArticle, "save", false, "Save article for offline reading")
}
