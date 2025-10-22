package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	lessons_pkg "github.com/tryoutshell/tryoutshell/internal/lessons"
	"log"
	"strings"
)

var (
	lesson string
	theme  string
	color  string
)

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:   "start [org] [--lesson] [--theme] [--color] ",
	Short: "select your lesson to start gaming",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalf("you need to specify the <id>")
		}
		if lesson == "" {
			log.Fatalf("--lesson is complusory flag")
		}
		orgList := getOrganizationList()
		org := args[0]
		isOrgPresent := false
		var orgDetail OrganizationDetails
		// check org is present in manifest.json (collection of all existing org/lesson)
		for _, d := range orgList {
			if strings.Contains(string(org), d.Id) {
				isOrgPresent = true
				orgDetail = d
				break
			}
		}
		if !isOrgPresent {
			log.Fatalf("%s organization id not present", org)
		}
		// check lesson is present or not under org
		lessons := orgDetail.Lessons
		isLessonPresent := false
		for _, d := range lessons {
			if strings.EqualFold(lesson, d) {
				isLessonPresent = true
				break
			}
		}
		if !isLessonPresent {
			log.Fatalf("%s lesson is not present under %s", lesson, org)
		}
		log.Printf("loading lesson %s 🚀 \n\n", strings.ToLower(lesson))
		// logic for loading the lesson
		lessonFileContent := lessons_pkg.GetLessonContent()
		lessonIntroduction := lessonFileContent.Introduction
		lessonMetadata := lessonFileContent.Metadata
		lessonSteps := lessonFileContent.Steps
		lessonConclusion := lessonFileContent.Conclusion

		fmt.Println("---------------------------------------------------")
		fmt.Println(lessonIntroduction)
		fmt.Println("---------------------------------------------------")
		fmt.Println(lessonMetadata)
		fmt.Println("---------------------------------------------------")
		fmt.Println(lessonSteps)
		fmt.Println("---------------------------------------------------")
		fmt.Println(lessonConclusion)

	},
}

func init() {
	rootCmd.AddCommand(selectCmd)
	selectCmd.AddCommand()
	selectCmd.Flags().StringVarP(&lesson, "lesson", "l", "", "lesson which you want to practice under orgs")
	//selectCmd.Flags().StringVarP(&theme, "theme", "t", "", "to override the default theme")
	//selectCmd.Flags().StringVar(&color, "color", "", "to override the default color")
}

//func checkOrgIdPresent(orgList []OrganizationDetails, org_id string) (OrganizationDetails, bool) {
//	isOrgPresent := false
//	var orgDetail OrganizationDetails
//	for _, d := range orgList {
//		if strings.Contains(string(lesson), d.Id) {
//			isOrgPresent = true
//			orgDetail = d
//			return orgDetail, isOrgPresent
//		}
//	}
//	return
//}

func loadLesson(lesson string) {

}
