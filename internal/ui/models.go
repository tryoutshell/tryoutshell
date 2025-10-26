package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	lessons_pkg "github.com/tryoutshell/tryoutshell/internal/lessons"
	"github.com/tryoutshell/tryoutshell/internal/runner"
)

// Key bindings
type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
	Quit  key.Binding
	Esc   key.Binding
	Help  key.Binding
	Skip  key.Binding
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "scroll up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "scroll down"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "continue/execute"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
	Esc: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "hint"),
	),
	Skip: key.NewBinding(
		key.WithKeys("ctrl+y"),
		key.WithHelp("ctrl+y", "skip"),
	),
}

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

	// Quiz state
	quizAnswers    map[string]int
	currentQuizQ   int
	selectedOption int

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

	// Runner
	runner      *runner.Runner
	sandboxInfo string
	//  store full result
	lastCommandResult runner.CommandResult
}

// NewModel creates a new model
func NewModel(orgID, lessonID string, lesson lessons_pkg.LessonFormat) Model {
	ti := textinput.New()
	ti.Placeholder = "Type command or ':skip' to skip..."
	ti.Focus()
	ti.CharLimit = 512
	ti.Width = 80
	ti.Prompt = ""
	ti.SetCursorMode(textinput.CursorBlink)

	theme := GetTheme("default")
	styles := NewStyles(theme)

	// NEW: Create runner and setup sandbox
	r := runner.NewRunner()
	sandboxInfo := ""
	if r.IsSandboxed() {
		sandboxInfo = fmt.Sprintf("🏖️  Sandbox: %s", r.GetWorkingDir())

		if err := r.SetupLesson(lessonID); err != nil {
			sandboxInfo += fmt.Sprintf("\n⚠️  Setup warning: %v", err)
		}
	} else {
		sandboxInfo = "⚠️  Running in current directory (sandbox creation failed)"
	}

	return Model{
		state:          StateIntroduction,
		currentStep:    0,
		totalSteps:     len(lesson.Steps),
		lesson:         lesson,
		orgID:          orgID,
		lessonID:       lessonID,
		textInput:      ti,
		styles:         styles,
		stepState:      StepPending,
		quizAnswers:    make(map[string]int),
		selectedOption: 0,
		runner:         r,           // CHANGED
		sandboxInfo:    sandboxInfo, // NEW
	}
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

// Update handles messages
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Get current step info
		var currentStepType string
		var isCommandInputActive bool

		if m.state == StateLesson && m.currentStep < len(m.lesson.Steps) {
			currentStepType = m.lesson.Steps[m.currentStep].Type
			isCommandInputActive = currentStepType == "command" &&
				(m.stepState == StepPending || m.stepState == StepFailed)
		}

		// Global keys that work everywhere
		if key.Matches(msg, keys.Quit) {
			m.quitting = true
			return m, tea.Batch(
				tea.Quit,
				cleanupSandbox(m.runner), // Clean up before quitting
			)
		}

		// Special handling for quiz steps
		if currentStepType == "quiz" {
			switch msg.String() {
			case "q", "ctrl+c":
				m.quitting = true
				return m, tea.Quit

			case "enter":
				return m.handleEnter()

			case "esc":
				return m.handleEscape()

			case "up", "k":
				// Quiz navigation
				if m.selectedOption > 0 {
					m.selectedOption--
					m.updateViewportContent()
				}
				return m, nil

			case "down", "j":
				// Quiz navigation
				step := m.lesson.Steps[m.currentStep]
				if m.currentQuizQ < len(step.Questions) &&
					m.selectedOption < len(step.Questions[m.currentQuizQ].Options)-1 {
					m.selectedOption++
					m.updateViewportContent()
				}
				return m, nil

			case "ctrl+y":
				return m.handleSkip()

			default:
				// Ignore other keys in quiz mode
				return m, nil
			}
		}

		// Special handling for command input (text entry)
		if isCommandInputActive {
			switch msg.String() {
			case "q":
				// In command input mode, 'q' should be typed, not quit
				// Only quit with ctrl+c
				m.textInput, cmd = m.textInput.Update(msg)
				m.updateViewportContent() // UPDATE VIEWPORT AFTER TYPING
				return m, cmd

			case "ctrl+c":
				m.quitting = true
				return m, tea.Quit

			case "enter":
				return m.handleEnter()

			case "esc":
				return m.handleEscape()

			case "ctrl+y":
				return m.handleSkip()

			case "?":
				return m.handleHintRequest()

			case "up", "down", "left", "right":
				// Allow cursor movement in text input
				m.textInput, cmd = m.textInput.Update(msg)
				m.updateViewportContent() // UPDATE VIEWPORT AFTER CURSOR MOVE
				return m, cmd

			default:
				// Pass all other keys to text input for typing
				m.textInput, cmd = m.textInput.Update(msg)
				m.updateViewportContent() // UPDATE VIEWPORT AFTER EVERY KEYSTROKE
				return m, cmd
			}
		}

		// Default key handling for info steps and other content
		switch {
		case key.Matches(msg, keys.Enter):
			return m.handleEnter()

		case key.Matches(msg, keys.Esc):
			return m.handleEscape()

		case key.Matches(msg, keys.Help):
			return m.handleHintRequest()

		case key.Matches(msg, keys.Up):
			m.viewport.LineUp(3)
			return m, nil

		case key.Matches(msg, keys.Down):
			m.viewport.LineDown(3)
			return m, nil

		case key.Matches(msg, keys.Skip):
			return m.handleSkip()
		}

	case DebugCommandResultMsg:
		// Show debug output but don't change step state
		m.commandOutput = msg.Output
		m.updateViewportContent()
		return m, nil

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		headerHeight := 3
		footerHeight := 5

		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-headerHeight-footerHeight)
			m.viewport.YPosition = headerHeight
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - headerHeight - footerHeight
		}

		m.updateViewportContent()

	case CommandResultMsg:
		return m.handleCommandResult(msg)

	case AdvanceStepMsg:
		return m.advanceStep()
	}

	// Update viewport (but not when typing in text input)
	if m.state != StateLesson ||
		m.currentStep >= len(m.lesson.Steps) ||
		m.lesson.Steps[m.currentStep].Type != "command" ||
		(m.stepState != StepPending && m.stepState != StepFailed) {
		m.viewport, cmd = m.viewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

// Add a cleanup message type
type CleanupCompleteMsg struct{}

func cleanupSandbox(r *runner.Runner) tea.Cmd {
	return func() tea.Msg {
		r.Cleanup()
		return CleanupCompleteMsg{}
	}
}

// handleEnter processes Enter key
func (m Model) handleEnter() (tea.Model, tea.Cmd) {
	switch m.state {
	case StateIntroduction:
		m.state = StateLesson
		m.currentStep = 0
		m.updateViewportContent()
		return m, nil

	case StateLesson:
		step := m.lesson.Steps[m.currentStep]

		switch step.Type {
		case "info":
			return m.advanceStep()

		case "command":
			userInput := strings.TrimSpace(m.textInput.Value())

			// Check for skip command
			if userInput == ":skip" {
				return m.handleSkip()
			}

			if userInput != "" && (m.stepState == StepPending || m.stepState == StepFailed) {
				m.stepState = StepExecuting
				m.updateViewportContent()
				return m, m.executeCommand(userInput, step)
			}

			// If already successful, advance
			if m.stepState == StepSuccess {
				return m.advanceStep()
			}
			return m, nil

		case "quiz":
			// Submit quiz answer
			if m.currentQuizQ < len(step.Questions) {
				q := step.Questions[m.currentQuizQ]
				m.quizAnswers[q.ID] = m.selectedOption

				m.currentQuizQ++
				m.selectedOption = 0

				// If all questions answered, advance
				if m.currentQuizQ >= len(step.Questions) {
					m.currentQuizQ = 0
					return m.advanceStep()
				}

				m.updateViewportContent()
			}
			return m, nil

		case "challenge":
			// Verify challenge completion
			return m.advanceStep()

		case "interview_prep":
			return m.advanceStep()

		default:
			return m.advanceStep()
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
		m.currentQuizQ = 0
		m.selectedOption = 0
		m.textInput.Reset()
		m.updateViewportContent()
	}
	return m, nil
}

// handleSkip skips the current step
func (m Model) handleSkip() (tea.Model, tea.Cmd) {
	if m.state != StateLesson || m.currentStep >= len(m.lesson.Steps) {
		return m, nil
	}

	step := m.lesson.Steps[m.currentStep]

	// Allow skip for command steps (always allow in learning environment)
	if step.Type == "command" {
		m.stepState = StepSuccess
		m.commandOutput = "⏭️  Skipped"
		m.textInput.SetValue(":skipped")
		m.updateViewportContent()

		// Show brief message then advance
		return m, tea.Tick(1*time.Second, func(t time.Time) tea.Msg {
			return AdvanceStepMsg{}
		})
	}

	// For other steps, just advance
	return m.advanceStep()
}

// handleHintRequest shows next hint
func (m Model) handleHintRequest() (tea.Model, tea.Cmd) {
	if m.state == StateLesson {
		step := m.lesson.Steps[m.currentStep]
		if step.Type == "command" && len(step.Hints) > 0 {
			if m.currentHint < len(step.Hints) {
				m.currentHint++
				m.updateViewportContent()
			}
		}
	}
	return m, nil
}

// handleCommandResult processes command execution result
func (m Model) handleCommandResult(msg CommandResultMsg) (tea.Model, tea.Cmd) {
	m.commandOutput = msg.Output
	m.lastError = msg.Error
	m.lastCommandResult = msg.FullResult // STORE FULL RESULT

	if msg.Success {
		m.stepState = StepSuccess
	} else {
		m.stepState = StepFailed
	}

	m.updateViewportContent()
	return m, nil
}

// advanceStep moves to next step
func (m Model) advanceStep() (tea.Model, tea.Cmd) {
	if m.currentStep < m.totalSteps-1 {
		m.currentStep++
		m.stepState = StepPending
		m.commandOutput = ""
		m.currentHint = 0
		m.currentQuizQ = 0
		m.selectedOption = 0
		m.textInput.Reset()
		m.viewport.GotoTop()
		m.updateViewportContent()
	} else {
		m.state = StateConclusion
		m.updateViewportContent()
	}
	return m, nil
}

// updateViewportContent refreshes viewport content
func (m *Model) updateViewportContent() {
	if !m.ready {
		return
	}

	var content string
	switch m.state {
	case StateIntroduction:
		content = m.renderIntroduction()
	case StateLesson:
		content = m.renderLesson()
	case StateConclusion:
		content = m.renderConclusion()
	}

	// // Center content horizontally
	// lines := strings.Split(content, "\n")
	// var centeredLines []string
	// for _, line := range lines {
	// 	lineWidth := lipgloss.Width(line)
	// 	if lineWidth < m.width {
	// 		padding := (m.width - lineWidth) / 2
	// 		centeredLines = append(centeredLines, strings.Repeat(" ", padding)+line)
	// 	} else {
	// 		centeredLines = append(centeredLines, line)
	// 	}
	// }

	// m.viewport.SetContent(strings.Join(centeredLines, "\n"))
	// Don't center - keep natural alignment for better readability
	m.viewport.SetContent(content)
}

// Fix the View() method
func (m Model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}

	if m.quitting {
		// Fix: Use direct style instead of non-existent method
		quitStyle := lipgloss.NewStyle().Foreground(m.styles.Theme.Success).Bold(true)
		return quitStyle.Render("\n  Thanks for using TryOutShell! 👋\n\n")
	}

	// Header
	header := m.renderHeader()

	// Footer with help
	footer := m.renderFooter()

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		m.viewport.View(),
		footer,
	)
}

