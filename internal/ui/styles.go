package ui

import (
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	Theme Theme

	// Headers
	AppTitle      lipgloss.Style
	SectionHeader lipgloss.Style
	StepTitle     lipgloss.Style

	// Content
	Paragraph  lipgloss.Style
	CodeBlock  lipgloss.Style
	InlineCode lipgloss.Style
	Bold       lipgloss.Style
	Muted      lipgloss.Style

	// UI Elements
	Border         lipgloss.Style
	SuccessMsg     lipgloss.Style
	ErrorMsg       lipgloss.Style
	HintBox        lipgloss.Style
	CalloutTip     lipgloss.Style
	CalloutWarning lipgloss.Style
	SuccessStyle   func() lipgloss.Style
	// Command execution
	CommandPrompt lipgloss.Style
	CommandInput  lipgloss.Style
	OutputBox     lipgloss.Style
	OutputSuccess lipgloss.Style
	OutputError   lipgloss.Style

	// Navigation
	HelpText lipgloss.Style
}

func NewStyles(theme Theme) *Styles {
	s := &Styles{Theme: theme}

	// Headers
	s.AppTitle = lipgloss.NewStyle().
		Bold(true).
		Foreground(theme.Primary).
		Align(lipgloss.Center).
		Padding(0, 2)

	s.SectionHeader = lipgloss.NewStyle().
		Bold(true).
		Foreground(theme.Primary).
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderForeground(theme.Primary).
		Padding(0, 1)

	s.StepTitle = lipgloss.NewStyle().
		Bold(true).
		Foreground(theme.Primary).
		Padding(1, 2).
		Border(theme.Border).
		BorderForeground(theme.Primary)

	// Content
	s.Paragraph = lipgloss.NewStyle().
		Padding(0, 2)

	s.CodeBlock = lipgloss.NewStyle().
		Foreground(lipgloss.Color("214")).
		Background(lipgloss.Color("235")).
		Padding(1, 2).
		Border(theme.Border).
		BorderForeground(theme.Secondary)

	s.InlineCode = lipgloss.NewStyle().
		Foreground(lipgloss.Color("214")).
		Background(lipgloss.Color("235")).
		Padding(0, 1)

	s.Bold = lipgloss.NewStyle().Bold(true)

	s.Muted = lipgloss.NewStyle().
		Foreground(theme.Muted)

	// UI Elements
	s.Border = lipgloss.NewStyle().
		Border(theme.Border).
		BorderForeground(theme.Primary).
		Padding(1, 2)

	s.SuccessMsg = lipgloss.NewStyle().
		Foreground(theme.Success).
		Bold(true)

	s.ErrorMsg = lipgloss.NewStyle().
		Foreground(theme.Error).
		Bold(true)

	s.HintBox = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(theme.Warning).
		Foreground(theme.Warning).
		Padding(1, 2)

	s.CalloutTip = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(theme.Success).
		Padding(1, 2)

	s.CalloutWarning = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(theme.Warning).
		Padding(1, 2)

	// Command execution
	s.CommandPrompt = lipgloss.NewStyle().
		Bold(true).
		Foreground(theme.Success)

	s.CommandInput = lipgloss.NewStyle().
		Border(theme.Border).
		BorderForeground(theme.Primary).
		Padding(1, 2)

	s.OutputBox = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2)

	s.OutputSuccess = s.OutputBox.Copy().
		BorderForeground(theme.Success)

	s.OutputError = s.OutputBox.Copy().
		BorderForeground(theme.Error)

	// Navigation
	s.HelpText = lipgloss.NewStyle().
		Foreground(theme.Muted).
		Padding(0, 2)

	return s
}

// Add helper method to Theme
func (t Theme) SuccessStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(t.Success)
}
