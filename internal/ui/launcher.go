package ui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	lessons_pkg "github.com/tryoutshell/tryoutshell/internal/lessons"
)

// LaunchInteractive starts the Bubble Tea UI
func LaunchInteractive(orgID, lessonID string) error {
	// Load lesson data
	lesson, err := lessons_pkg.GetLessonContent(orgID, lessonID)
	if err != nil {
		return fmt.Errorf("failed to load lesson: %w", err)
	}

	// Create model
	m := NewModel(orgID, lessonID, lesson)

	// Create program with proper mouse support
	p := tea.NewProgram(
		m,
		tea.WithAltScreen(),       // Use alternate screen buffer
		tea.WithMouseCellMotion(), // Enable mouse motion events
	)

	// Run the program
	if _, err := p.Run(); err != nil {
		m.runner.Cleanup()
		return fmt.Errorf("error running program: %w", err)
	}
	m.runner.Cleanup()
	return nil
}
