package ui

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
	lessons_pkg "github.com/tryoutshell/tryoutshell/internal/lessons"
)

// renderIntroduction renders the introduction screen
func (m Model) renderIntroduction() string {
	var b strings.Builder

	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(m.styles.Theme.Primary).
		Padding(1, 2).
		Render(m.lesson.Metadata.Title)

	b.WriteString(title + "\n\n")

	meta := fmt.Sprintf("📚 %s  |  ⭐ %s  |  ⏱  %s",
		m.lesson.Metadata.Author,
		m.lesson.Metadata.Difficulty,
		m.lesson.Metadata.Duration,
	)
	b.WriteString(m.styles.HelpText.Render(meta) + "\n\n")

	introBox := m.styles.Border.Width(m.getContentWidth()).Render(
		m.styles.Bold.Render(m.lesson.Introduction.Title) + "\n\n" +
			m.formatMarkdown(m.lesson.Introduction.Content),
	)
	b.WriteString(introBox + "\n")

	return b.String()
}

// renderLesson renders the current lesson step
func (m Model) renderLesson() string {
	if m.currentStep >= len(m.lesson.Steps) {
		return m.renderConclusion()
	}

	step := m.lesson.Steps[m.currentStep]

	var b strings.Builder

	stepHeader := m.styles.SectionHeader.Width(m.getContentWidth()).Render(
		fmt.Sprintf("Step %d/%d  •  %s", m.currentStep+1, m.totalSteps, strings.ToUpper(step.Type)),
	)
	b.WriteString(stepHeader + "\n\n")

	switch step.Type {
	case "info":
		b.WriteString(m.renderInfoStep(step))
	case "command":
		b.WriteString(m.renderCommandStep(step))
	case "quiz":
		b.WriteString(m.renderQuizStep(step))
	case "challenge":
		b.WriteString(m.renderChallengeStep(step))
	case "interview_prep":
		b.WriteString(m.renderInterviewStep(step))
	default:
		b.WriteString("Unknown step type: " + step.Type + "\n")
	}

	return b.String()
}

// renderCommandStep renders a command execution step
func (m Model) renderCommandStep(step lessons_pkg.StepType) string {
	var b strings.Builder

	width := m.getContentWidth()

	// Title
	title := m.styles.StepTitle.Width(width).Render("  " + step.Prompt + "  ")
	b.WriteString("\n" + title + "\n\n")

	// Pre-content
	if step.PreContent != "" {
		preContent := m.formatMarkdown(step.PreContent)
		lines := strings.Split(preContent, "\n")
		for _, line := range lines {
			b.WriteString("  " + line + "\n")
		}
		b.WriteString("\n")
	}

	// Instruction
	if step.Instruction != "" {
		instruction := lipgloss.NewStyle().
			Foreground(m.styles.Theme.Primary).
			Render("  → " + step.Instruction)
		b.WriteString(instruction + "\n\n")
	}

	// Example command
	if step.Example != "" {
		exampleHeader := m.styles.Muted.Render("  Example:")
		b.WriteString(exampleHeader + "\n\n")

		exampleContent := "    " + m.styles.CommandPrompt.Render() + step.Example
		exampleBox := m.styles.CodeBlock.Width(width - 4).Render(exampleContent)
		b.WriteString("  " + exampleBox + "\n\n")
	}

	// Input box or output
	if m.stepState == StepPending || m.stepState == StepFailed {
		inputLabel := m.styles.Bold.Render("  Your turn:")
		b.WriteString(inputLabel + "\n\n")

		// Show input with prompt
		inputContent := "    " + m.styles.CommandPrompt.Render() + m.textInput.View()
		inputBox := m.styles.CommandInput.Width(width - 4).Render(inputContent)
		b.WriteString("  " + inputBox + "\n\n")

		// Skip hint
		skipHint := m.styles.Muted.Render("  💡 Type ':skip' or press Ctrl+Y to skip | Press '?' for hints")
		b.WriteString(skipHint + "\n")

		// Show hints
		if m.currentHint > 0 && m.currentHint <= len(step.Hints) {
			hint := step.Hints[m.currentHint-1]
			hintContent := fmt.Sprintf("💡 Hint %d/%d\n\n%s", m.currentHint, len(step.Hints), hint.Text)
			hintBox := m.styles.HintBox.Width(width - 6).Render(hintContent)
			b.WriteString("\n  " + hintBox + "\n")
		}

		// Error message
		if m.stepState == StepFailed {
			errorMsg := m.styles.ErrorMsg.Render("\n  ❌ " + step.FailMsg)
			b.WriteString(errorMsg + "\n\n")
			b.WriteString(m.styles.Muted.Render("  Press Enter to try again...") + "\n")
		}

	} else if m.stepState == StepExecuting {
		executing := lipgloss.NewStyle().
			Foreground(m.styles.Theme.Warning).
			Render("  ⏳ Executing command...")
		b.WriteString(executing + "\n")

	} else if m.stepState == StepSuccess {
		// Command executed
		cmdLabel := m.styles.Muted.Render("  Command executed:")
		b.WriteString(cmdLabel + "\n\n")

		cmdContent := "    " + m.styles.CommandPrompt.Render() + m.textInput.Value()
		cmdBox := m.styles.Border.Width(width - 4).Render(cmdContent)
		b.WriteString("  " + cmdBox + "\n\n")

		// Output
		if m.commandOutput != "" {
			outputLabel := m.styles.Muted.Render("  Output:")
			b.WriteString(outputLabel + "\n\n")

			// Add padding to output lines
			outputLines := strings.Split(m.commandOutput, "\n")
			var paddedOutput []string
			for _, line := range outputLines {
				paddedOutput = append(paddedOutput, "    "+line)
			}

			outputBox := m.styles.OutputSuccess.Width(width - 4).Render(strings.Join(paddedOutput, "\n"))
			b.WriteString("  " + outputBox + "\n\n")
		}

		// Success message
		successMsg := m.styles.SuccessMsg.Render("  ✅ " + step.SuccessMsg)
		b.WriteString(successMsg + "\n")

		// Post-content
		if step.PostContent != "" {
			b.WriteString("\n")
			postLines := strings.Split(m.formatMarkdown(step.PostContent), "\n")
			for _, line := range postLines {
				b.WriteString("  " + line + "\n")
			}
		}

		// Continue hint
		continueHint := m.styles.Muted.Render("\n  Press Enter to continue...")
		b.WriteString(continueHint + "\n")
	}

	return b.String()
}

