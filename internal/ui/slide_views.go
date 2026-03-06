package ui

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// numberedListRe matches lines beginning with a number followed by ". "
var numberedListRe = regexp.MustCompile(`^\d+\.\s`)

// renderCurrentSlide formats the current slide's markdown for display.
func (m SlideModel) renderCurrentSlide() string {
	if m.total == 0 || m.current >= m.total {
		return ""
	}

	slide := m.slides[m.current]

	// Use a lightweight stand-alone formatter so we don't need the full
	// lessons Model.  We re-use the same style system already in the package.
	sf := &slideFormatter{
		styles:       m.styles,
		contentWidth: m.width - 8, // leave breathing room on both sides
	}
	if sf.contentWidth < 40 {
		sf.contentWidth = 40
	}

	return sf.format(slide.Content)
}

// ──────────────────────────────────────────────
// Stand-alone slide formatter
// ──────────────────────────────────────────────

type slideFormatter struct {
	styles       *Styles
	contentWidth int
}

// format converts markdown text to a styled string ready for display.
func (sf *slideFormatter) format(markdown string) string {
	lines := strings.Split(markdown, "\n")
	var out []string

	inCodeBlock := false
	codeBlockLang := ""
	var codeLines []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// ── Code block toggle ─────────────────────────────────────────────
		if strings.HasPrefix(trimmed, "```") {
			if !inCodeBlock {
				inCodeBlock = true
				codeBlockLang = strings.TrimPrefix(trimmed, "```")
				if codeBlockLang == "" {
					codeBlockLang = "text"
				}
				codeLines = nil
			} else {
				// End of code block – render it
				code := strings.Join(codeLines, "\n")
				out = append(out, "")
				out = append(out, RenderCodeBlock(code, codeBlockLang, "", sf.contentWidth))
				out = append(out, "")
				inCodeBlock = false
				codeLines = nil
			}
			continue
		}

		if inCodeBlock {
			codeLines = append(codeLines, line)
			continue
		}

		// ── Headings ──────────────────────────────────────────────────────
		if strings.HasPrefix(trimmed, "# ") {
			heading := strings.TrimPrefix(trimmed, "# ")
			h1Style := lipgloss.NewStyle().
				Bold(true).
				Foreground(sf.styles.Theme.Primary).
				Width(sf.contentWidth).
				Align(lipgloss.Center).
				MarginTop(1).
				MarginBottom(1)
			out = append(out, "")
			out = append(out, h1Style.Render(heading))
			out = append(out, lipgloss.NewStyle().Foreground(sf.styles.Theme.Primary).
				Render(strings.Repeat("═", sf.contentWidth)))
			out = append(out, "")
			continue
		}

		if strings.HasPrefix(trimmed, "## ") {
			heading := strings.TrimPrefix(trimmed, "## ")
			h2Style := lipgloss.NewStyle().
				Bold(true).
				Foreground(sf.styles.Theme.Primary).
				MarginTop(1)
			out = append(out, "")
			out = append(out, h2Style.Render("  "+heading))
			out = append(out, "")
			continue
		}

		if strings.HasPrefix(trimmed, "### ") {
			heading := strings.TrimPrefix(trimmed, "### ")
			h3Style := lipgloss.NewStyle().
				Bold(true).
				Foreground(sf.styles.Theme.Warning)
			out = append(out, "")
			out = append(out, h3Style.Render("    "+heading))
			out = append(out, "")
			continue
		}

		// ── Horizontal rules (already used as separators, skip extras) ────
		if trimmed == "---" || trimmed == "***" || trimmed == "___" {
			out = append(out, "")
			out = append(out, lipgloss.NewStyle().
				Foreground(lipgloss.Color("240")).
				Render(strings.Repeat("─", sf.contentWidth)))
			out = append(out, "")
			continue
		}

		// ── Blockquotes ───────────────────────────────────────────────────
		if strings.HasPrefix(trimmed, "> ") {
			quote := strings.TrimPrefix(trimmed, "> ")
			quoteStyle := lipgloss.NewStyle().
				Border(lipgloss.ThickBorder(), false, false, false, true).
				BorderForeground(sf.styles.Theme.Primary).
				Foreground(lipgloss.Color("252")).
				Padding(0, 2).
				Width(sf.contentWidth - 2)
			out = append(out, quoteStyle.Render(sf.inline(quote)))
			continue
		}

		// ── Bullet lists ──────────────────────────────────────────────────
		if strings.HasPrefix(trimmed, "- ") || strings.HasPrefix(trimmed, "* ") {
			bullet := strings.TrimPrefix(trimmed, "- ")
			bullet = strings.TrimPrefix(bullet, "* ")
			bulletStyle := lipgloss.NewStyle().
				Foreground(sf.styles.Theme.Primary).
				Bold(true)
			out = append(out, "  "+bulletStyle.Render("•")+" "+sf.inline(bullet))
			continue
		}

		// ── Numbered lists ────────────────────────────────────────────────
		if numberedListRe.MatchString(trimmed) {
			out = append(out, "  "+sf.inline(trimmed))
			continue
		}

		// ── Empty line ────────────────────────────────────────────────────
		if trimmed == "" {
			out = append(out, "")
			continue
		}

		// ── Regular paragraph ─────────────────────────────────────────────
		out = append(out, "  "+sf.inline(trimmed))
	}

	return strings.Join(out, "\n")
}