// renderHeader renders the top header
func (m Model) renderHeader() string {
	title := m.styles.AppTitle.Width(m.width).Render(
		fmt.Sprintf("TryOutShell | %s - %s", m.lesson.Metadata.Org, m.lesson.Metadata.Title),
	)
	return title
}

// renderFooter renders the bottom footer
// renderFooter renders the bottom footer
func (m Model) renderFooter() string {
	step := ""
	stepState := ""

	if m.state == StateLesson && m.currentStep < len(m.lesson.Steps) {
		currentStep := m.lesson.Steps[m.currentStep]
		step = currentStep.Type

		// Debug info (remove after fixing)
		switch m.stepState {
		case StepPending:
			stepState = "PENDING"
		case StepExecuting:
			stepState = "EXECUTING"
		case StepSuccess:
			stepState = "SUCCESS"
		case StepFailed:
			stepState = "FAILED"
		}
	}

	help := m.renderHelpText(step)
	progress := m.renderProgressBar()

	// Debug line (remove after fixing)
	debug := m.styles.Muted.Render(fmt.Sprintf("  [Debug: State=%v | Step=%s | StepState=%s | QuizQ=%d | QuizOpt=%d]",
		m.state, step, stepState, m.currentQuizQ, m.selectedOption))

	footerContent := lipgloss.JoinVertical(
		lipgloss.Left,
		progress,
		help,
		debug, // Remove this line after fixing
	)

	return m.styles.HelpText.Width(m.width).Render(footerContent)
}

