package ui

import (
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	lessons_pkg "github.com/tryoutshell/tryoutshell/internal/lessons"
)

// LaunchPresentation reads a markdown file, parses it into slides, and starts
// the terminal-based slide presentation UI.
func LaunchPresentation(filePath string) error {
	raw, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %q: %w", filePath, err)
	}

	slides := ParseSlides(string(raw))

	m := NewSlideModel(slides)

	p := tea.NewProgram(
		m,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		return fmt.Errorf("error running presentation: %w", err)
	}

	return nil
}

// LaunchInteractive starts the Bubble Tea UI
func LaunchInteractive(orgID, lessonID string) error {
	// Load lesson data
	lesson, err := lessons_pkg.GetLessonContent(orgID, lessonID)
	if err != nil {
		return fmt.Errorf("failed to load lesson: %w", err)
	}

	// Create working directory for this lesson
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	workingDir := filepath.Join(homeDir, "tryoutshell-labs", orgID, lessonID)

	// Create directory if it doesn't exist
	if err := os.MkdirAll(workingDir, 0755); err != nil {
		return fmt.Errorf("failed to create working directory: %w", err)
	}

	// Change to working directory FIRST
	if err := os.Chdir(workingDir); err != nil {
		return fmt.Errorf("failed to change to working directory: %w", err)
	}

	fmt.Printf("📂 Working directory: %s\n", workingDir)
	fmt.Printf("📝 All files will be created here\n\n")

	// Create model AFTER changing directory (so runner picks up correct dir)
	m := NewModel(orgID, lessonID, lesson)

	// Verify working directory is set correctly in runner
	fmt.Printf("🔍 Runner working directory: %s\n\n", m.runner.GetWorkingDir())

	// Create program with proper mouse support
	p := tea.NewProgram(
		m,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	// Run the program
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("error running program: %w", err)
	}

	return nil
}
