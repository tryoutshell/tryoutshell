package ui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

type OrgItem struct {
	ID          string // Capitalized
	Name        string
	Description string
	Logo        string
}

func (i OrgItem) FilterValue() string { return i.Name }

type orgItemDelegate struct{}

func (d orgItemDelegate) Height() int                             { return 1 }
func (d orgItemDelegate) Spacing() int                            { return 0 }
func (d orgItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d orgItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(OrgItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%s  %s", i.Logo, i.Name)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("▸ " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type OrgListModel struct {
	list     list.Model
	choice   string
	quitting bool
}

func NewOrgListModel(orgs []OrgItem) OrgListModel {
	items := make([]list.Item, len(orgs))
	for i, org := range orgs {
		items[i] = org
	}

	const defaultWidth = 80
	const listHeight = 14

	l := list.New(items, orgItemDelegate{}, defaultWidth, listHeight)
	l.Title = "Select an Organization"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return OrgListModel{list: l}
}

func (m OrgListModel) Init() tea.Cmd {
	return nil
}

func (m OrgListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(OrgItem)
			if ok {
				m.choice = i.ID
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m OrgListModel) View() string {
	if m.quitting {
		return ""
	}
	if m.state == ListStateOrganizations {
		return m.renderOrganizations()
	}
	return "\n" + m.list.View()
}

func (m OrgListModel) SelectedOrg() string {
	return m.choice
}
