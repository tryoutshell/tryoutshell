package reader

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type pane int

const (
	paneArticle pane = iota
	paneChat
)

type chatResponseMsg struct {
	content string
	err     error
}

type ReaderModel struct {
	title   string
	article string
	url     string

	width  int
	height int
	ready  bool

	activePane pane
	splitRatio float64

	articleVP viewport.Model
	chatVP    viewport.Model
	chatInput textinput.Model

	chatHistory []ChatMessage
	chatClient  *ChatClient
	waiting     bool

	themeName string
	theme     ReaderTheme
}

func NewReaderModel(title, markdown, url string) ReaderModel {
	ti := textinput.New()
	ti.Placeholder = "Ask about this article..."
	ti.CharLimit = 500

	theme := getReaderTheme("default")
	client := NewChatClient(markdown)

	return ReaderModel{
		title:      title,
		article:    markdown,
		url:        url,
		splitRatio: 0.6,
		chatInput:  ti,
		chatClient: client,
		themeName:  "default",
		theme:      theme,
	}
}

func (m ReaderModel) Init() tea.Cmd {
	return nil
}

func (m ReaderModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m = m.initViewports()
		return m, nil

	case chatResponseMsg:
		m.waiting = false
		if msg.err != nil {
			m.chatHistory = append(m.chatHistory, ChatMessage{Role: "assistant", Content: fmt.Sprintf("Error: %v", msg.err)})
		} else {
			m.chatHistory = append(m.chatHistory, ChatMessage{Role: "assistant", Content: msg.content})
		}
		m.chatVP.SetContent(m.renderChatHistory())
		m.chatVP.GotoBottom()
		return m, nil

	case tea.KeyMsg:
		return m.handleKey(msg)
	}

	var cmds []tea.Cmd
	if m.activePane == paneArticle {
		var cmd tea.Cmd
		m.articleVP, cmd = m.articleVP.Update(msg)
		cmds = append(cmds, cmd)
	} else {
		var cmd tea.Cmd
		m.chatVP, cmd = m.chatVP.Update(msg)
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m ReaderModel) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	k := msg.String()

	switch k {
	case "ctrl+c":
		return m, tea.Quit
	case "tab":
		if m.activePane == paneArticle {
			m.activePane = paneChat
			m.chatInput.Focus()
		} else {
			m.activePane = paneArticle
			m.chatInput.Blur()
		}
		return m, nil
	case "[":
		if m.splitRatio > 0.2 {
			m.splitRatio -= 0.05
			m = m.initViewports()
		}
		return m, nil
	case "]":
		if m.splitRatio < 0.8 {
			m.splitRatio += 0.05
			m = m.initViewports()
		}
		return m, nil
	case "t":
		if m.activePane == paneArticle {
			m.themeName = nextThemeName(m.themeName)
			m.theme = getReaderTheme(m.themeName)
			return m, nil
		}
	}

	if m.activePane == paneArticle {
		return m.handleArticleKey(msg)
	}
	return m.handleChatKey(msg)
}

func (m ReaderModel) handleArticleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q":
		return m, tea.Quit
	case "j", "down":
		m.articleVP.LineDown(1)
	case "k", "up":
		m.articleVP.LineUp(1)
	case "d", "ctrl+d":
		m.articleVP.HalfViewDown()
	case "u", "ctrl+u":
		m.articleVP.HalfViewUp()
	case "G":
		m.articleVP.GotoBottom()
	case "g":
		m.articleVP.GotoTop()
	}

	var cmd tea.Cmd
	m.articleVP, cmd = m.articleVP.Update(msg)
	return m, cmd
}

func (m ReaderModel) handleChatKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		input := strings.TrimSpace(m.chatInput.Value())
		if input == "" || m.waiting {
			return m, nil
		}

		if !m.chatClient.Available() {
			m.chatHistory = append(m.chatHistory, ChatMessage{Role: "user", Content: input})
			m.chatHistory = append(m.chatHistory, ChatMessage{
				Role:    "assistant",
				Content: "No API key found. Set OPENAI_API_KEY, ANTHROPIC_API_KEY, or GEMINI_API_KEY to enable chat.",
			})
			m.chatInput.Reset()
			m.chatVP.SetContent(m.renderChatHistory())
			m.chatVP.GotoBottom()
			return m, nil
		}

		m.chatHistory = append(m.chatHistory, ChatMessage{Role: "user", Content: input})
		m.chatInput.Reset()
		m.waiting = true
		m.chatVP.SetContent(m.renderChatHistory())
		m.chatVP.GotoBottom()

		client := m.chatClient
		return m, func() tea.Msg {
			reply, err := client.Send(input)
			return chatResponseMsg{content: reply, err: err}
		}

	case "esc":
		m.activePane = paneArticle
		m.chatInput.Blur()
		return m, nil
	}

	var cmd tea.Cmd
	m.chatInput, cmd = m.chatInput.Update(msg)
	return m, cmd
}

