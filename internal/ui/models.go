package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	lessons_pkg "github.com/tryoutshell/tryoutshell/internal/lessons"
	"github.com/tryoutshell/tryoutshell/internal/runner"
	"time"
)

// AppState represents different screens
type AppState int

const (
	StateIntroduction AppState = iota
	StateLesson
	StateConclusion
	StateQuitting
)

// StepState represents the state of current step
type StepState int

const (
	StepPending StepState = iota
	StepExecuting
	StepSuccess
	StepFailed
)

// Model is the main Bubble Tea model
type Model struct {
	// State
	state       AppState
	currentStep int
	stepState   StepState
	totalSteps  int

	// Lesson data
	lesson   lessons_pkg.LessonFormat
	orgID    string
	lessonID string

	// UI components
	textInput textinput.Model
	viewport  viewport.Model
	styles    *Styles

	// Command execution
	commandOutput string
	lastError     error
	currentHint   int

	// Dimensions
	width  int
	height int

	// Flags
	ready    bool
	quitting bool
	runner   *runner.Runner
}

// NewModel creates a new model
func NewModel(orgID, lessonID string, lesson lessons_pkg.LessonFormat) Model {
	ti := textinput.New()
	ti.Placeholder = "Type your command here..."
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 80

	theme := GetTheme("default")
	styles := NewStyles(theme)

	return Model{
		state:       StateIntroduction,
		currentStep: 0,
		totalSteps:  len(lesson.Steps),
		lesson:      lesson,
		orgID:       orgID,
		lessonID:    lessonID,
		textInput:   ti,
		styles:      styles,
		stepState:   StepPending,
		runner:      runner.NewRunner(),
	}
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

// Update handles messages
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			return m.handleEnter()

		case "esc":
			return m.handleEscape()

		case "?":
			return m.handleHintRequest()

		default:
			// Update text input for command steps
			if m.state == StateLesson &&
				m.lesson.Steps[m.currentStep].Type == "command" {
				m.textInput, cmd = m.textInput.Update(msg)
				return m, cmd
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-10)
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - 10
		}

	case CommandResultMsg:
		return m.handleCommandResult(msg)
	}

	return m, cmd
}

// handleEnter processes Enter key
func (m Model) handleEnter() (tea.Model, tea.Cmd) {
	switch m.state {
	case StateIntroduction:
		m.state = StateLesson
		m.currentStep = 0
		return m, nil

	case StateLesson:
		step := m.lesson.Steps[m.currentStep]

		switch step.Type {
		case "info":
			// Move to next step
			if m.currentStep < m.totalSteps-1 {
				m.currentStep++
				m.stepState = StepPending
				m.currentHint = 0
				m.commandOutput = ""
				m.textInput.Reset()
			} else {
				m.state = StateConclusion
			}
			return m, nil

		case "command":
			if m.textInput.Value() != "" {
				m.stepState = StepExecuting
				return m, executeCommand(m.textInput.Value(), step, m.runner)
			}
			return m, nil

		default:
			// Handle quiz, challenge, etc.
			if m.currentStep < m.totalSteps-1 {
				m.currentStep++
			} else {
				m.state = StateConclusion
			}
			return m, nil
		}

	case StateConclusion:
		m.quitting = true
		return m, tea.Quit
	}

	return m, nil
}

// handleEscape processes Escape key
func (m Model) handleEscape() (tea.Model, tea.Cmd) {
	if m.currentStep > 0 {
		m.currentStep--
		m.stepState = StepPending
		m.commandOutput = ""
		m.currentHint = 0
		m.textInput.Reset()
	}
	return m, nil
}

// handleHintRequest shows next hint
func (m Model) handleHintRequest() (tea.Model, tea.Cmd) {
	if m.state == StateLesson {
		step := m.lesson.Steps[m.currentStep]
		if step.Type == "command" && len(step.Hints) > 0 {
			if m.currentHint < len(step.Hints) {
				m.currentHint++
			}
		}
	}
	return m, nil
}

// handleCommandResult processes command execution result
func (m Model) handleCommandResult(msg CommandResultMsg) (tea.Model, tea.Cmd) {
	m.commandOutput = msg.Output
	m.lastError = msg.Error

	if msg.Success {
		m.stepState = StepSuccess
		// Auto-advance after short delay (2 seconds)
		return m, tea.Tick(2*time.Second, func(t time.Time) tea.Msg {
			return AdvanceStepMsg{}
		})
	} else {
		m.stepState = StepFailed
	}

	return m, nil
}

// View renders the UI
func (m Model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}

	if m.quitting {
		return "\n  Thanks for using TryOutShell! 👋\n\n"
	}

	switch m.state {
	case StateIntroduction:
		return m.renderIntroduction()
	case StateLesson:
		return m.renderLesson()
	case StateConclusion:
		return m.renderConclusion()
	}

	return ""
}

// Custom messages
type CommandResultMsg struct {
	Output  string
	Error   error
	Success bool
}

type AdvanceStepMsg struct{}

// executeCommand runs a shell command (stub for now)
func executeCommand(cmd string, step lessons_pkg.StepType, r *runner.Runner) tea.Cmd {
	return func() tea.Msg {
		// TODO: Implement actual command execution
		result := r.Execute(cmd, step.Timeout)
		success := false
		if result.Error == nil {
			success = r.Verify(result, step.Validation)
			if !success && len(step.AlternativeValidations) > 0 {
				for _, altVal := range step.AlternativeValidations {
					if r.Verify(result, altVal) {
						success = true
						break
					}
				}
			}
		}

		// For now, simulate success
		// Format output
		output := fmt.Sprintf("$ %s\n\n%s\n\n⏱ Completed in %.1fs",
			cmd,
			result.Output,
			result.Duration.Seconds(),
		)

		if result.Error != nil {
			output += fmt.Sprintf("\n\nExit code: %d", result.ExitCode)
		}

		return CommandResultMsg{
			Output:  output,
			Error:   result.Error,
			Success: success,
		}
	}
}
