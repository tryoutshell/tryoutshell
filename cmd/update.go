package cmd

import (
	"fmt"
	"log"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/tryoutshell/tryoutshell/internal/updater"
)

var checkOnly bool

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Check for and download lesson updates",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Checking for updates...")

		updates, err := updater.CheckForUpdates()
		if err != nil {
			log.Fatalf("Error checking for updates: %v", err)
		}

		if len(updates) == 0 {
			fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("42")).Render("✓ All lessons are up to date."))
			return
		}

		headerStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63"))
		newStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("42"))
		updateStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("226"))

		fmt.Printf("\n%s\n\n", headerStyle.Render(fmt.Sprintf("  %d update(s) available:", len(updates))))

		for _, u := range updates {
			if u.OldVersion == "" {
				fmt.Printf("  %s %s/%s (%s)\n", newStyle.Render("NEW"), u.Org, u.Lesson, u.NewVersion)
			} else {
				fmt.Printf("  %s %s/%s (%s → %s)\n", updateStyle.Render("UPD"), u.Org, u.Lesson, u.OldVersion, u.NewVersion)
			}
		}

		if checkOnly {
			fmt.Println("\n  Run 'tryoutshell update' to download.")
			return
		}

		fmt.Println("\nDownloading updates...")
		if err := updater.DownloadUpdates(updates); err != nil {
			log.Fatalf("Error downloading updates: %v", err)
		}

		fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("42")).Render("\n✓ Updates downloaded successfully."))
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().BoolVar(&checkOnly, "check", false, "Only check for updates without downloading")
}
