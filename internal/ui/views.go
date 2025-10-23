package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	lessons_pkg "github.com/tryoutshell/tryoutshell/internal/lessons"
	"regexp"
	"strings"
)

// renderIntroduction renders the introduction screen
func (m Model) renderIntroduction() string {
	var b strings.Builder

	// Header
	header := m.styles.SectionHeader.Render(
		fmt.Sprintf("╔═══════════════════════════════════════════════════════════════════════╗\n"+
			"║  %s\n"+
			"║  by %s  •  v%s\n"+
			"╚═══════════════════════════════════════════════════════════════════════╝",
			m.lesson.Metadata.Title,
			m.lesson.Metadata.Author,
			m.lesson.Metadata.Version,
		),
	)
	b.WriteString(header + "\n\n")

	// Introduction content
	introBox := m.styles.Border.Render(
		m.styles.Bold.Render(m.lesson.Introduction.Title) + "\n\n" +
			m.formatMarkdown(m.lesson.Introduction.Content),
	)
	b.WriteString(introBox + "\n\n")

	// Progress bar
	b.WriteString(m.renderProgressBar() + "\n\n")

	// Help text
	help := m.styles.HelpText.Render(
		"  Press Enter to continue  •  q: Quit",
	)
	b.WriteString(help + "\n")

	return b.String()
}

// renderLesson renders the current lesson step
func (m Model) renderLesson() string {
	if m.currentStep >= len(m.lesson.Steps) {
		m.state = StateConclusion
		return m.renderConclusion()
	}

	step := m.lesson.Steps[m.currentStep]

	var b strings.Builder

	// Step header
	header := m.styles.SectionHeader.Render(
		fmt.Sprintf("╔═══════════════════════════════════════════════════════════════════════╗\n"+
			"║  Step %d/%d%s\n"+
			"╚═══════════════════════════════════════════════════════════════════════╝",
			m.currentStep+1,
			m.totalSteps,
			m.getStepTypeLabel(step.Type),
		),
	)
	b.WriteString(header + "\n\n")

	// Render based on step type
	switch step.Type {
	case "info":
		b.WriteString(m.renderInfoStep(step))
	case "command":
		b.WriteString(m.renderCommandStep(step))
	case "quiz":
		b.WriteString(m.renderQuizStep(step))
	case "challenge":
		b.WriteString(m.renderChallengeStep(step))
	default:
		b.WriteString("Unknown step type\n")
	}

	// Progress bar
	b.WriteString("\n" + m.renderProgressBar() + "\n\n")

	// Help text
	b.WriteString(m.renderHelpText(step.Type) + "\n")

	return b.String()
}

// renderInfoStep renders an info step
func (m Model) renderInfoStep(step lessons_pkg.StepType) string {
	var b strings.Builder

	// Title
	title := m.styles.StepTitle.Render(step.Title)
	b.WriteString(title + "\n\n")

	// Content
	content := m.formatMarkdown(step.Content)
	b.WriteString(content + "\n")

	// Code blocks
	for _, block := range step.CodeBlocks {
		b.WriteString("\n" + m.renderCodeBlock(block) + "\n")
	}

	// Callouts
	for _, callout := range step.Callouts {
		b.WriteString("\n" + m.renderCallout(callout) + "\n")
	}

	// Diagram
	if step.Diagram != "" {
		b.WriteString("\n" + m.renderDiagram(step.Diagram) + "\n")
	}

	return b.String()
}

