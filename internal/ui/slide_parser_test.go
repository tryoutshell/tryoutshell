package ui

import (
	"testing"
)

func TestParseSlides_BasicSeparator(t *testing.T) {
	input := "# Slide One\n\nHello\n\n---\n\n# Slide Two\n\nWorld"
	slides := ParseSlides(input)
	if len(slides) != 2 {
		t.Fatalf("expected 2 slides, got %d", len(slides))
	}
	if slides[0].Title != "Slide One" {
		t.Errorf("expected title 'Slide One', got %q", slides[0].Title)
	}
	if slides[1].Title != "Slide Two" {
		t.Errorf("expected title 'Slide Two', got %q", slides[1].Title)
	}
}

func TestParseSlides_EmptyInput(t *testing.T) {
	slides := ParseSlides("")
	if len(slides) != 0 {
		t.Errorf("expected 0 slides for empty input, got %d", len(slides))
	}
}

func TestParseSlides_NoSeparator(t *testing.T) {
	input := "# Single slide\n\nSome content here."
	slides := ParseSlides(input)
	if len(slides) != 1 {
		t.Fatalf("expected 1 slide, got %d", len(slides))
	}
	if slides[0].Title != "Single slide" {
		t.Errorf("expected title 'Single slide', got %q", slides[0].Title)
	}
}

func TestParseSlides_SkipsBlankSlides(t *testing.T) {
	// Leading/trailing separators produce blank slides that should be skipped
	input := "---\n\n# Real Slide\n\nContent\n\n---\n"
	slides := ParseSlides(input)
	if len(slides) != 1 {
		t.Fatalf("expected 1 slide (blank ones filtered), got %d", len(slides))
	}
}

func TestParseSlides_MultipleSeparators(t *testing.T) {
	input := "A\n---\nB\n---\nC\n---\nD"
	slides := ParseSlides(input)
	if len(slides) != 4 {
		t.Fatalf("expected 4 slides, got %d", len(slides))
	}
}

func TestExtractTitle_H1(t *testing.T) {
	title := extractTitle("# My Title\n\nSome text")
	if title != "My Title" {
		t.Errorf("expected 'My Title', got %q", title)
	}
}

func TestExtractTitle_H2(t *testing.T) {
	title := extractTitle("## Sub Title\n\nMore text")
	if title != "Sub Title" {
		t.Errorf("expected 'Sub Title', got %q", title)
	}
}

func TestExtractTitle_NoHeading(t *testing.T) {
	title := extractTitle("Just a paragraph without heading")
	if title != "" {
		t.Errorf("expected empty string, got %q", title)
	}
}

func TestSlideModel_FindSlide(t *testing.T) {
	slides := []Slide{
		{Content: "hello world", Index: 0},
		{Content: "foo bar", Index: 1},
		{Content: "Hello Again", Index: 2},
	}
	m := NewSlideModel(slides)

	// Should find first occurrence
	idx := m.findSlide("hello", 0)
	if idx != 0 {
		t.Errorf("expected index 0, got %d", idx)
	}

	// Should find starting from offset (wraps around)
	idx = m.findSlide("hello", 1)
	if idx != 2 {
		t.Errorf("expected index 2, got %d", idx)
	}

	// Not found returns -1
	idx = m.findSlide("notfound", 0)
	if idx != -1 {
		t.Errorf("expected -1, got %d", idx)
	}
}

func TestParseNum(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"5", 5},
		{"10", 10},
		{"1", 1},
		{"", 1},
		{"abc", 1},
		{"0", 1}, // 0 is treated as 1
	}
	for _, c := range cases {
		got := parseNum(c.input)
		if got != c.expected {
			t.Errorf("parseNum(%q) = %d, want %d", c.input, got, c.expected)
		}
	}
}
