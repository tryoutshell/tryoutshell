package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tryoutshell/tryoutshell/internal/reader"
)

var savedCmd = &cobra.Command{
	Use:   "saved",
	Short: "List and open saved articles",
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Error getting home directory: %v", err)
		}

		saveDir := filepath.Join(homeDir, ".local", "share", "tryoutshell", "saved")
		entries, err := os.ReadDir(saveDir)
		if err != nil {
			fmt.Println("No saved articles found.")
			fmt.Printf("Save articles with: tryoutshell read <url> --save\n")
			return
		}

		var files []string
		for _, e := range entries {
			if !e.IsDir() && strings.HasSuffix(e.Name(), ".md") {
				files = append(files, e.Name())
			}
		}

		if len(files) == 0 {
			fmt.Println("No saved articles found.")
			return
		}

		fmt.Println("\n  📚 Saved Articles")
		fmt.Println()

		query := ""
		if len(args) > 0 {
			query = strings.ToLower(strings.Join(args, " "))
		}

		var matched []string
		for _, f := range files {
			name := strings.TrimSuffix(f, ".md")
			if query == "" || strings.Contains(strings.ToLower(name), query) {
				matched = append(matched, f)
			}
		}

		if len(matched) == 0 {
			fmt.Printf("  No articles matching '%s'\n\n", query)
			return
		}

		for i, f := range matched {
			name := strings.TrimSuffix(f, ".md")
			fmt.Printf("  %d. %s\n", i+1, name)
		}
		fmt.Println()

		if len(matched) == 1 {
			openSavedArticle(filepath.Join(saveDir, matched[0]))
			return
		}

		fmt.Print("  Enter number to open (or q to quit): ")
		var input string
		if _, err := fmt.Scanln(&input); err != nil {
			return
		}

		if input == "q" || input == "" {
			return
		}

		var idx int
		if _, err := fmt.Sscanf(input, "%d", &idx); err != nil || idx < 1 || idx > len(matched) {
			fmt.Println("  Invalid selection.")
			return
		}

		openSavedArticle(filepath.Join(saveDir, matched[idx-1]))
	},
}

func openSavedArticle(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading saved article: %v", err)
	}

	content := string(data)
	title := strings.TrimSuffix(filepath.Base(path), ".md")
	url := ""

	if strings.HasPrefix(content, "---\n") {
		end := strings.Index(content[4:], "\n---\n")
		if end >= 0 {
			frontmatter := content[4 : 4+end]
			for _, line := range strings.Split(frontmatter, "\n") {
				if strings.HasPrefix(line, "title: ") {
					title = strings.TrimPrefix(line, "title: ")
				}
				if strings.HasPrefix(line, "url: ") {
					url = strings.TrimPrefix(line, "url: ")
				}
			}
			content = content[4+end+5:]
		}
	}

	if err := reader.Launch(title, content, url); err != nil {
		log.Fatalf("Error launching reader: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(savedCmd)
}
