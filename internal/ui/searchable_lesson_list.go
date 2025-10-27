package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	lessons_pkg "github.com/tryoutshell/tryoutshell/internal/lessons"
)

type SearchableLessonListModel struct {
	list          list.Model
	searchInput   textinput.Model
	searchActive  bool
	originalItems []list.Item
	orgID         string
	orgName       string
	choice        string
	quitting      bool
}

func NewSearchableLessonList(orgID, orgName string, lessons []lessons_pkg.LessonMetadata) SearchableLessonListModel {
	items := make([]list.Item, len(lessons))
	for i, lesson := range lessons {
		items[i] = LessonItem{metadata: lesson}
	}

	l := list.New(items, lessonItemDelegate{}, 100, 20)
	l.Title = fmt.Sprintf("Select a Lesson from %s", orgName)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = lessonTitleStyle

	ti := textinput.New()
	ti.Placeholder = "Type to search lessons..."
	ti.CharLimit = 50
	ti.Width = 50

	return SearchableLessonListModel{
		list:          l,
		searchInput:   ti,
		originalItems: items,
		orgID:         orgID,
		orgName:       orgName,
	}
}

func (m SearchableLessonListModel) Init() tea.Cmd {
	return nil
}

func (m SearchableLessonListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		if m.searchActive {
			switch msg.String() {
			case "enter":
				if len(m.list.Items()) > 0 {
					i, ok := m.list.SelectedItem().(LessonItem)
					if ok {
						m.choice = i.metadata.ID
					}
					return m, tea.Quit
				}

			case "esc":
				m.searchActive = false
				m.searchInput.Reset()
				m.list.SetItems(m.originalItems)
				return m, nil

			case "ctrl+c":
				m.quitting = true
				return m, tea.Quit

			default:
				m.searchInput, cmd = m.searchInput.Update(msg)
				m.filterList()
				return m, cmd
			}
		} else {
			switch msg.String() {
			case ":", "/":
				m.searchActive = true
				m.searchInput.Focus()
				return m, textinput.Blink

			case "enter":
				i, ok := m.list.SelectedItem().(LessonItem)
				if ok {
					m.choice = i.metadata.ID
				}
				return m, tea.Quit

			case "ctrl+c":
				m.quitting = true
				return m, tea.Quit

			case "q", "esc":
				m.choice = ""
				return m, tea.Quit
			}
		}
	}

	if !m.searchActive {
		m.list, cmd = m.list.Update(msg)
	}

	return m, cmd
}

func (m *SearchableLessonListModel) filterList() {
	query := strings.ToLower(m.searchInput.Value())

	if query == "" {
		m.list.SetItems(m.originalItems)
		return
	}

	var filtered []list.Item
	for _, item := range m.originalItems {
		lesson := item.(LessonItem)
		meta := lesson.metadata

		// Search in title, description, tags, ID
		searchText := strings.ToLower(fmt.Sprintf("%s %s %s %s",
			meta.ID, meta.Title, meta.Description, strings.Join(meta.Tags, " ")))

		if strings.Contains(searchText, query) {
			filtered = append(filtered, item)
		}
	}

	m.list.SetItems(filtered)
	if len(filtered) > 0 {
		m.list.Select(0)
	}
}

func (m SearchableLessonListModel) View() string {
	if m.quitting {
		return ""
	}

	var s strings.Builder

	// Search bar if active
	if m.searchActive {
		searchBox := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("39")).
			Padding(0, 1).
			Render("🔍 " + m.searchInput.View())

		s.WriteString(searchBox + "\n\n")

		resultsInfo := lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Render(fmt.Sprintf("Found %d lessons", len(m.list.Items())))
		s.WriteString(resultsInfo + "\n\n")
	}

	// List
	s.WriteString(m.list.View() + "\n")

	// Help text
	if m.searchActive {
		help := lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Render("  Type to search  •  Enter: Select  •  Esc: Cancel  •  Ctrl+C: Quit")
		s.WriteString(help)
	} else {
		help := lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Render("  ↑/↓: Navigate  •  : or /: Search  •  Enter: Select  •  Esc: Back  •  q: Quit")
		s.WriteString(help)
	}

	return s.String()
}

func (m SearchableLessonListModel) SelectedLesson() string {
	return m.choice
}

func (m SearchableLessonListModel) WasQuit() bool {
	return m.quitting
}
