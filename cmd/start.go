package cmd

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/tryoutshell/tryoutshell/internal/lesson"
	lessons_pkg "github.com/tryoutshell/tryoutshell/internal/lessons"
	"github.com/tryoutshell/tryoutshell/internal/ui"
	"github.com/tryoutshell/tryoutshell/types"
)

var (
	lessonFlag string
	themeFlag  string
	colorFlag  string
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
	allLessons := lesson.DiscoverLessons()
	orgList := buildOrgList(allLessons)

	var selectedOrg string
	var selectedLesson string

	if len(args) == 0 {
		selectedOrg = selectOrganization(orgList)
		if selectedOrg == "" {
			return
		}
	} else {
		selectedOrg = args[0]

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

	var orgDetail types.OrganizationDetails
	for _, org := range orgList {
		if org.Id == selectedOrg {
			orgDetail = org
			break
		}
	}

	if lessonFlag != "" {
		selectedLesson = lessonFlag
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
		selectedLesson = selectLesson(selectedOrg, orgDetail.Name, orgDetail.Lessons)
		if selectedLesson == "" {
			startHandler(cmd, []string{})
			return
		}
	}

	fmt.Printf("🚀 Loading lesson: %s/%s\n\n", selectedOrg, selectedLesson)

	dl := lesson.FindLesson(allLessons, selectedOrg, selectedLesson)
	if dl != nil && dl.HasSlides && !dl.HasLegacy {
		slidesContent, err := lesson.LoadSlides(dl.Dir)
		if err != nil {
			log.Fatalf("Error loading slides: %v", err)
		}
		if err := ui.LaunchSlideLesson(dl, slidesContent); err != nil {
			log.Fatalf("Error launching lesson: %v", err)
		}
		return
	}

	if err := ui.LaunchInteractive(selectedOrg, selectedLesson); err != nil {
		log.Fatalf("Error launching UI: %v", err)
	}
}

func selectOrganization(orgs []types.OrganizationDetails) string {
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
	lessons := lessons_pkg.GetAllLessonMetadata(orgID, lessonIDs)

	m := ui.NewLessonListModel(orgID, orgName, lessons)
	p := tea.NewProgram(m)

	finalModel, err := p.Run()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if lessonModel, ok := finalModel.(ui.LessonListModel); ok {
		if lessonModel.WasQuit() {
			return ""
		}
		return lessonModel.SelectedLesson()
	}

	return ""
}

func init() {
	rootCmd.AddCommand(StartCmd)
	StartCmd.Flags().StringVarP(&lessonFlag, "lesson", "l", "", "lesson ID to practice")
	StartCmd.Flags().StringVarP(&themeFlag, "theme", "t", "default", "UI theme")
	StartCmd.Flags().StringVar(&colorFlag, "color", "", "custom color")
}