// renderCommandStep renders a command execution step
func (m Model) renderCommandStep(step lessons_pkg.StepType) string {
	var b strings.Builder

	// Title
	title := m.styles.StepTitle.Render(step.Prompt)
	b.WriteString(title + "\n\n")

	// Pre-content
	if step.PreContent != "" {
		b.WriteString(m.formatMarkdown(step.PreContent) + "\n\n")
	}

	// Instruction
	if step.Instruction != "" {
		instruction := m.styles.Paragraph.Render(step.Instruction)
		b.WriteString(instruction + "\n\n")
	}

	// Example command
	if step.Example != "" {
		exampleBox := m.styles.CodeBlock.Render(
			m.styles.Bold.Render("╭─ Example ─────────────────────────────────────────────────────────────╮") +
				"\n│\n│  " + m.styles.CommandPrompt.Render() + step.Example +
				"\n│\n" +
				m.styles.Bold.Render("╰────────────────────────────────────────────────────────────────────────╯"),
		)
		b.WriteString(exampleBox + "\n\n")
	}

	// Command input or output
	if m.stepState == StepPending || m.stepState == StepFailed {
		// Show input box
		inputBox := m.styles.CommandInput.Render(
			m.styles.Bold.Render("┌─ Your turn: ──────────────────────────────────────────────────────────┐") +
				"\n│\n│  " + m.styles.CommandPrompt.Render() + m.textInput.View() +
				"\n│\n" +
				m.styles.Bold.Render("└────────────────────────────────────────────────────────────────────────┘"),
		)
		b.WriteString(inputBox + "\n")

		// Show hints if requested
		if m.currentHint > 0 && m.currentHint <= len(step.Hints) {
			hint := step.Hints[m.currentHint-1]
			hintBox := m.styles.HintBox.Render(
				fmt.Sprintf("╭─ Hint (Level %d/%d) ────────────────────────────────────────────────────╮\n"+
					"│\n│  💡 %s\n│\n"+
					"╰────────────────────────────────────────────────────────────────────────╯",
					m.currentHint, len(step.Hints), hint.Text),
			)
			b.WriteString("\n" + hintBox + "\n")
		}

		// Show error message if failed
		if m.stepState == StepFailed {
			errorMsg := m.styles.ErrorMsg.Render("❌ " + step.FailMsg)
			b.WriteString("\n" + errorMsg + "\n")
		}

	} else if m.stepState == StepExecuting {
		b.WriteString(m.styles.Paragraph.Render("⏳ Executing command...") + "\n")

	} else if m.stepState == StepSuccess {
		// Show command that was executed
		cmdBox := m.styles.Border.Render(
			m.styles.Bold.Render("┌─ Your command: ───────────────────────────────────────────────────────┐") +
				"\n│\n│  " + m.styles.CommandPrompt.Render() + m.textInput.Value() +
				"\n│\n" +
				m.styles.Bold.Render("└────────────────────────────────────────────────────────────────────────┘"),
		)
		b.WriteString(cmdBox + "\n\n")

		// Show output
		if m.commandOutput != "" {
			outputBox := m.styles.OutputSuccess.Render(
				m.styles.Bold.Render("╭─ Output ──────────────────────────────────────────────────────────────╮") +
					"\n│\n│  " + strings.ReplaceAll(m.commandOutput, "\n", "\n│  ") +
					"\n│\n" +
					m.styles.Bold.Render("╰────────────────────────────────────────────────────────────────────────╯"),
			)
			b.WriteString(outputBox + "\n\n")
		}

		// Success message
		successMsg := m.styles.SuccessMsg.Render("✅ " + step.SuccessMsg)
		b.WriteString(successMsg + "\n")

		// Post-content
		if step.PostContent != "" {
			b.WriteString("\n" + m.formatMarkdown(step.PostContent) + "\n")
		}
	}

	return b.String()
}

// renderQuizStep renders a quiz step (placeholder)
func (m Model) renderQuizStep(step lessons_pkg.StepType) string {
	return m.styles.Border.Render("Quiz functionality coming soon...")
}

// renderChallengeStep renders a challenge step (placeholder)
func (m Model) renderChallengeStep(step lessons_pkg.StepType) string {
	return m.styles.Border.Render("Challenge functionality coming soon...")
}

