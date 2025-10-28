package ui

import (
	"bytes"
	"strings"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/charmbracelet/lipgloss"
)

// SyntaxHighlight applies syntax highlighting to code
func SyntaxHighlight(code, language string) string {
	// Get lexer for language
	lexer := lexers.Get(language)
	if lexer == nil {
		lexer = lexers.Fallback
	}
	lexer = chroma.Coalesce(lexer)

	// Use dracula style (dark theme)
	style := styles.Get("dracula")
	if style == nil {
		style = styles.Fallback
	}

	// Create terminal formatter with 16-color support
	formatter := formatters.Get("terminal16m")
	if formatter == nil {
		formatter = formatters.Fallback
	}

	// Tokenize and format
	iterator, err := lexer.Tokenise(nil, code)
	if err != nil {
		return code // Fallback to plain text
	}

	var buf bytes.Buffer
	err = formatter.Format(&buf, style, iterator)
	if err != nil {
		return code
	}

	return buf.String()
}

// RenderCodeBlock renders a code block with syntax highlighting and border
func RenderCodeBlock(code, language, label string, width int) string {
	// Syntax highlight
	highlighted := SyntaxHighlight(code, language)

	// Add language badge
	langBadge := lipgloss.NewStyle().
		Background(lipgloss.Color("39")).
		Foreground(lipgloss.Color("0")).
		Padding(0, 1).
		Bold(true).
		Render(strings.ToUpper(language))

	// Add label if provided
	header := ""
	if label != "" {
		labelStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Bold(true)
		header = labelStyle.Render(label) + " " + langBadge + "\n\n"
	} else {
		header = langBadge + "\n\n"
	}

	// Wrap in box
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		Padding(1, 2).
		Width(width)

	return boxStyle.Render(header + highlighted)
}