// inline handles inline markdown within a line.
func (sf *slideFormatter) inline(text string) string {
	// Bold must be processed before italic to avoid partial marker consumption.
	// Double markers (**/__) are processed before single (* /_).
	boldStyle := lipgloss.NewStyle().Bold(true).Foreground(sf.styles.Theme.Primary)
	italicStyle := lipgloss.NewStyle().Italic(true)

	// Bold (**text**)
	text = replaceEnclosed(text, "**", boldStyle)
	// Bold (__text__)
	text = replaceEnclosed(text, "__", boldStyle)

	// Italic (*text*) — only after bold so remaining single '*' are italic
	text = replaceEnclosed(text, "*", italicStyle)
	// Italic (_text_) — only after bold so remaining single '_' are italic
	text = replaceEnclosed(text, "_", italicStyle)

	// Inline code (`code`)
	text = replaceInlineCode(text, lipgloss.NewStyle().
		Foreground(lipgloss.Color("214")).
		Background(lipgloss.Color("235")).
		Padding(0, 1))

	// Links [text](url) – show as "text (url)"
	text = replaceLinks(text, sf.styles)

	return text
}

// replaceEnclosed replaces `marker text marker` with styled output.
// Only handles symmetric single-character or two-character markers
// on a single line.
func replaceEnclosed(text, marker string, style lipgloss.Style) string {
	for {
		start := strings.Index(text, marker)
		if start < 0 {
			break
		}
		end := strings.Index(text[start+len(marker):], marker)
		if end < 0 {
			break
		}
		end += start + len(marker) // absolute position of second marker

		inner := text[start+len(marker) : end]
		if strings.ContainsAny(inner, "\n\r") {
			break // don't span lines
		}
		styled := style.Render(inner)
		text = text[:start] + styled + text[end+len(marker):]
	}
	return text
}

// replaceInlineCode replaces `code` spans with styled text.
func replaceInlineCode(text string, style lipgloss.Style) string {
	for {
		start := strings.Index(text, "`")
		if start < 0 {
			break
		}
		end := strings.Index(text[start+1:], "`")
		if end < 0 {
			break
		}
		end += start + 1 // absolute position of closing backtick
		inner := text[start+1 : end]
		text = text[:start] + style.Render(inner) + text[end+1:]
	}
	return text
}

// replaceLinks replaces [text](url) with "text (url)" styled.
func replaceLinks(text string, styles *Styles) string {
	for {
		start := strings.Index(text, "[")
		if start < 0 {
			break
		}
		mid := strings.Index(text[start:], "](")
		if mid < 0 {
			break
		}
		mid += start
		end := strings.Index(text[mid+2:], ")")
		if end < 0 {
			break
		}
		end += mid + 2

		linkText := text[start+1 : mid]
		url := text[mid+2 : end]

		urlStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("39")).
			Underline(true)

		replacement := fmt.Sprintf("%s (%s)", styles.Bold.Render(linkText), urlStyle.Render(url))
		text = text[:start] + replacement + text[end+1:]
	}
	return text
}