// Custom messages
type CommandResultMsg struct {
	Output     string
	Error      error
	Success    bool
	FullResult runner.CommandResult
}

type AdvanceStepMsg struct{}

// executeCommand runs a shell command
func (m Model) executeCommand(cmd string, step lessons_pkg.StepType) tea.Cmd {
	return func() tea.Msg {
		// Execute command
		result := m.runner.Execute(cmd, step.Timeout)

		// Verify result WITH detailed validation info
		result, success := m.runner.Verify(result, step.Validation)

		// Try alternative validations if primary fails
		if !success && len(step.AlternativeValidations) > 0 {
			for _, altVal := range step.AlternativeValidations {
				result, success = m.runner.Verify(result, altVal)
				if success {
					break
				}
			}
		}

		// Format output for display
		output := fmt.Sprintf("$ %s\n\n%s",
			cmd,
			result.Output,
		)

		if result.Duration > 0 {
			output += fmt.Sprintf("\n\n⏱  Completed in %.2fs", result.Duration.Seconds())
		}

		return CommandResultMsg{
			Output:     output,
			Error:      result.Error,
			Success:    success,
			FullResult: result, // ADD THIS
		}
	}
}

// executeDebugCommand runs debug commands without validation
func (m Model) executeDebugCommand(cmd string) tea.Cmd {
	return func() tea.Msg {
		// Execute debug command
		result := m.runner.Execute(cmd, 5) // 5 second timeout

		// Format output
		output := fmt.Sprintf("$ %s\n\n%s", cmd, result.Output)

		// Return as success (debug commands don't affect step progress)
		return DebugCommandResultMsg{
			Output:  output,
			Command: cmd,
		}
	}
}

// Add new message type
type DebugCommandResultMsg struct {
	Output  string
	Command string
}
