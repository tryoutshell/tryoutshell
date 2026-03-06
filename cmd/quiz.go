package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tryoutshell/tryoutshell/internal/lesson"
	"github.com/tryoutshell/tryoutshell/internal/quiz"
)

var quizCmd = &cobra.Command{
	Use:   "quiz <org> <lesson>",
	Short: "Launch quiz mode for a lesson",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		orgID := args[0]
		lessonID := args[1]

		allLessons := lesson.DiscoverLessons()
		dl := lesson.FindLesson(allLessons, orgID, lessonID)
		if dl == nil {
			log.Fatalf("Lesson '%s/%s' not found", orgID, lessonID)
		}

		if len(dl.LessonMeta.Quiz) == 0 {
			fmt.Println("This lesson has no quiz questions.")
			return
		}

		if err := quiz.Launch(orgID, lessonID, dl.LessonMeta.Quiz); err != nil {
			log.Fatalf("Error running quiz: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(quizCmd)
}
