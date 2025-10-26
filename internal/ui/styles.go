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
	SubHeading    lipgloss.Style

	// Content
	Paragraph  lipgloss.Style
	CodeBlock  lipgloss.Style
	InlineCode lipgloss.Style
	Bold       lipgloss.Style
	Muted      lipgloss.Style
	Link       lipgloss.Style

	// UI Elements
	Border         lipgloss.Style
	BoxBorder      lipgloss.Style
	SuccessMsg     lipgloss.Style
	ErrorMsg       lipgloss.Style
	WarningMsg     lipgloss.Style
	InfoMsg        lipgloss.Style
	HintBox        lipgloss.Style
	CalloutTip     lipgloss.Style
	CalloutWarning lipgloss.Style

	// Command execution
	CommandPrompt  lipgloss.Style
	CommandInput   lipgloss.Style
	CommandExample lipgloss.Style
	OutputBox      lipgloss.Style
	OutputSuccess  lipgloss.Style
	OutputError    lipgloss.Style

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
		Padding(0, 2).
		MarginBottom(1)

	s.StepTitle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("15")). // Bright white
		Background(theme.Primary).
		Padding(0, 2).
		MarginBottom(1)

	s.SubHeading = lipgloss.NewStyle().
		Bold(true).
		Foreground(theme.Primary).
		MarginTop(1).
		MarginBottom(1)

	// Content
	s.Paragraph = lipgloss.NewStyle().
		Padding(0, 2).
		MarginBottom(1)

	s.CodeBlock = lipgloss.NewStyle().
		Foreground(lipgloss.Color("252")).
		Background(lipgloss.Color("235")).
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		MarginTop(1).
		MarginBottom(1)

	s.InlineCode = lipgloss.NewStyle().
		Foreground(lipgloss.Color("214")). // Orange
		Background(lipgloss.Color("235")).
		Padding(0, 1)

	s.Bold = lipgloss.NewStyle().Bold(true)

	s.Muted = lipgloss.NewStyle().
		Foreground(lipgloss.Color("240"))

	s.Link = lipgloss.NewStyle().
		Foreground(lipgloss.Color("39")). // Cyan
		Underline(true)

	// UI Elements
	s.Border = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(theme.Primary).
		Padding(1, 2)

	s.BoxBorder = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		Padding(1, 2)

	s.SuccessMsg = lipgloss.NewStyle().
		Foreground(theme.Success).
		Bold(true).
		MarginTop(1)

	s.ErrorMsg = lipgloss.NewStyle().
		Foreground(theme.Error).
		Bold(true).
		MarginTop(1)

	s.WarningMsg = lipgloss.NewStyle().
		Foreground(theme.Warning).
		Bold(true)

	s.InfoMsg = lipgloss.NewStyle().
		Foreground(theme.Primary).
		Bold(true)

	s.HintBox = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(theme.Warning).
		Foreground(lipgloss.Color("252")).
		Padding(1, 2).
		MarginTop(1)

	s.CalloutTip = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(theme.Success).
		Padding(1, 2).
		MarginTop(1)

	s.CalloutWarning = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(theme.Warning).
		Padding(1, 2).
		MarginTop(1)

	// Command execution
	s.CommandPrompt = lipgloss.NewStyle().
		Bold(true).
		Foreground(theme.Success).
		SetString("$ ")

	s.CommandInput = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(theme.Primary).
		Padding(1, 2).
		MarginTop(1).
		MarginBottom(1)

	s.CommandExample = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		Padding(1, 2).
		MarginTop(1).
		MarginBottom(1)

	s.OutputBox = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		Padding(1, 2).
		MarginTop(1)

	s.OutputSuccess = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(theme.Success).
		Padding(1, 2).
		MarginTop(1)

	s.OutputError = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(theme.Error).
		Padding(1, 2).
		MarginTop(1)

	// Navigation
	s.HelpText = lipgloss.NewStyle().
		Foreground(theme.Muted).
		Padding(0, 2)

	return s
}
