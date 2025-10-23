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

	// UI Elements
	Border         lipgloss.Style
	ProgressBar    lipgloss.Style
	SuccessMsg     lipgloss.Style
	ErrorMsg       lipgloss.Style
	HintBox        lipgloss.Style
	CalloutTip     lipgloss.Style
	CalloutWarning lipgloss.Style

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
		MarginBottom(1).
		Align(lipgloss.Center)

	s.SectionHeader = lipgloss.NewStyle().
		Bold(true).
		Foreground(theme.Primary).
		BorderStyle(lipgloss.ThickBorder()).
		BorderBottom(true).
		BorderForeground(theme.Primary).
		Padding(0, 1).
		MarginBottom(1)

	s.StepTitle = lipgloss.NewStyle().
		Bold(true).
		Foreground(theme.Primary).
		Padding(0, 1).
		Border(theme.Border).
		BorderForeground(theme.Primary)

	// Content
	s.Paragraph = lipgloss.NewStyle().
		Foreground(lipgloss.Color("")).
		MarginBottom(1)

	s.CodeBlock = lipgloss.NewStyle().
		Foreground(lipgloss.Color("214")). // Orange
		Background(lipgloss.Color("235")). // Dark gray
		Padding(1, 2).
		Border(theme.Border).
		BorderForeground(theme.Secondary).
		MarginTop(1).
		MarginBottom(1)

	s.InlineCode = lipgloss.NewStyle().
		Foreground(lipgloss.Color("214")).
		Background(lipgloss.Color("235")).
		Padding(0, 1)

	s.Bold = lipgloss.NewStyle().Bold(true)

	// UI Elements
	s.Border = lipgloss.NewStyle().
		Border(theme.Border).
		BorderForeground(theme.Primary).
		Padding(1, 2)

	s.SuccessMsg = lipgloss.NewStyle().
		Foreground(theme.Success).
		Bold(true).
		MarginTop(1)

	s.ErrorMsg = lipgloss.NewStyle().
		Foreground(theme.Error).
		Bold(true).
		MarginTop(1)

	s.HintBox = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(theme.Warning).
		Foreground(theme.Warning).
		Padding(0, 1).
		MarginTop(1)

	s.CalloutTip = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(theme.Success).
		Padding(0, 1).
		MarginTop(1)

	s.CalloutWarning = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(theme.Warning).
		Padding(0, 1).
		MarginTop(1)

	// Command execution
	s.CommandPrompt = lipgloss.NewStyle().
		Bold(true).
		Foreground(theme.Success).
		SetString("$ ")

	s.CommandInput = lipgloss.NewStyle().
		Border(theme.Border).
		BorderForeground(theme.Primary).
		Padding(0, 1).
		MarginTop(1)

	s.OutputBox = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		MarginTop(1)

	s.OutputSuccess = s.OutputBox.Copy().
		BorderForeground(theme.Success)

	s.OutputError = s.OutputBox.Copy().
		BorderForeground(theme.Error)

	// Navigation
	s.HelpText = lipgloss.NewStyle().
		Foreground(theme.Muted).
		MarginTop(2)

	return s
}