func (m ReaderModel) initViewports() ReaderModel {
	headerH := 1
	footerH := 1
	contentH := m.height - headerH - footerH
	if contentH < 1 {
		contentH = 1
	}

	leftW := int(float64(m.width) * m.splitRatio)
	rightW := m.width - leftW - 1
	if rightW < 10 {
		rightW = 10
	}

	m.articleVP = viewport.New(leftW-2, contentH)
	m.articleVP.SetContent(m.article)

	chatContentH := contentH - 3
	if chatContentH < 1 {
		chatContentH = 1
	}
	m.chatVP = viewport.New(rightW-2, chatContentH)
	m.chatVP.SetContent(m.renderChatHistory())

	m.chatInput.Width = rightW - 4
	m.ready = true
	return m
}

func (m ReaderModel) renderChatHistory() string {
	if len(m.chatHistory) == 0 {
		hint := "Ask questions about the article here."
		if !m.chatClient.Available() {
			hint += "\n\n(No API key detected. Set OPENAI_API_KEY, ANTHROPIC_API_KEY, or GEMINI_API_KEY)"
		} else {
			hint += fmt.Sprintf("\n\nUsing %s", m.chatClient.Provider())
		}
		return lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render(hint)
	}

	var lines []string
	userStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("63")).Bold(true)
	aiStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("252"))

	for _, msg := range m.chatHistory {
		if msg.Role == "user" {
			lines = append(lines, userStyle.Render("You: ")+msg.Content)
		} else {
			lines = append(lines, aiStyle.Render("AI: ")+msg.Content)
		}
		lines = append(lines, "")
	}

	if m.waiting {
		lines = append(lines, lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("Thinking..."))
	}

	return strings.Join(lines, "\n")
}

func (m ReaderModel) View() string {
	if !m.ready {
		return "\n  Loading..."
	}

	header := m.renderHeader()
	body := m.renderBody()
	footer := m.renderFooter()

	return lipgloss.JoinVertical(lipgloss.Left, header, body, footer)
}

func (m ReaderModel) renderHeader() string {
	style := lipgloss.NewStyle().
		Background(lipgloss.Color(string(m.theme.Header))).
		Foreground(lipgloss.Color("0")).
		Bold(true).
		Padding(0, 2).
		Width(m.width)

	title := m.title
	if len(title) > m.width-10 {
		title = title[:m.width-13] + "..."
	}

	return style.Render("📖 " + title)
}

func (m ReaderModel) renderBody() string {
	headerH := 1
	footerH := 1
	contentH := m.height - headerH - footerH
	if contentH < 1 {
		contentH = 1
	}

	leftW := int(float64(m.width) * m.splitRatio)
	rightW := m.width - leftW - 1
	if rightW < 10 {
		rightW = 10
	}

	borderColor := m.theme.Border
	activeBorder := lipgloss.NewStyle().Foreground(lipgloss.Color(string(m.theme.Accent)))
	inactiveBorder := lipgloss.NewStyle().Foreground(lipgloss.Color(string(borderColor)))

	leftBorderStyle := inactiveBorder
	rightBorderStyle := inactiveBorder
	if m.activePane == paneArticle {
		leftBorderStyle = activeBorder
	} else {
		rightBorderStyle = activeBorder
	}

	leftPane := leftBorderStyle.
		Border(lipgloss.RoundedBorder()).
		Width(leftW - 2).
		Height(contentH - 2).
		Render(m.articleVP.View())

	chatContent := m.chatVP.View() + "\n" + m.chatInput.View()
	rightPane := rightBorderStyle.
		Border(lipgloss.RoundedBorder()).
		Width(rightW - 2).
		Height(contentH - 2).
		Render(chatContent)

	separator := lipgloss.NewStyle().
		Foreground(lipgloss.Color(string(borderColor))).
		Render(strings.Repeat("│\n", contentH))

	return lipgloss.JoinHorizontal(lipgloss.Top, leftPane, separator, rightPane)
}

func (m ReaderModel) renderFooter() string {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Width(m.width)

	paneLabel := "article"
	if m.activePane == paneChat {
		paneLabel = "chat"
	}

	return style.Render(fmt.Sprintf(" tab: switch pane (%s) │ [/]: resize │ t: theme (%s) │ j/k: scroll │ q: quit", paneLabel, m.themeName))
}

func Launch(title, markdown, url string) error {
	m := NewReaderModel(title, markdown, url)
	p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseCellMotion())
	_, err := p.Run()
	return err
}
