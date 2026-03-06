package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/tryoutshell/tryoutshell/internal/lesson"
	"github.com/tryoutshell/tryoutshell/internal/progress"
)

var progressCmd = &cobra.Command{
	Use:   "progress",
	Short: "Show a summary of your lesson progress",
	Run: func(cmd *cobra.Command, args []string) {
		store := progress.NewStore()
		allLessons := lesson.DiscoverLessons()
		groups := lesson.GroupByOrg(allLessons)
		allProgress := store.GetAllProgress()

		headerStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63"))
		completedStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("42"))
		inProgressStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("226"))
		notStartedStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
		accentStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63"))

		fmt.Println()
		fmt.Println(headerStyle.Render("  📊 Lesson Progress"))
		fmt.Println()

		colOrg := 16
		colLesson := 28
		colStatus := 14
		colQuiz := 14
		colTime := 14

		header := fmt.Sprintf("  %-*s %-*s %-*s %-*s %-*s",
			colOrg, "Org",
			colLesson, "Lesson",
			colStatus, "Status",
			colQuiz, "Quiz Score",
			colTime, "Time Spent",
		)
		fmt.Println(accentStyle.Render(header))
		fmt.Println(accentStyle.Render("  " + strings.Repeat("─", colOrg+colLesson+colStatus+colQuiz+colTime+4)))

		totalLessons := 0
		completedCount := 0

		for orgID, lessons := range groups {
			for i, dl := range lessons {
				totalLessons++
				key := orgID + "/" + dl.LessonID
				p, exists := allProgress[key]

				orgLabel := ""
				if i == 0 {
					orgLabel = orgID
				}

				status := notStartedStyle.Render("not started")
				if exists && p.Completed {
					status = completedStyle.Render("✓ completed")
					completedCount++
				} else if exists && !p.LastAccess.IsZero() {
					status = inProgressStyle.Render("⟳ in progress")
				}

				quiz := notStartedStyle.Render("—")
				if exists && p.QuizTotal > 0 {
					quiz = fmt.Sprintf("%d/%d", p.QuizScore, p.QuizTotal)
				}

				timeStr := notStartedStyle.Render("—")
				if exists && p.TimeSpentMs > 0 {
					timeStr = formatDuration(time.Duration(p.TimeSpentMs) * time.Millisecond)
				}

				fmt.Printf("  %-*s %-*s %-*s %-*s %-*s\n",
					colOrg, orgLabel,
					colLesson, dl.LessonMeta.Title,
					colStatus+10, status,
					colQuiz, quiz,
					colTime, timeStr,
				)
			}
		}

		fmt.Println()
		pct := 0.0
		if totalLessons > 0 {
			pct = float64(completedCount) / float64(totalLessons) * 100
		}
		fmt.Printf("  Overall: %s of %d lessons completed (%.0f%%)\n\n",
			completedStyle.Render(fmt.Sprintf("%d", completedCount)),
			totalLessons, pct,
		)
	},
}

func formatDuration(d time.Duration) string {
	if d < time.Minute {
		return fmt.Sprintf("%ds", int(d.Seconds()))
	}
	if d < time.Hour {
		return fmt.Sprintf("%dm %ds", int(d.Minutes()), int(d.Seconds())%60)
	}
	return fmt.Sprintf("%dh %dm", int(d.Hours()), int(d.Minutes())%60)
}

func init() {
	rootCmd.AddCommand(progressCmd)
}