// renderQuizStep renders a quiz step
// renderQuizStep with better formatting
func (m Model) renderQuizStep(step lessons_pkg.StepType) string {
	var b strings.Builder

	width := m.getContentWidth()

	if len(step.Questions) == 0 {
		return m.styles.ErrorMsg.Render("  ❌ No quiz questions found")
	}

	if m.currentQuizQ >= len(step.Questions) {
		m.currentQuizQ = 0
	}

	q := step.Questions[m.currentQuizQ]

	// Title
	title := m.styles.StepTitle.Width(width).Render(
		fmt.Sprintf("  Quiz - Question %d/%d  ", m.currentQuizQ+1, len(step.Questions)),
	)
	b.WriteString("\n" + title + "\n\n")

	// Question
	question := m.styles.Bold.Render(q.Question)
	b.WriteString("  " + question + "\n\n")

	// Options
	for i, option := range q.Options {
		var optionLine string

		if i == m.selectedOption {
			// Highlighted option
			optionStyle := lipgloss.NewStyle().
				Foreground(m.styles.Theme.Primary).
				Bold(true).
				Background(lipgloss.Color("240"))

			optionLine = optionStyle.Render(fmt.Sprintf("  ▸ %s  ", option))
		} else {
			optionLine = fmt.Sprintf("    %s", option)
		}

		b.WriteString(optionLine + "\n")
	}

	b.WriteString("\n")

	// Help text
	helpText := m.styles.Muted.Render("  ↑/↓ or j/k: Navigate  •  Enter: Submit answer")
	b.WriteString(helpText + "\n")

	// Show result if already answered
	if answered, ok := m.quizAnswers[q.ID]; ok {
		b.WriteString("\n")
		if answered == q.Answer {
			result := m.styles.SuccessMsg.Render("  ✅ Correct!")
			b.WriteString(result + "\n")
		} else {
			result := m.styles.ErrorMsg.Render("  ❌ Incorrect")
			b.WriteString(result + "\n")
			correctAnswer := m.styles.Paragraph.Render(
				"  Correct answer: " + q.Options[q.Answer],
			)
			b.WriteString(correctAnswer + "\n")
		}

		if q.Explanation != "" {
			b.WriteString("\n")
			explanationBox := m.styles.CalloutTip.Width(width - 4).Render(
				"💡 Explanation\n\n" + q.Explanation,
			)
			b.WriteString("  " + explanationBox + "\n")
		}

		continueHint := m.styles.Muted.Render("\n  Press Enter to continue...")
		b.WriteString(continueHint + "\n")
	}

	return b.String()
}

