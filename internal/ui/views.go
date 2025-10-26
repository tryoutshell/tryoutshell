package ui

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
	lessons_pkg "github.com/tryoutshell/tryoutshell/internal/lessons"
	"github.com/tryoutshell/tryoutshell/internal/runner"
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

	// Step title bar (like codecrafters)
	titleBar := m.styles.StepTitle.
		Width(width).
		Render(step.Prompt)
	b.WriteString(titleBar + "\n\n")

	// Working directory (subtle)
	workingDirInfo := m.styles.Muted.Render(
		fmt.Sprintf("📂 %s", m.runner.GetWorkingDir()),
	)
	b.WriteString(workingDirInfo + "\n\n")

	// Pre-content with proper formatting
	if step.PreContent != "" {
		preContent := m.formatMarkdown(step.PreContent)
		b.WriteString(preContent + "\n\n")
	}

	// Instruction (arrow pointer like codecrafters)
	if step.Instruction != "" {
		instruction := m.styles.InfoMsg.Render("→ " + step.Instruction)
		b.WriteString(instruction + "\n\n")
	}

	// Example command in clean box
	if step.Example != "" {
		exampleLabel := m.styles.Muted.Render("Example:")
		b.WriteString(exampleLabel + "\n\n")

		exampleContent := m.styles.CommandPrompt.Render() + step.Example
		exampleBox := m.styles.CommandExample.
			Width(width).
			Render(exampleContent)
		b.WriteString(exampleBox + "\n")
	}

	// Input area or output
	if m.stepState == StepPending || m.stepState == StepFailed {
		// Input label
		inputLabel := m.styles.InfoMsg.Render("Your turn:")
		b.WriteString(inputLabel + "\n\n")

		// Input box with proper width
		inputContent := m.styles.CommandPrompt.Render() + m.textInput.View()
		inputBox := m.styles.CommandInput.
			Width(width).
			Render(inputContent)
		b.WriteString(inputBox + "\n")

		// Debug commands hint (subtle, like codecrafters footer)
		debugHint := m.styles.Muted.Render(
			"💡 Debug commands: :pwd :ls :state | Type ':skip' or Ctrl+Y to skip | '?' for hints",
		)
		b.WriteString(debugHint + "\n")

		// Show hints if requested
		if m.currentHint > 0 && m.currentHint <= len(step.Hints) {
			hint := step.Hints[m.currentHint-1]
			hintContent := fmt.Sprintf("💡 Hint %d/%d\n\n%s",
				m.currentHint, len(step.Hints), hint.Text)
			hintBox := m.styles.HintBox.
				Width(width).
				Render(hintContent)
			b.WriteString("\n" + hintBox + "\n")
		}

		// Error details if failed
		if m.stepState == StepFailed {
			b.WriteString("\n" + m.renderFullErrorDetails(step, m.lastCommandResult) + "\n")
		}

	} else if m.stepState == StepExecuting {
		executing := m.styles.WarningMsg.Render("⏳ Executing command...")
		b.WriteString(executing + "\n")

	} else if m.stepState == StepSuccess {
		// Command executed box
		cmdLabel := m.styles.Muted.Render("Command executed:")
		b.WriteString(cmdLabel + "\n\n")

		cmdContent := m.styles.CommandPrompt.Render() + m.textInput.Value()
		cmdBox := m.styles.BoxBorder.
			Width(width).
			Render(cmdContent)
		b.WriteString(cmdBox + "\n\n")

		// Output box
		if m.commandOutput != "" {
			outputLabel := m.styles.Muted.Render("Output:")
			b.WriteString(outputLabel + "\n\n")

			outputBox := m.styles.OutputSuccess.
				Width(width).
				Render(m.commandOutput)
			b.WriteString(outputBox + "\n\n")
		}

		// Success message
		successMsg := m.styles.SuccessMsg.Render("✅ " + step.SuccessMsg)
		b.WriteString(successMsg + "\n")

		// Post-content
		if step.PostContent != "" {
			b.WriteString("\n" + m.formatMarkdown(step.PostContent) + "\n")
		}

		// Continue hint
		continueHint := m.styles.Muted.Render("\nPress Enter to continue...")
		b.WriteString(continueHint + "\n")
	}

	return b.String()
}

