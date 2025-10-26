package ui

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	Name       string
	Primary    lipgloss.Color
	Secondary  lipgloss.Color
	Success    lipgloss.Color
	Error      lipgloss.Color
	Warning    lipgloss.Color
	Muted      lipgloss.Color
	Background lipgloss.Color
	Border     lipgloss.Border
}

var Themes = map[string]Theme{
	"default": {
		Name:       "default",
		Primary:    lipgloss.Color("63"),  // Blue
		Secondary:  lipgloss.Color("240"), // Gray
		Success:    lipgloss.Color("42"),  // Green
		Error:      lipgloss.Color("196"), // Red
		Warning:    lipgloss.Color("226"), // Yellow
		Muted:      lipgloss.Color("240"), // Dark gray
		Background: lipgloss.Color(""),    // Terminal default
		Border:     lipgloss.RoundedBorder(),
	},
	"dark": {
		Name:       "dark",
		Primary:    lipgloss.Color("219"), // Pink
		Secondary:  lipgloss.Color("240"),
		Success:    lipgloss.Color("120"), // Light green
		Error:      lipgloss.Color("203"), // Light red
		Warning:    lipgloss.Color("221"),
		Muted:      lipgloss.Color("240"),
		Background: lipgloss.Color("0"), // Black
		Border:     lipgloss.ThickBorder(),
	},
	"cyberpunk": {
		Name:       "cyberpunk",
		Primary:    lipgloss.Color("201"), // Magenta
		Secondary:  lipgloss.Color("240"),
		Success:    lipgloss.Color("46"), // Neon green
		Error:      lipgloss.Color("196"),
		Warning:    lipgloss.Color("226"),
		Muted:      lipgloss.Color("240"),
		Background: lipgloss.Color("16"), // Dark gray
		Border:     lipgloss.DoubleBorder(),
	},
}

func GetTheme(name string) Theme {
	if theme, ok := Themes[name]; ok {
		return theme
	}
	return Themes["default"]
}
