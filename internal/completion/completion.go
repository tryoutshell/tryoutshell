package completion

import (
	"github.com/spf13/cobra"
	"github.com/tryoutshell/tryoutshell/internal/lesson"
)

func OrgCompletionFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) > 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	allLessons := lesson.DiscoverLessons()
	orgs := lesson.GetOrgList(allLessons)

	var suggestions []string
	for _, org := range orgs {
		suggestions = append(suggestions, org.ID)
	}
	return suggestions, cobra.ShellCompDirectiveNoFileComp
}

func LessonCompletionFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) != 1 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	orgID := args[0]
	allLessons := lesson.DiscoverLessons()
	groups := lesson.GroupByOrg(allLessons)

	lessons, ok := groups[orgID]
	if !ok {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	var suggestions []string
	for _, l := range lessons {
		suggestions = append(suggestions, l.LessonID)
	}
	return suggestions, cobra.ShellCompDirectiveNoFileComp
}
