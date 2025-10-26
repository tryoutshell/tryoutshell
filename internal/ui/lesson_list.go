package ui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	lessons_pkg "github.com/tryoutshell/tryoutshell/internal/lessons"
)

var (
	lessonTitleStyle        = lipgloss.NewStyle().MarginLeft(2).Bold(true)
	lessonItemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	lessonSelectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Foreground(lipgloss.Color("170")).
				BorderLeft(true).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("170"))
)

type LessonItem struct {
	metadata lessons_pkg.LessonMetadata
}

func (i LessonItem) FilterValue() string { return i.metadata.Title }

type lessonItemDelegate struct{}

func (d lessonItemDelegate) Height() int                             { return 4 }
func (d lessonItemDelegate) Spacing() int                            { return 1 }
func (d lessonItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d lessonItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(LessonItem)
	if !ok {
		return
	}

	meta := i.metadata

	// Format difficulty stars
	difficultyStars := "⭐"
	switch meta.Difficulty {
	case "intermediate":
		difficultyStars = "⭐⭐"
	case "advanced":
		difficultyStars = "⭐⭐⭐"
	}

	// Format tags
	tagStr := ""
	if len(meta.Tags) > 0 {
		tagStr = strings.Join(meta.Tags, ", ")
		if len(tagStr) > 40 {
			tagStr = tagStr[:40] + "..."
		}
	}

	// Main title line
	title := fmt.Sprintf("%s", meta.Title)

	// Metadata line
	metaLine := fmt.Sprintf("%s %s  •  ⏱ %s  •  🏷 %s",
		difficultyStars, meta.Difficulty, meta.Duration, tagStr)

	// Description
	desc := meta.Description
	if len(desc) > 60 {
		desc = desc[:60] + "..."
	}

	// Prerequisites
	prereqLine := ""
	if len(meta.Prerequisites) > 0 {
		prereqLine = "\n    📋 Prerequisites: " + strings.Join(meta.Prerequisites, ", ")
	}

	fn := lessonItemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			content := strings.Join(s, "\n")
			return lessonSelectedItemStyle.Render(content)
		}
		title = "│ " + title
	} else {
		title = "  " + title
	}

	content := fmt.Sprintf("%s\n    %s\n    %s%s",
		title, metaLine, desc, prereqLine)

	fmt.Fprint(w, fn(content))
}

type LessonListModel struct {
	list     list.Model
	orgID    string
	choice   string
	quitting bool
}

func NewLessonListModel(orgID, orgName string, lessons []lessons_pkg.LessonMetadata) LessonListModel {
	items := make([]list.Item, len(lessons))
	for i, lesson := range lessons {
		items[i] = LessonItem{metadata: lesson}
	}
	totalLessons := len(lessons)
	const defaultWidth = 100
	const listHeight = 20

	l := list.New(items, lessonItemDelegate{}, defaultWidth, listHeight)
	l.Title = fmt.Sprintf("Select a Lesson from %s | Total Lessons: %d", orgName, totalLessons)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = lessonTitleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return LessonListModel{
		list:  l,
		orgID: orgID,
	}
}

func (m LessonListModel) Init() tea.Cmd {
	return nil
}

func (m LessonListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "q", "esc":
			// Return empty choice to go back
			m.choice = ""
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(LessonItem)
			if ok {
				m.choice = i.metadata.ID
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m LessonListModel) View() string {
	if m.quitting {
		return ""
	}
	return "\n" + m.list.View()
}

func (m LessonListModel) SelectedLesson() string {
	return m.choice
}

func (m LessonListModel) WasQuit() bool {
	return m.quitting
}
