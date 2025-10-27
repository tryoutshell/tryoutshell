package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// SearchableListModel wraps list with custom search
type SearchableListModel struct {
	list          list.Model
	searchInput   textinput.Model
	searchActive  bool
	originalItems []list.Item
	choice        string
	quitting      bool
}

func NewSearchableOrgList(orgs []OrgItem) SearchableListModel {
	items := make([]list.Item, len(orgs))
	for i, org := range orgs {
		items[i] = org
	}

	l := list.New(items, orgItemDelegate{}, 80, 14)
	l.Title = "Select an Organization"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false) // We handle search ourselves
	l.Styles.Title = titleStyle

	// Create search input
	ti := textinput.New()
	ti.Placeholder = "Type to search..."
	ti.CharLimit = 50
	ti.Width = 50

	return SearchableListModel{
		list:          l,
		searchInput:   ti,
		originalItems: items,
	}
}

func (m SearchableListModel) Init() tea.Cmd {
	return nil
}

func (m SearchableListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		// If search is active, handle search input
		if m.searchActive {
			switch msg.String() {
			case "enter":
				// If search has results, select first
				if len(m.list.Items()) > 0 {
					i, ok := m.list.SelectedItem().(OrgItem)
					if ok {
						m.choice = i.ID
					}
					return m, tea.Quit
				}

			case "esc":
				// Cancel search, restore all items
				m.searchActive = false
				m.searchInput.Reset()
				m.list.SetItems(m.originalItems)
				return m, nil

			case "ctrl+c", "q":
				m.quitting = true
				return m, tea.Quit

			default:
				// Update search input
				m.searchInput, cmd = m.searchInput.Update(msg)

				// Filter list based on search
				m.filterList()
				return m, cmd
			}
		} else {
			// Normal navigation mode
			switch msg.String() {
			case ":":
				// Activate search
				m.searchActive = true
				m.searchInput.Focus()
				return m, textinput.Blink

			case "/":
				// Also activate search (alternative)
				m.searchActive = true
				m.searchInput.Focus()
				return m, textinput.Blink

			case "enter":
				i, ok := m.list.SelectedItem().(OrgItem)
				if ok {
					m.choice = i.ID
				}
				return m, tea.Quit

			case "ctrl+c", "q":
				m.quitting = true
				return m, tea.Quit
			}
		}
	}

	// Update list (for arrow keys, etc.)
	if !m.searchActive {
		m.list, cmd = m.list.Update(msg)
	}

	return m, cmd
}

func (m *SearchableListModel) filterList() {
	query := strings.ToLower(m.searchInput.Value())

	if query == "" {
		// No search query, show all
		m.list.SetItems(m.originalItems)
		return
	}

	// Filter items
	var filtered []list.Item
	for _, item := range m.originalItems {
		org := item.(OrgItem)

		// Match against ID, Name, or Description
		if strings.Contains(strings.ToLower(org.ID), query) ||
			strings.Contains(strings.ToLower(org.Name), query) ||
			strings.Contains(strings.ToLower(org.Description), query) {
			filtered = append(filtered, item)
		}
	}

	m.list.SetItems(filtered)

	// Reset selection to first item
	if len(filtered) > 0 {
		m.list.Select(0)
	}
}

func (m SearchableListModel) View() string {
	if m.quitting {
		return ""
	}

	var s strings.Builder

	// Banner
	banner := getBanner()
	bannerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("39")).
		Bold(true).
		Align(lipgloss.Center)
	s.WriteString(bannerStyle.Render(banner) + "\n\n")

	// Subtitle
	subtitle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Align(lipgloss.Center).
		Render("🚀 Interactive DevOps Learning in Your Terminal")
	s.WriteString(subtitle + "\n\n")

	// Search bar (if active)
	if m.searchActive {
		searchBox := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("39")).
			Padding(0, 1).
			Render("🔍 " + m.searchInput.View())

		s.WriteString(searchBox + "\n\n")

		resultsInfo := lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Render(fmt.Sprintf("Found %d results", len(m.list.Items())))
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
			Render("  ↑/↓: Navigate  •  : or /: Search  •  Enter: Select  •  q: Quit")
		s.WriteString(help)
	}

	return s.String()
}

func (m SearchableListModel) SelectedOrg() string {
	return m.choice
}
