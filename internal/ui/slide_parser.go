package ui

import (
	"strings"
)

// Slide represents a single presentation slide parsed from markdown
type Slide struct {
	// Raw markdown content of this slide
	Content string
	// Inferred title (first heading found, or empty)
	Title string
	// Index (0-based)
	Index int
}

// ParseSlides splits a markdown document into individual slides.
// Slides are separated by lines that contain only "---".
func ParseSlides(markdown string) []Slide {
	// Normalize line endings
	markdown = strings.ReplaceAll(markdown, "\r\n", "\n")

	// Split on horizontal rule separators ("---" on its own line)
	rawSlides := splitOnSeparator(markdown)

	var slides []Slide
	for _, raw := range rawSlides {
		trimmed := strings.TrimSpace(raw)
		if trimmed == "" {
			continue
		}
		slides = append(slides, Slide{
			Content: trimmed,
			Title:   extractTitle(trimmed),
			Index:   len(slides),
		})
	}
	return slides
}

// splitOnSeparator splits markdown text on lines that are exactly "---".
// It avoids splitting on YAML front-matter (the very first "---" block).
func splitOnSeparator(markdown string) []string {
	lines := strings.Split(markdown, "\n")

	var parts []string
	var current strings.Builder

	for _, line := range lines {
		if strings.TrimSpace(line) == "---" {
			parts = append(parts, current.String())
			current.Reset()
		} else {
			current.WriteString(line)
			current.WriteByte('\n')
		}
	}

	// Flush remaining content
	if current.Len() > 0 {
		parts = append(parts, current.String())
	}

	return parts
}

// extractTitle returns the first heading found in the slide content.
func extractTitle(content string) string {
	for _, line := range strings.Split(content, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "# ") {
			return strings.TrimPrefix(trimmed, "# ")
		}
		if strings.HasPrefix(trimmed, "## ") {
			return strings.TrimPrefix(trimmed, "## ")
		}
		if strings.HasPrefix(trimmed, "### ") {
			return strings.TrimPrefix(trimmed, "### ")
		}
	}
	return ""
}
