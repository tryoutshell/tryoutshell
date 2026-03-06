package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tryoutshell/tryoutshell/internal/lesson"
	"github.com/tryoutshell/tryoutshell/internal/ui"
	"github.com/tryoutshell/tryoutshell/types"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Interactive list of organizations and lessons",
	Run: func(cmd *cobra.Command, args []string) {
		allLessons := lesson.DiscoverLessons()
		orgList := buildOrgList(allLessons)

		selectedOrg := selectOrganization(orgList)
		if selectedOrg == "" {
			return
		}

		var orgDetail types.OrganizationDetails
		for _, org := range orgList {
			if org.Id == selectedOrg {
				orgDetail = org
				break
			}
		}

		selectedLesson := selectLesson(selectedOrg, orgDetail.Name, orgDetail.Lessons)
		if selectedLesson == "" {
			return
		}

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
			log.Fatalf("Error launching lesson: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func buildOrgList(allLessons []types.DiscoveredLesson) []types.OrganizationDetails {
	groups := lesson.GroupByOrg(allLessons)
	var orgList []types.OrganizationDetails

	for orgID, lessons := range groups {
		var lessonIDs []string
		for _, l := range lessons {
			lessonIDs = append(lessonIDs, l.LessonID)
		}

		meta := lessons[0].OrgMeta
		orgList = append(orgList, types.OrganizationDetails{
			Id:          orgID,
			Name:        meta.Name,
			Description: meta.Description,
			Logo:        meta.Logo,
			Lessons:     lessonIDs,
		})
	}

	return orgList
}