// renderChallengeStep renders a challenge step
func (m Model) renderChallengeStep(step lessons_pkg.StepType) string {
	var b strings.Builder

	width := m.getContentWidth()

	title := m.styles.StepTitle.Width(width).Render(step.Title)
	b.WriteString(title + "\n\n")

	desc := m.formatMarkdown(step.Description)
	b.WriteString(desc + "\n\n")

	// Show hints if available
	if len(step.Hints) > 0 && m.currentHint > 0 {
		hint := step.Hints[min(m.currentHint-1, len(step.Hints)-1)]
		hintBox := m.styles.HintBox.Width(width).Render(
			fmt.Sprintf("💡 Hint: %s", hint.Text),
		)
		b.WriteString(hintBox + "\n\n")
	}

	continueHint := m.styles.HelpText.Render("Press Enter when ready to continue...")
	b.WriteString(continueHint + "\n")

	return b.String()
}

// renderInterviewStep renders an interview prep step
func (m Model) renderInterviewStep(step lessons_pkg.StepType) string {
	var b strings.Builder

	width := m.getContentWidth()

	title := m.styles.StepTitle.Width(width).Render(step.Title)
	b.WriteString(title + "\n\n")

	if step.Description != "" {
		b.WriteString(m.formatMarkdown(step.Description) + "\n\n")
	}

	if len(step.Questions) > 0 {
		for i, q := range step.Questions {
			questionText := fmt.Sprintf("%d. %s", i+1, q.Question)
			b.WriteString(m.styles.Paragraph.Render(questionText) + "\n\n")
		}
	}

	note := m.styles.CalloutTip.Width(width).Render(
		"💡 Take time to think about these questions. Your understanding will be tested in real scenarios!",
	)
	b.WriteString(note + "\n\n")

	continueHint := m.styles.HelpText.Render("Press Enter to continue...")
	b.WriteString(continueHint + "\n")

	return b.String()
}

// renderInfoStep renders an info step
// renderInfoStep with padding
func (m Model) renderInfoStep(step lessons_pkg.StepType) string {
	var b strings.Builder

	width := m.getContentWidth()

	// Title with padding
	title := m.styles.StepTitle.Width(width).Render("  " + step.Title + "  ")
	b.WriteString("\n" + title + "\n\n")

	// Content with proper formatting
	content := m.formatMarkdown(step.Content)
	// Add left padding to each line
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		b.WriteString("  " + line + "\n")
	}

	// Code blocks
	for _, block := range step.CodeBlocks {
		b.WriteString("\n" + m.renderCodeBlock(block) + "\n")
	}

	// Callouts
	for _, callout := range step.Callouts {
		b.WriteString("\n  " + m.renderCallout(callout) + "\n")
	}

	// Diagram
	if step.Diagram != "" {
		b.WriteString("\n" + m.renderDiagram(step.Diagram) + "\n")
	}

	// Continue hint
	continueHint := m.styles.Muted.Render("\n  Press Enter to continue...")
	b.WriteString(continueHint + "\n")

	return b.String()
}

// renderConclusion renders the conclusion screen
func (m Model) renderConclusion() string {
	var b strings.Builder

	width := m.getContentWidth()

	header := m.styles.AppTitle.Width(width).Render("🎉 Lesson Complete!")
	b.WriteString(header + "\n\n")

	conclusionBox := m.styles.Border.Width(width).Render(
		m.styles.Bold.Render(m.lesson.Conclusion.Title) + "\n\n" +
			m.formatMarkdown(m.lesson.Conclusion.Content),
	)
	b.WriteString(conclusionBox + "\n\n")

	if len(m.lesson.Conclusion.Badges) > 0 {
		b.WriteString(m.styles.Bold.Render("🏆 Badges Earned:") + "\n\n")
		for _, badge := range m.lesson.Conclusion.Badges {
			b.WriteString(fmt.Sprintf("  %s %s\n", badge.Icon, badge.Name))
		}
		b.WriteString("\n")
	}

	exitHint := m.styles.HelpText.Render("Press Enter to exit...")
	b.WriteString(exitHint + "\n")

	return b.String()
}