// Render full error details with debugging context
func (m Model) renderFullErrorDetails(step lessons_pkg.StepType, result runner.CommandResult) string {
	var b strings.Builder

	width := m.getContentWidth()

	// Error header
	errorHeader := m.styles.ErrorMsg.Render("  ❌ " + step.FailMsg)
	b.WriteString(errorHeader + "\n\n")

	// Debug information box
	debugBox := strings.Builder{}
	debugBox.WriteString("📋 DEBUG INFORMATION\n")
	debugBox.WriteString(strings.Repeat("━", width-6) + "\n\n")

	// Command info
	debugBox.WriteString("Command executed:\n")
	debugBox.WriteString("  $ " + result.Command + "\n\n")

	// Working directory
	debugBox.WriteString("Working directory:\n")
	debugBox.WriteString("  " + result.WorkingDir + "\n\n")

	// Exit code
	debugBox.WriteString("Exit code: ")
	if result.ExitCode == 0 {
		debugBox.WriteString("0 (success)\n\n")
	} else {
		debugBox.WriteString(fmt.Sprintf("%d (error)\n\n", result.ExitCode))
	}

	// Duration
	debugBox.WriteString(fmt.Sprintf("Duration: %.2fs\n\n", result.Duration.Seconds()))

	// Full output
	if result.Output != "" {
		debugBox.WriteString("Full output:\n")
		outputLines := strings.Split(result.Output, "\n")
		for _, line := range outputLines {
			debugBox.WriteString("  " + line + "\n")
		}
		debugBox.WriteString("\n")
	}

	// Validation details
	debugBox.WriteString("Validation attempted:\n")
	debugBox.WriteString(fmt.Sprintf("  Type: %s\n", result.ValidationInfo.ValidationType))

	if !result.ValidationInfo.Passed {
		debugBox.WriteString("  Status: ✗ FAILED\n\n")

		debugBox.WriteString(fmt.Sprintf("  Expected: %v\n", result.ValidationInfo.Expected))
		debugBox.WriteString(fmt.Sprintf("  Actual: %v\n\n", result.ValidationInfo.Actual))

		if len(result.ValidationInfo.Details) > 0 {
			debugBox.WriteString("  Details:\n")
			for _, detail := range result.ValidationInfo.Details {
				debugBox.WriteString("    • " + detail + "\n")
			}
			debugBox.WriteString("\n")
		}
	}

	// Troubleshooting suggestions
	debugBox.WriteString("What to check:\n")
	debugBox.WriteString("  1. Run ':pwd' to see your current directory\n")
	debugBox.WriteString("  2. Run ':ls' to list files in the directory\n")
	debugBox.WriteString("  3. Run ':state' to see expected lesson state\n")

	// Specific suggestions based on validation type
	if result.ValidationInfo.ValidationType == "file_exists" {
		debugBox.WriteString("  4. Check file permissions: ls -la\n")
		debugBox.WriteString("  5. Verify you're in a writable directory: touch test.txt\n")
	}

	debugBox.WriteString("\n" + strings.Repeat("━", width-6))

	// Render the debug box
	debugBoxStyled := m.styles.OutputError.Width(width - 4).Render(debugBox.String())
	b.WriteString("  " + debugBoxStyled + "\n\n")

	// Action prompt
	actionPrompt := m.styles.Muted.Render("  Press Enter to try again, '?' for hints, or Ctrl+Y to skip")
	b.WriteString(actionPrompt + "\n")

	return b.String()
}

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

	// Title with progress
	title := m.styles.StepTitle.Width(width).Render(
		fmt.Sprintf("  📝 Quiz - Question %d of %d  ", m.currentQuizQ+1, len(step.Questions)),
	)
	b.WriteString("\n" + title + "\n\n")

	// Question with nice formatting
	questionBox := m.styles.Border.Width(width - 4).Render(
		m.styles.Bold.Render("❓ " + q.Question),
	)
	b.WriteString("  " + questionBox + "\n\n")

	// Check if already answered
	answered, hasAnswered := m.quizAnswers[q.ID]

	// Options with better styling
	for i, option := range q.Options {
		var optionLine string
		var optionPrefix string

		if hasAnswered {
			// Show which option was selected and which was correct
			if i == answered && i == q.Answer {
				// User selected correct answer
				optionPrefix = "  ✅ "
				optionStyle := lipgloss.NewStyle().
					Foreground(m.styles.Theme.Success).
					Bold(true)
				optionLine = optionStyle.Render(optionPrefix + option)

			} else if i == answered && i != q.Answer {
				// User selected wrong answer
				optionPrefix = "  ❌ "
				optionStyle := lipgloss.NewStyle().
					Foreground(m.styles.Theme.Error).
					Bold(true).
					Strikethrough(true)
				optionLine = optionStyle.Render(optionPrefix + option)

			} else if i == q.Answer {
				// Show correct answer
				optionPrefix = "  ✓  "
				optionStyle := lipgloss.NewStyle().
					Foreground(m.styles.Theme.Success)
				optionLine = optionStyle.Render(optionPrefix + option + " (correct)")

			} else {
				// Other options
				optionPrefix = "     "
				optionLine = m.styles.Muted.Render(optionPrefix + option)
			}

		} else {
			// Not answered yet - show selection cursor
			if i == m.selectedOption {
				// Highlighted option
				optionStyle := lipgloss.NewStyle().
					Foreground(m.styles.Theme.Primary).
					Bold(true).
					Background(lipgloss.Color("237"))

				optionLine = "  " + optionStyle.Render(fmt.Sprintf(" ▸ %s ", option))
			} else {
				optionLine = fmt.Sprintf("     %s", option)
			}
		}

		b.WriteString(optionLine + "\n")
	}

	b.WriteString("\n")

	// Show result if already answered
	if hasAnswered {
		if answered == q.Answer {
			result := m.styles.SuccessMsg.Render("  🎉 Correct! Well done!")
			b.WriteString(result + "\n\n")
		} else {
			result := m.styles.ErrorMsg.Render("  ❌ Incorrect")
			b.WriteString(result + "\n\n")

			correctAnswer := lipgloss.NewStyle().
				Foreground(m.styles.Theme.Success).
				Render(fmt.Sprintf("  The correct answer is: %s", q.Options[q.Answer]))
			b.WriteString(correctAnswer + "\n\n")
		}

		// Explanation
		if q.Explanation != "" {
			explanationBox := m.styles.CalloutTip.Width(width - 6).Render(
				"💡 Explanation\n\n" + q.Explanation,
			)
			b.WriteString("  " + explanationBox + "\n\n")
		}

		// Show score
		correctCount := 0
		for qID, ans := range m.quizAnswers {
			for _, question := range step.Questions {
				if question.ID == qID && ans == question.Answer {
					correctCount++
				}
			}
		}
		score := lipgloss.NewStyle().
			Foreground(m.styles.Theme.Primary).
			Render(fmt.Sprintf("  📊 Score: %d/%d correct", correctCount, len(m.quizAnswers)))
		b.WriteString(score + "\n\n")

		continueHint := m.styles.Muted.Render("  Press Enter to continue...")
		b.WriteString(continueHint + "\n")

	} else {
		// Help text for navigation
		helpBox := m.styles.HintBox.Width(width - 6).Render(
			"⌨️  Use ↑↓ or j/k to navigate\n" +
				"   Press Enter to submit your answer",
		)
		b.WriteString("  " + helpBox + "\n")
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

	// Title bar
	titleBar := m.styles.StepTitle.
		Width(width).
		Render(step.Title)
	b.WriteString(titleBar + "\n\n")

	// Content with proper formatting
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

	// Continue hint
	continueHint := m.styles.Muted.Render("\nPress Enter to continue...")
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

// renderDiagram renders ASCII diagrams with syntax highlighting
func (m Model) renderDiagram(diagram string) string {
	lines := strings.Split(diagram, "\n")
	var highlighted []string

	for _, line := range lines {
		// Highlight box drawing characters
		if strings.Contains(line, "┌") || strings.Contains(line, "└") ||
			strings.Contains(line, "│") || strings.Contains(line, "─") ||
			strings.Contains(line, "┐") || strings.Contains(line, "┘") {
			line = lipgloss.NewStyle().
				Foreground(m.styles.Theme.Primary).
				Render(line)
		}

		// Highlight arrows
		if strings.Contains(line, "→") || strings.Contains(line, "▶") {
			line = lipgloss.NewStyle().
				Foreground(m.styles.Theme.Success).
				Bold(true).
				Render(line)
		}

		// Highlight keywords
		line = strings.ReplaceAll(line, "Image",
			lipgloss.NewStyle().Foreground(lipgloss.Color("214")).Bold(true).Render("Image"))
		line = strings.ReplaceAll(line, "Cosign",
			lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true).Render("Cosign"))
		line = strings.ReplaceAll(line, ".sig",
			lipgloss.NewStyle().Foreground(lipgloss.Color("46")).Render(".sig"))
		line = strings.ReplaceAll(line, "✓",
			lipgloss.NewStyle().Foreground(m.styles.Theme.Success).Bold(true).Render("✓"))

		highlighted = append(highlighted, line)
	}

	diagramBox := m.styles.BoxBorder.
		Width(m.getContentWidth()).
		Render(strings.Join(highlighted, "\n"))

	return diagramBox
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
	inCodeBlock := false

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Code blocks (```)
		if strings.HasPrefix(trimmed, "```") {
			inCodeBlock = !inCodeBlock
			continue
		}

		if inCodeBlock {
			// Render code block line
			codeLine := lipgloss.NewStyle().
				Foreground(lipgloss.Color("252")).
				Render("  " + line)
			formatted = append(formatted, codeLine)
			continue
		}

		// Headings
		if strings.HasPrefix(trimmed, "### ") {
			heading := strings.TrimPrefix(trimmed, "### ")
			formatted = append(formatted, "")
			formatted = append(formatted, m.styles.SubHeading.Render(heading))
			formatted = append(formatted, "")
			continue
		}

		if strings.HasPrefix(trimmed, "## ") {
			heading := strings.TrimPrefix(trimmed, "## ")
			formatted = append(formatted, "")
			formatted = append(formatted, m.styles.Bold.Render(heading))
			formatted = append(formatted, "")
			continue
		}

		// Bullet points
		if strings.HasPrefix(trimmed, "- ") || strings.HasPrefix(trimmed, "* ") {
			bullet := strings.TrimPrefix(trimmed, "- ")
			bullet = strings.TrimPrefix(bullet, "* ")
			formatted = append(formatted, "  • "+m.formatInline(bullet))
			continue
		}

		// Numbered lists
		if regexp.MustCompile(`^\d+\.`).MatchString(trimmed) {
			formatted = append(formatted, "  "+m.formatInline(trimmed))
			continue
		}

		// Blockquotes
		if strings.HasPrefix(trimmed, "> ") {
			quote := strings.TrimPrefix(trimmed, "> ")
			quoteLine := m.styles.Muted.Render("│ " + quote)
			formatted = append(formatted, "  "+quoteLine)
			continue
		}

		// Empty lines
		if trimmed == "" {
			formatted = append(formatted, "")
			continue
		}

		// Regular paragraphs
		formatted = append(formatted, "  "+m.formatInline(line))

		// Add spacing after paragraphs (if next line is not empty)
		if i < len(lines)-1 && strings.TrimSpace(lines[i+1]) != "" &&
			!strings.HasPrefix(strings.TrimSpace(lines[i+1]), "-") {
			// Don't add extra space
		}
	}

	return strings.Join(formatted, "\n")
}

// formatInline handles inline markdown (bold, code, links)
func (m Model) formatInline(text string) string {
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

	// Links ([text](url))
	re = regexp.MustCompile(`\[([^\]]+)\]\(([^\)]+)\)`)
	text = re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) == 3 {
			return m.styles.Link.Render(parts[1])
		}
		return match
	})

	return text
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

// func (m Model) getContentWidth() int {
// 	maxWidth := 120 // Increased for better readability
// 	padding := 20

// 	if m.width < maxWidth+padding {
// 		return m.width - padding
// 	}
// 	return maxWidth
// }

func (m Model) getBoxWidth() int {
	// Account for: border (2 chars) + padding (4 chars) = 6 chars total
	return m.getContentWidth() - 6
}

func (m Model) getContentWidth() int {
	if m.width == 0 {
		return 100 // Default width
	}

	// Leave margins for borders and padding
	maxWidth := 120
	margins := 8 // Account for borders, padding

	availableWidth := m.width - margins

	if availableWidth < maxWidth {
		return availableWidth
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
