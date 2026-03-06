package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tryoutshell/tryoutshell/types"
)

// ──────────────────────────────────────────────
// Key bindings
// ──────────────────────────────────────────────

type slideKeyMap struct {
	Next     key.Binding
	Prev     key.Binding
	First    key.Binding
	Last     key.Binding
	Search   key.Binding
	Quit     key.Binding
	ScrollUp key.Binding
	ScrollDn key.Binding
}

var slideKeys = slideKeyMap{
	Next: key.NewBinding(
		key.WithKeys("space", "right", "down", "enter", "n", "j", "l", "pgdown"),
		key.WithHelp("space/→/↓/enter", "next slide"),
	),
	Prev: key.NewBinding(
		key.WithKeys("left", "up", "p", "h", "k", "N", "pgup"),
		key.WithHelp("←/↑/p/h", "prev slide"),
	),
	First: key.NewBinding(
		key.WithKeys("g"),
		key.WithHelp("gg", "first slide"),
	),
	Last: key.NewBinding(
		key.WithKeys("G"),
		key.WithHelp("G", "last slide"),
	),
	Search: key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "search"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q/ctrl+c", "quit"),
	),
	ScrollUp: key.NewBinding(
		key.WithKeys("ctrl+u"),
		key.WithHelp("ctrl+u", "scroll up"),
	),
	ScrollDn: key.NewBinding(
		key.WithKeys("ctrl+d"),
		key.WithHelp("ctrl+d", "scroll down"),
	),
}

// ──────────────────────────────────────────────
// Presentation mode state
// ──────────────────────────────────────────────

type slideMode int

const (
	slideModePresenting slideMode = iota
	slideModeSearch
)

// ──────────────────────────────────────────────
// SlideModel is the Bubble Tea model for a slide presentation
// ──────────────────────────────────────────────

// SlideModel is the Bubble Tea model for presenting slides.
type SlideModel struct {
	slides       []Slide
	current      int
	total        int
	width        int
	height       int
	ready        bool
	viewport     viewport.Model
	mode         slideMode
	searchInput  textinput.Model
	searchQuery  string
	searchResult int // index of matching slide (-1 = none)
	gPressed     bool
	numBuf       string
	styles       *Styles

	lessonTitle string
	orgID       string
	lessonID    string
	quiz        []types.QuizQuestion
	showHelp    bool
}

// NewSlideModel creates a new SlideModel from a slice of parsed slides.
func NewSlideModel(slides []Slide) SlideModel {
	ti := textinput.New()
	ti.Placeholder = "search..."
	ti.CharLimit = 128

	theme := GetTheme("default")
	styles := NewStyles(theme)

	return SlideModel{
		slides:      slides,
		current:     0,
		total:       len(slides),
		searchInput: ti,
		styles:      styles,
	}
}

// ──────────────────────────────────────────────
// Bubble Tea interface
// ──────────────────────────────────────────────

// Init is called once when the program starts.
func (m SlideModel) Init() tea.Cmd {
	return nil
}

// Update handles incoming messages.
func (m SlideModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		headerH := 1
		footerH := 2
		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-headerH-footerH)
			m.viewport.YPosition = headerH
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - headerH - footerH
		}
		m.refreshContent()
		return m, nil

	case tea.KeyMsg:
		// Search mode handles its own keys
		if m.mode == slideModeSearch {
			return m.updateSearch(msg)
		}
		return m.updatePresent(msg)
	}

	// Forward remaining messages to the viewport when not in search mode
	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

// updatePresent handles keys in normal (presenting) mode.
func (m SlideModel) updatePresent(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	k := msg.String()

	// Number prefix accumulation (for "5G" etc.)
	if k >= "1" && k <= "9" {
		m.numBuf += k
		m.gPressed = false
		return m, nil
	}
	if k == "0" && m.numBuf != "" {
		m.numBuf += k
		m.gPressed = false
		return m, nil
	}

	switch {
	case key.Matches(msg, slideKeys.Quit):
		return m, tea.Quit

	case k == "?":
		m.showHelp = !m.showHelp
		return m, nil

	case k == "g":
		if m.gPressed {
			// "gg" – go to first slide
			m.current = 0
			m.gPressed = false
			m.numBuf = ""
			m.refreshContent()
		} else {
			m.gPressed = true
		}
		return m, nil

	case key.Matches(msg, slideKeys.Last):
		// Optional: number + G = go to slide N
		if m.numBuf != "" {
			n := parseNum(m.numBuf) - 1 // slides are 1-indexed in UX
			if n >= 0 && n < m.total {
				m.current = n
			}
			m.numBuf = ""
		} else {
			m.current = m.total - 1
		}
		m.gPressed = false
		m.refreshContent()
		return m, nil

	case key.Matches(msg, slideKeys.Next):
		offset := 1
		if m.numBuf != "" {
			offset = parseNum(m.numBuf)
			m.numBuf = ""
		}
		m.gPressed = false
		if m.current+offset < m.total {
			m.current += offset
		} else {
			m.current = m.total - 1
		}
		m.refreshContent()
		return m, nil

	case key.Matches(msg, slideKeys.Prev):
		offset := 1
		if m.numBuf != "" {
			offset = parseNum(m.numBuf)
			m.numBuf = ""
		}
		m.gPressed = false
		if m.current-offset >= 0 {
			m.current -= offset
		} else {
			m.current = 0
		}
		m.refreshContent()
		return m, nil

	case key.Matches(msg, slideKeys.Search):
		m.mode = slideModeSearch
		m.searchInput.Reset()
		m.searchInput.Focus()
		m.gPressed = false
		m.numBuf = ""
		return m, textinput.Blink

	case key.Matches(msg, slideKeys.ScrollUp):
		m.viewport.LineUp(5)
		m.gPressed = false
		return m, nil

	case key.Matches(msg, slideKeys.ScrollDn):
		m.viewport.LineDown(5)
		m.gPressed = false
		return m, nil
	}

	m.gPressed = false
	m.numBuf = ""

	// Forward to viewport for scrolling
	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

// updateSearch handles keys in search mode.
func (m SlideModel) updateSearch(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		query := strings.TrimSpace(m.searchInput.Value())
		m.searchQuery = query
		m.mode = slideModePresenting
		m.searchInput.Blur()

		if query != "" {
			idx := m.findSlide(query, m.current)
			if idx >= 0 {
				m.current = idx
				m.searchResult = idx
			} else {
				m.searchResult = -1
			}
			m.refreshContent()
		}
		return m, nil

	case "esc", "ctrl+c":
		m.mode = slideModePresenting
		m.searchInput.Blur()
		return m, nil

	case "ctrl+n":
		// next search result
		if m.searchQuery != "" {
			start := (m.current + 1) % m.total
			idx := m.findSlide(m.searchQuery, start)
			if idx >= 0 {
				m.current = idx
				m.searchResult = idx
				m.refreshContent()
			}
		}
		return m, nil
	}

	var cmd tea.Cmd
	m.searchInput, cmd = m.searchInput.Update(msg)
	return m, cmd
}

