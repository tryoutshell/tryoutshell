package cmd

import (
	"encoding/json"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/tryoutshell/tryoutshell/internal/ui"
	"github.com/tryoutshell/tryoutshell/types"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Interactive list of organizations and lessons",
	Run: func(cmd *cobra.Command, args []string) {
		orgList := getOrganizationList()

		// Launch interactive TUI
		m := ui.NewListModel(orgList)
		p := tea.NewProgram(
			m,
			tea.WithAltScreen(),
		)

		finalModel, err := p.Run()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		// Check if user selected a lesson to start
		if listModel, ok := finalModel.(ui.ListModel); ok {
			if listModel.ShouldStartLesson() {
				orgID, lessonID := listModel.GetSelectedLesson()

				// Start the lesson
				if err := ui.LaunchInteractive(orgID, lessonID); err != nil {
					log.Fatalf("Error launching lesson: %v", err)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func getOrganizationList() []types.OrganizationDetails {
	filePath := "manifest.json"
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	var orgStruct types.OrganizationList
	err = json.Unmarshal(fileContent, &orgStruct)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}
	return orgStruct.Organizations
}
