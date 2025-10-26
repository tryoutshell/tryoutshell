package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tryoutshell/tryoutshell/internal/ui"
	"github.com/tryoutshell/tryoutshell/types"
)

var (
	lesson string
	theme  string
	color  string
)

var selectCmd = &cobra.Command{
	Use:   "start [org] [--lesson] [--theme] [--color]",
	Short: "select your lesson to start gaming",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalf("you need to specify the <org-id>")
		}
		if lesson == "" {
			log.Fatalf("--lesson is compulsory flag")
		}

		orgList := getOrganizationList()
		org := args[0]
		isOrgPresent := false
		var orgDetail types.OrganizationDetails

		// Check org is present in manifest.json
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

		// Check lesson is present under org
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

		fmt.Printf("🚀 Loading lesson: %s/%s\n\n", org, lesson)

		// Launch interactive UI with org and lesson IDs
		if err := ui.LaunchInteractive(org, lesson); err != nil {
			log.Fatalf("Error launching UI: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)
	selectCmd.Flags().StringVarP(&lesson, "lesson", "l", "", "lesson which you want to practice under orgs")
	selectCmd.Flags().StringVarP(&theme, "theme", "t", "default", "to override the default theme")
	selectCmd.Flags().StringVar(&color, "color", "", "to override the default color")
}
