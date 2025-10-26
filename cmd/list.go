package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tryoutshell/tryoutshell/internal/ui"
	"github.com/tryoutshell/tryoutshell/types"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Interactive list of organizations and lessons",
	Run: func(cmd *cobra.Command, args []string) {
		orgList := getOrganizationList()

		// Show org selection
		selectedOrg := selectOrganization(orgList)
		if selectedOrg == "" {
			return
		}

		// Get org details
		var orgDetail types.OrganizationDetails
		for _, org := range orgList {
			if org.Id == selectedOrg {
				orgDetail = org
				break
			}
		}

		// Show lesson selection
		selectedLesson := selectLesson(selectedOrg, orgDetail.Name, orgDetail.Lessons)
		if selectedLesson == "" {
			return
		}

		// Start the lesson
		if err := ui.LaunchInteractive(selectedOrg, selectedLesson); err != nil {
			log.Fatalf("Error launching lesson: %v", err)
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