// renderConclusion renders the conclusion screen
func (m Model) renderConclusion() string {
	var b strings.Builder

	// Celebration header
	header := m.styles.AppTitle.Render(
		"╔═══════════════════════════════════════════════════════════════════════╗\n" +
			"║  🎉 Lesson Complete!                                                   ║\n" +
			"╚═══════════════════════════════════════════════════════════════════════╝",
	)
	b.WriteString(header + "\n\n")

	// Conclusion content
	conclusionBox := m.styles.Border.Render(
		m.styles.Bold.Render(m.lesson.Conclusion.Title) + "\n\n" +
			m.formatMarkdown(m.lesson.Conclusion.Content),
	)
	b.WriteString(conclusionBox + "\n\n")

	// Badges
	if len(m.lesson.Conclusion.Badges) > 0 {
		b.WriteString(m.styles.Bold.Render("🏆 Badges Earned:") + "\n\n")
		for _, badge := range m.lesson.Conclusion.Badges {
			b.WriteString(fmt.Sprintf("  %s %s\n", badge.Icon, badge.Name))
		}
		b.WriteString("\n")
	}

	// Help text
	help := m.styles.HelpText.Render("  Press Enter to exit  •  q: Quit")
	b.WriteString(help + "\n")

	return b.String()
}

// Helper rendering functions

func (m Model) renderProgressBar() string {
	progress := float64(m.currentStep) / float64(m.totalSteps)
	filled := int(progress * 20)
	empty := 20 - filled

	// Create style for filled portion
	filledStyle := lipgloss.NewStyle().Foreground(m.styles.Theme.Success)

	bar := "Progress: [" +
		filledStyle.Render(strings.Repeat("█", filled)) +
		strings.Repeat("░", empty) + "]" +
		fmt.Sprintf(" %d%% (%d/%d steps)", int(progress*100), m.currentStep, m.totalSteps)

	return bar
}

func (m Model) renderCodeBlock(block lessons_pkg.CodeBlockType) string {
	return m.styles.CodeBlock.Render(
		m.styles.Bold.Render("╭─ "+block.Label+" ──────────────────────────────────────────────────────────────╮") +
			"\n│\n│  " + m.styles.CommandPrompt.Render() + strings.ReplaceAll(block.Code, "\n", "\n│  ") +
			"\n│\n" +
			m.styles.Bold.Render("╰────────────────────────────────────────────────────────────────────────╯"),
	)
}

func (m Model) renderCallout(callout lessons_pkg.CalloutType) string {
	var icon string
	var style lipgloss.Style

	switch callout.Type {
	case "tip":
		icon = "💡"
		style = m.styles.CalloutTip
	case "warning":
		icon = "⚠️"
		style = m.styles.CalloutWarning
	default:
		icon = "ℹ️"
		style = m.styles.CalloutTip
	}

	return style.Render(
		fmt.Sprintf("┌─────────────────────────────────────────────────────────────────────┐\n"+
			"│  %s %s\n"+
			"└─────────────────────────────────────────────────────────────────────┘",
			icon, callout.Text),
	)
}

func (m Model) renderDiagram(diagram string) string {
	return m.styles.Border.Render(diagram)
}

func (m Model) getStepTypeLabel(stepType string) string {
	switch stepType {
	case "command":
		return "  •  Command Execution"
	case "quiz":
		return "  •  Quiz"
	case "challenge":
		return "  •  Challenge"
	default:
		return ""
	}
}

func (m Model) renderHelpText(stepType string) string {
	switch stepType {
	case "info":
		return m.styles.HelpText.Render("  Press Enter to continue  •  Esc: Back  •  q: Quit")
	case "command":
		return m.styles.HelpText.Render("  Type command and press Enter  •  ?: Hint  •  Esc: Back  •  q: Quit")
	default:
		return m.styles.HelpText.Render("  Enter: Continue  •  Esc: Back  •  q: Quit")
	}
}

// formatMarkdown applies basic markdown formatting
func (m Model) formatMarkdown(text string) string {
	// Bold (**text**)
	re := regexp.MustCompile(`\*\*(.*?)\*\*`)
	text = re.ReplaceAllStringFunc(text, func(match string) string {
		content := strings.Trim(match, "*")
		return m.styles.Bold.Render(content)
	})

	// Inline code (`code`)
	re = regexp.MustCompile("`([^`]+)`")
	text = re.ReplaceAllStringFunc(text, func(match string) string {
		code := strings.Trim(match, "`")
		return m.styles.InlineCode.Render(code)
	})

	// Bullet points
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "- ") {
			lines[i] = "  • " + strings.TrimPrefix(strings.TrimSpace(line), "- ")
		}
	}
	text = strings.Join(lines, "\n")

	return text
}