// findSlide returns the index of the first slide at or after `startFrom`
// whose content (case-insensitive) contains the query string.
// Returns -1 if not found.
func (m SlideModel) findSlide(query string, startFrom int) int {
	q := strings.ToLower(query)
	for i := 0; i < m.total; i++ {
		idx := (startFrom + i) % m.total
		if strings.Contains(strings.ToLower(m.slides[idx].Content), q) {
			return idx
		}
	}
	return -1
}

// refreshContent updates the viewport content for the current slide.
func (m *SlideModel) refreshContent() {
	if !m.ready || m.total == 0 {
		return
	}
	m.viewport.SetContent(m.renderCurrentSlide())
	m.viewport.GotoTop()
}

// ──────────────────────────────────────────────
// View
// ──────────────────────────────────────────────

// View renders the full TUI.
func (m SlideModel) View() string {
	if !m.ready {
		return "\n  Loading slides..."
	}
	if m.total == 0 {
		return m.styles.ErrorMsg.Render("\n  No slides found in file. Make sure slides are separated by '---'.\n")
	}

	if m.showHelp {
		return m.renderHelpOverlay()
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.renderSlideHeader(),
		m.viewport.View(),
		m.renderSlideFooter(),
	)
}

func (m SlideModel) renderHelpOverlay() string {
	helpStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		Padding(1, 3).
		Width(m.width - 4).
		Align(lipgloss.Left)

	help := `Keybindings

  Navigation
  ──────────────────────────
  space / → / ↓ / enter    Next slide
  ← / ↑ / p / h / k        Previous slide
  gg                        First slide
  G                         Last slide
  <number> G                Jump to slide

  Scrolling
  ──────────────────────────
  ctrl+u                    Scroll up
  ctrl+d                    Scroll down

  Search
  ──────────────────────────
  /                         Search slides
  ctrl+n                    Next result

  Other
  ──────────────────────────
  ?                         Toggle this help
  q / ctrl+c                Quit

  Press any key to dismiss`

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, helpStyle.Render(help))
}

// parseNum parses a numeric string to int; returns 1 on failure.
func parseNum(s string) int {
	n := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 1
		}
		n = n*10 + int(c-'0')
	}
	if n == 0 {
		return 1
	}
	return n
}

// renderSlideHeader renders the top status bar.
func (m SlideModel) renderSlideHeader() string {
	leftStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("62")).
		Foreground(lipgloss.Color("0")).
		Bold(true).
		Padding(0, 2)

	rightStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("62")).
		Foreground(lipgloss.Color("0")).
		Padding(0, 2).
		Align(lipgloss.Right)

	title := "TryOutShell Slides"
	if m.lessonTitle != "" {
		title = m.lessonTitle
	}
	if m.total > 0 && m.slides[m.current].Title != "" {
		title = m.slides[m.current].Title
	}

	left := leftStyle.Render("🎞  " + title)
	right := rightStyle.Render(fmt.Sprintf("Slide %d / %d", m.current+1, m.total))

	leftW := lipgloss.Width(left)
	rightW := lipgloss.Width(right)
	midW := m.width - leftW - rightW
	if midW < 0 {
		midW = 0
	}

	mid := lipgloss.NewStyle().
		Background(lipgloss.Color("62")).
		Width(midW).
		Render("")

	return lipgloss.JoinHorizontal(lipgloss.Top, left, mid, right)
}

// renderSlideFooter renders the bottom help/search bar.
func (m SlideModel) renderSlideFooter() string {
	footerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Padding(0, 2)

	if m.mode == slideModeSearch {
		return footerStyle.Render("/ " + m.searchInput.View() + "  (Enter: jump, Esc: cancel, ctrl+n: next result)")
	}

	hints := "space/→: next  ←: prev  G: last  gg: first  /: search  q: quit"
	if m.numBuf != "" {
		hints = fmt.Sprintf("[%s]  %s", m.numBuf, hints)
	}

	// Progress bar
	progress := 0.0
	if m.total > 1 {
		progress = float64(m.current) / float64(m.total-1)
	}
	barWidth := m.width - 4
	if barWidth < 10 {
		barWidth = 10
	}
	filled := int(progress * float64(barWidth))
	if filled > barWidth {
		filled = barWidth
	}
	empty := barWidth - filled

	filledStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("62"))
	bar := filledStyle.Render(strings.Repeat("─", filled)) + strings.Repeat("─", empty)

	return footerStyle.Render(bar+"\n"+hints)
}
