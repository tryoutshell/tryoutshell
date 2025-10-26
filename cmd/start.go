package cmd

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	lessons_pkg "github.com/tryoutshell/tryoutshell/internal/lessons"
	"github.com/tryoutshell/tryoutshell/internal/ui"
	"github.com/tryoutshell/tryoutshell/types"
)

var (
	lesson string
	theme  string
	color  string
)

var StartCmd = &cobra.Command{
	Use:   "start [org] [--lesson lesson-id]",
	Short: "Start an interactive learning session",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		startHandler(cmd, args)
	},
}

func startHandler(cmd *cobra.Command, args []string) {
	var selectedOrg string
	var selectedLesson string

	orgList := getOrganizationList()

	// Case 1: No args - show org selection
	if len(args) == 0 {
		selectedOrg = selectOrganization(orgList)
		if selectedOrg == "" {
			return // User quit
		}
	} else {
		// Case 2: Org provided
		selectedOrg = args[0]

		// Validate org exists
		orgExists := false
		for _, org := range orgList {
			if org.Id == selectedOrg {
				orgExists = true
				break
			}
		}

		if !orgExists {
			log.Fatalf("Organization '%s' not found", selectedOrg)
		}
	}

	// Get org details
	var orgDetail types.OrganizationDetails
	for _, org := range orgList {
		if org.Id == selectedOrg {
			orgDetail = org
			break
		}
	}

	// Case 3: Lesson flag provided
	if lesson != "" {
		selectedLesson = lesson
		lessonExists := false
		for _, l := range orgDetail.Lessons {
			if l == selectedLesson {
				lessonExists = true
				break
			}
		}

		if !lessonExists {
			log.Fatalf("Lesson '%s' not found in organization '%s'", selectedLesson, selectedOrg)
		}
	} else {
		// Case 4: Show lesson selection
		selectedLesson = selectLesson(selectedOrg, orgDetail.Name, orgDetail.Lessons)
		if selectedLesson == "" {
			// Instead of calling StartCmd.Run again — just call the function itself
			startHandler(cmd, []string{})
			return
		}
	}

	fmt.Printf("🚀 Loading lesson: %s/%s\n\n", selectedOrg, selectedLesson)

	if err := ui.LaunchInteractive(selectedOrg, selectedLesson); err != nil {
		log.Fatalf("Error launching UI: %v", err)
	}
}

func selectOrganization(orgs []types.OrganizationDetails) string {
	// Convert to OrgItem
	orgItems := make([]ui.OrgItem, len(orgs))
	for i, org := range orgs {
		orgItems[i] = ui.OrgItem{
			ID:          org.Id,
			Name:        org.Name,
			Description: org.Description,
			Logo:        org.Logo,
		}
	}

	m := ui.NewOrgListModel(orgItems)
	p := tea.NewProgram(m)

	finalModel, err := p.Run()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if orgModel, ok := finalModel.(ui.OrgListModel); ok {
		return orgModel.SelectedOrg()
	}

	return ""
}

func selectLesson(orgID, orgName string, lessonIDs []string) string {
	// Load lesson metadata
	lessons := lessons_pkg.GetAllLessonMetadata(orgID, lessonIDs)

	m := ui.NewLessonListModel(orgID, orgName, lessons)
	p := tea.NewProgram(m)

	finalModel, err := p.Run()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if lessonModel, ok := finalModel.(ui.LessonListModel); ok {
		if lessonModel.WasQuit() {
			return "" // User pressed q/esc
		}
		return lessonModel.SelectedLesson()
	}

	return ""
}

func init() {
	rootCmd.AddCommand(StartCmd)
	StartCmd.Flags().StringVarP(&lesson, "lesson", "l", "", "lesson ID to practice")
	StartCmd.Flags().StringVarP(&theme, "theme", "t", "default", "UI theme")
	StartCmd.Flags().StringVar(&color, "color", "", "custom color")
}
