package quiz

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tryoutshell/tryoutshell/internal/progress"
	"github.com/tryoutshell/tryoutshell/types"
)

type phase int

const (
	phaseQuestion phase = iota
	phaseFeedback
	phaseResults
)

type Model struct {
	questions []types.QuizQuestion
	current   int
	cursor    int
	selected  int
	phase     phase
	correct   int
	answers   []int
	width     int
	height    int

	orgID    string
	lessonID string
}

func New(orgID, lessonID string, questions []types.QuizQuestion) Model {
	return Model{
		questions: questions,
		orgID:     orgID,
		lessonID:  lessonID,
		selected:  -1,
		answers:   make([]int, len(questions)),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch m.phase {
		case phaseQuestion:
			return m.updateQuestion(msg)
		case phaseFeedback:
			return m.updateFeedback(msg)
		case phaseResults:
			return m.updateResults(msg)
		}
	}
	return m, nil
}

func (m Model) updateQuestion(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	q := m.questions[m.current]
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(q.Options)-1 {
			m.cursor++
		}
	case "enter", " ":
		m.selected = m.cursor
		m.answers[m.current] = m.selected
		if m.selected == q.Answer {
			m.correct++
		}
		m.phase = phaseFeedback
	}
	return m, nil
}

func (m Model) updateFeedback(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "enter", " ", "n":
		if m.current+1 < len(m.questions) {
			m.current++
			m.cursor = 0
			m.selected = -1
			m.phase = phaseQuestion
		} else {
			m.phase = phaseResults
			m.saveScore()
		}
	}
	return m, nil
}

func (m Model) updateResults(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c", "enter":
		return m, tea.Quit
	}
	return m, nil
}

func (m Model) saveScore() {
	store := progress.NewStore()
	store.SaveQuizScore(m.orgID, m.lessonID, m.correct, len(m.questions))
}

func (m Model) View() string {
	switch m.phase {
	case phaseQuestion:
		return m.viewQuestion()
	case phaseFeedback:
		return m.viewFeedback()
	case phaseResults:
		return m.viewResults()
	}
	return ""
}

func (m Model) viewQuestion() string {
	q := m.questions[m.current]

	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63")).Padding(1, 2)
	counterStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Padding(0, 2)
	optionStyle := lipgloss.NewStyle().Padding(0, 4)
	cursorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("63")).Bold(true)

	var b strings.Builder
	b.WriteString(titleStyle.Render(fmt.Sprintf("Quiz: %s", q.Question)))
	b.WriteString("\n")
	b.WriteString(counterStyle.Render(fmt.Sprintf("Question %d of %d", m.current+1, len(m.questions))))
	b.WriteString("\n\n")

	for i, opt := range q.Options {
		prefix := "  "
		if i == m.cursor {
			prefix = cursorStyle.Render("▸ ")
		}
		b.WriteString(optionStyle.Render(fmt.Sprintf("%s%c) %s", prefix, 'A'+rune(i), opt)))
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Padding(0, 2).Render("↑/↓ navigate • enter select • q quit"))

	return b.String()
}

func (m Model) viewFeedback() string {
	q := m.questions[m.current]
	isCorrect := m.selected == q.Answer

	var b strings.Builder

	titleStyle := lipgloss.NewStyle().Bold(true).Padding(1, 2)
	explainStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Padding(0, 4).Width(m.width - 8)

	if isCorrect {
		b.WriteString(titleStyle.Foreground(lipgloss.Color("42")).Render("✓ Correct!"))
	} else {
		b.WriteString(titleStyle.Foreground(lipgloss.Color("196")).Render("✗ Incorrect"))
		b.WriteString("\n")
		correctOpt := ""
		if q.Answer >= 0 && q.Answer < len(q.Options) {
			correctOpt = q.Options[q.Answer]
		}
		b.WriteString(lipgloss.NewStyle().Padding(0, 4).Render(
			fmt.Sprintf("The correct answer was: %c) %s", 'A'+rune(q.Answer), correctOpt),
		))
	}

	if q.Explain != "" {
		b.WriteString("\n\n")
		b.WriteString(explainStyle.Render(q.Explain))
	}

	b.WriteString("\n\n")
	b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Padding(0, 2).Render("press enter to continue"))

	return b.String()
}

func (m Model) viewResults() string {
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63")).Padding(1, 2)
	scoreStyle := lipgloss.NewStyle().Bold(true).Padding(0, 2)

	pct := 0.0
	if len(m.questions) > 0 {
		pct = float64(m.correct) / float64(len(m.questions)) * 100
	}

	var b strings.Builder
	b.WriteString(titleStyle.Render("📊 Quiz Complete!"))
	b.WriteString("\n\n")

	color := lipgloss.Color("196")
	if pct >= 80 {
		color = lipgloss.Color("42")
	} else if pct >= 50 {
		color = lipgloss.Color("226")
	}

	b.WriteString(scoreStyle.Foreground(color).Render(
		fmt.Sprintf("  Score: %d / %d  (%.0f%%)", m.correct, len(m.questions), pct),
	))
	b.WriteString("\n\n")

	for i, q := range m.questions {
		icon := lipgloss.NewStyle().Foreground(lipgloss.Color("42")).Render("✓")
		if m.answers[i] != q.Answer {
			icon = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render("✗")
		}
		b.WriteString(fmt.Sprintf("    %s %s\n", icon, q.Question))
	}

	b.WriteString("\n")
	b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Padding(0, 2).Render("press q or enter to exit"))

	return b.String()
}

func Launch(orgID, lessonID string, questions []types.QuizQuestion) error {
	m := New(orgID, lessonID, questions)
	p := tea.NewProgram(m, tea.WithAltScreen())
	_, err := p.Run()
	return err
}
