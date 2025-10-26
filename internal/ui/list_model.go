package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tryoutshell/tryoutshell/types"
)

type ListState int

const (
	ListStateOrganizations ListState = iota
	ListStateLessons
)

// type OrganizationDetails struct {
// 	Id          string   `json:"id"`
// 	Name        string   `json:"name"`
// 	Description string   `json:"description"`
// 	Logo        string   `json:"logo"`
// 	Lessons     []string `json:"lessons"`
// }

type ListModel struct {
	state             ListState
	organizations     []types.OrganizationDetails
	selectedOrg       int
	selectedLesson    int
	currentOrg        types.OrganizationDetails
	styles            *Styles
	width             int
	height            int
	quitting          bool
	shouldStartLesson bool
}

func NewListModel(orgs []types.OrganizationDetails) ListModel {
	theme := GetTheme("default")
	styles := NewStyles(theme)

	return ListModel{
		state:          ListStateOrganizations,
		organizations:  orgs,
		selectedOrg:    0,
		selectedLesson: 0,
		styles:         styles,
	}
}
func (m ListModel) ShouldStartLesson() bool {
	return m.shouldStartLesson
}

func (m ListModel) GetSelectedLesson() (string, string) {
	if len(m.currentOrg.Lessons) > m.selectedLesson {
		return m.currentOrg.Id, m.currentOrg.Lessons[m.selectedLesson]
	}
	return "", ""
}

func (m ListModel) Init() tea.Cmd {
	return nil
}

func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "esc":
			if m.state == ListStateLessons {
				// Go back to org selection
				m.state = ListStateOrganizations
				m.selectedLesson = 0
			} else {
				m.quitting = true
				return m, tea.Quit
			}

		case "up", "k":
			if m.state == ListStateOrganizations {
				if m.selectedOrg > 0 {
					m.selectedOrg--
				}
			} else {
				if m.selectedLesson > 0 {
					m.selectedLesson--
				}
			}

		case "down", "j":
			if m.state == ListStateOrganizations {
				if m.selectedOrg < len(m.organizations)-1 {
					m.selectedOrg++
				}
			} else {
				if m.selectedLesson < len(m.currentOrg.Lessons)-1 {
					m.selectedLesson++
				}
			}

		case "enter":
			if m.state == ListStateOrganizations {
				// Select org, show lessons
				m.currentOrg = m.organizations[m.selectedOrg]
				m.state = ListStateLessons
				m.selectedLesson = 0
			} else {
				// Start lesson
				m.shouldStartLesson = true
				m.quitting = true
				return m, tea.Quit
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}

func (m ListModel) View() string {
	if m.quitting {
		return ""
	}

	if m.state == ListStateOrganizations {
		return m.renderOrganizations()
	}
	return m.renderLessons()
}

func (m ListModel) renderOrganizations() string {
	var s string

	// Banner
	banner := m.styles.AppTitle.Render(m.getBanner())
	s += banner + "\n\n"

	// Subtitle
	subtitle := m.styles.Muted.Render("🚀 Interactive Learning in Your Terminal")
	s += lipgloss.NewStyle().Align(lipgloss.Center).Width(m.width).Render(subtitle) + "\n\n"

	// Section header
	header := m.styles.SubHeading.Render("Select an organization:")
	s += "  " + header + "\n\n"

	// Organizations list
	for i, org := range m.organizations {
		var orgLine string

		if i == m.selectedOrg {
			// Selected org - highlighted
			orgStyle := lipgloss.NewStyle().
				Foreground(m.styles.Theme.Primary).
				Bold(true).
				Background(lipgloss.Color("237"))

			orgLine = orgStyle.Render(fmt.Sprintf("  ▸ %s  %-20s  %s  ",
				org.Logo, org.Name, org.Description))
		} else {
			// Unselected org
			orgLine = fmt.Sprintf("    %s  %-20s  %s",
				org.Logo, org.Name, org.Description)
		}

		s += orgLine + "\n"
	}

	s += "\n"

	// Help text
	help := m.styles.HelpText.Render("  ↑/↓ or j/k: Navigate  •  Enter: Select  •  q: Quit")
	s += help + "\n"

	return s
}

func (m ListModel) renderLessons() string {
	var s string

	// Header with org info
	headerContent := fmt.Sprintf("%s  %s Lessons", m.currentOrg.Logo, m.currentOrg.Name)
	header := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("15")).
		Background(m.styles.Theme.Primary).
		Padding(0, 2).
		Width(m.width - 4).
		Render(headerContent)
	s += header + "\n\n"

	// Section header
	sectionHeader := m.styles.SubHeading.Render("Available Lessons:")
	s += "  " + sectionHeader + "\n\n"

	// Lessons list
	for i, lessonID := range m.currentOrg.Lessons {
		// TODO: Load lesson metadata to show title, duration, difficulty
		// For now, just show lesson ID

		var lessonLine string

		if i == m.selectedLesson {
			// Selected lesson - highlighted with details
			lessonStyle := lipgloss.NewStyle().
				Foreground(m.styles.Theme.Primary).
				Bold(true).
				Background(lipgloss.Color("237")).
				Width(m.width - 8)

			lessonContent := fmt.Sprintf("  ▸ %s  ", lessonID)
			lessonLine = lessonStyle.Render(lessonContent)

			// Add metadata (placeholder - you'd load this from lesson file)
			metaStyle := lipgloss.NewStyle().
				Foreground(m.styles.Theme.Muted).
				MarginLeft(4)
			meta := metaStyle.Render("    ⭐ Beginner  •  ⏱ 20 min  •  📚 Learn by doing")

			s += lessonLine + "\n" + meta + "\n\n"
		} else {
			// Unselected lesson
			lessonLine = fmt.Sprintf("    %s", lessonID)
			s += lessonLine + "\n\n"
		}

		// Separator between lessons
		if i < len(m.currentOrg.Lessons)-1 {
			separator := m.styles.Muted.Render("  " + strings.Repeat("─", m.width-6))
			s += separator + "\n\n"
		}
	}

	s += "\n"

	// Help text
	help := m.styles.HelpText.Render("  ↑/↓ or j/k: Navigate  •  Enter: Start Lesson  •  Esc: Back  •  q: Quit")
	s += help + "\n"

	return s
}

func (m ListModel) getBanner() string {
	return `
████████╗██████╗ ██╗   ██╗ ██████╗ ██╗   ██╗████████╗███████╗██╗  ██╗
╚══██╔══╝██╔══██╗╚██╗ ██╔╝██╔═══██╗██║   ██║╚══██╔══╝██╔════╝██║  ██║
   ██║   ██████╔╝ ╚████╔╝ ██║   ██║██║   ██║   ██║   ███████╗███████║
   ██║   ██╔══██╗  ╚██╔╝  ██║   ██║██║   ██║   ██║   ╚════██║██╔══██║
   ██║   ██║  ██║   ██║   ╚██████╔╝╚██████╔╝   ██║   ███████║██║  ██║
   ╚═╝   ╚═╝  ╚═╝   ╚═╝    ╚═════╝  ╚═════╝    ╚═╝   ╚══════╝╚═╝  ╚═╝`
}

// Command to start a lesson
func startLesson(orgID, lessonID string) tea.Cmd {
	return func() tea.Msg {
		return StartLessonMsg{
			OrgID:    orgID,
			LessonID: lessonID,
		}
	}
}

type StartLessonMsg struct {
	OrgID    string
	LessonID string
}
