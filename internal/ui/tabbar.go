package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type TabBar struct {
	appName    string
	lessonName string
	author     string
	version    string
	width      int
}

func NewTabBar(orgID, lessonTitle, author, version string, width int) TabBar {
	return TabBar{
		appName:    "TryOutShell",
		lessonName: fmt.Sprintf("%s - %s", orgID, lessonTitle),
		author:     author,
		version:    version,
		width:      width,
	}
}

func (t TabBar) Render() string {
	// Create flawz-style tab bar
	leftStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("183")). // Purple like flawz
		Foreground(lipgloss.Color("0")).   // Black text
		Bold(true).
		Padding(0, 2)

	rightStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("183")).
		Foreground(lipgloss.Color("0")).
		Padding(0, 2).
		Align(lipgloss.Right)

	// Left side: app name | lesson
	leftContent := fmt.Sprintf("%s | %s", t.appName, t.lessonName)

	// Right side: version with ❤ by author
	rightContent := fmt.Sprintf("%s with ❤ by %s", t.version, t.author)

	// Calculate spacing
	leftRendered := leftStyle.Render(leftContent)
	rightRendered := rightStyle.Render(rightContent)

	leftWidth := lipgloss.Width(leftRendered)
	rightWidth := lipgloss.Width(rightRendered)

	// Fill middle with background color
	middleWidth := t.width - leftWidth - rightWidth
	if middleWidth < 0 {
		middleWidth = 0
	}

	middleStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("183")).
		Width(middleWidth)

	middle := middleStyle.Render("")

	return lipgloss.JoinHorizontal(lipgloss.Top, leftRendered, middle, rightRendered)
}