// Helper rendering functions

func (m Model) renderProgressBar() string {
	progress := float64(m.currentStep) / float64(m.totalSteps)
	if m.state == StateConclusion {
		progress = 1.0
	}

	filled := int(progress * 20)
	empty := 20 - filled

	filledStyle := lipgloss.NewStyle().Foreground(m.styles.Theme.Success)

	bar := "Progress: [" +
		filledStyle.Render(strings.Repeat("█", filled)) +
		strings.Repeat("░", empty) + "] " +
		fmt.Sprintf("%d%%", int(progress*100))

	return bar
}

func (m Model) renderCodeBlock(block lessons_pkg.CodeBlockType) string {
	width := m.getContentWidth()

	header := m.styles.Bold.Render(block.Label)
	code := block.Code

	return m.styles.CodeBlock.Width(width).Render(header + "\n\n" + code)
}

func (m Model) renderCallout(callout lessons_pkg.CalloutType) string {
	width := m.getContentWidth()

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

	return style.Width(width).Render(fmt.Sprintf("%s %s", icon, callout.Text))
}

func (m Model) renderDiagram(diagram string) string {
	width := m.getContentWidth()
	return m.styles.Border.Width(width).Render(diagram)
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
	var help string

	switch stepType {
	case "command":
		help = "Type to enter command  •  Enter: Execute  •  ?: Hint  •  Ctrl+Y: Skip  •  Esc: Back  •  Ctrl+C: Quit"
	case "quiz":
		help = "↑↓/jk: Select option  •  Enter: Submit  •  Esc: Back  •  Ctrl+C: Quit"
	default:
		help = "Enter: Continue  •  ↑↓/jk: Scroll  •  Esc: Back  •  Ctrl+C: Quit"
	}

	return m.styles.Muted.Render(help)
}

// formatMarkdown applies basic markdown formatting
func (m Model) formatMarkdown(text string) string {
	lines := strings.Split(text, "\n")
	var formatted []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Handle different markdown elements
		if strings.HasPrefix(trimmed, "###") {
			// Subheading
			heading := strings.TrimPrefix(trimmed, "###")
			heading = strings.TrimSpace(heading)
			formatted = append(formatted, "")
			formatted = append(formatted, m.styles.Bold.Render(heading))
			formatted = append(formatted, "")
		} else if strings.HasPrefix(trimmed, "##") {
			// Heading
			heading := strings.TrimPrefix(trimmed, "##")
			heading = strings.TrimSpace(heading)
			formatted = append(formatted, "")
			formatted = append(formatted, m.styles.Bold.Render(heading))
			formatted = append(formatted, "")
		} else if strings.HasPrefix(trimmed, "- ") || strings.HasPrefix(trimmed, "* ") {
			// Bullet point - maintain indentation
			bullet := strings.TrimPrefix(trimmed, "- ")
			bullet = strings.TrimPrefix(bullet, "* ")
			formatted = append(formatted, "  • "+m.formatInlineMarkdown(bullet))
		} else if regexp.MustCompile(`^\d+\.`).MatchString(trimmed) {
			// Numbered list - maintain indentation
			formatted = append(formatted, "  "+m.formatInlineMarkdown(trimmed))
		} else if strings.HasPrefix(trimmed, ">") {
			// Blockquote
			quote := strings.TrimPrefix(trimmed, ">")
			quote = strings.TrimSpace(quote)
			formatted = append(formatted, "  "+m.styles.Muted.Render("│ "+quote))
		} else if trimmed == "" {
			// Empty line
			formatted = append(formatted, "")
		} else {
			// Regular paragraph - add left padding
			formatted = append(formatted, m.formatInlineMarkdown(trimmed))
		}
	}

	return strings.Join(formatted, "\n")
}

// Helper for inline markdown (bold, code)
func (m Model) formatInlineMarkdown(text string) string {
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

	return text
}

func (m Model) getContentWidth() int {
	maxWidth := 120 // Increased for better readability
	padding := 20

	if m.width < maxWidth+padding {
		return m.width - padding
	}
	return maxWidth
}

// Helper min function
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
