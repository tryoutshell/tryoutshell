package reader

import "github.com/charmbracelet/lipgloss"

type ReaderTheme struct {
	Name      string
	ArticleBg lipgloss.Color
	ArticleFg lipgloss.Color
	ChatBg    lipgloss.Color
	ChatFg    lipgloss.Color
	Border    lipgloss.Color
	Accent    lipgloss.Color
	Header    lipgloss.Color
}

var readerThemes = []ReaderTheme{
	{
		Name:      "default",
		ArticleBg: lipgloss.Color(""),
		ArticleFg: lipgloss.Color("252"),
		ChatBg:    lipgloss.Color(""),
		ChatFg:    lipgloss.Color("252"),
		Border:    lipgloss.Color("63"),
		Accent:    lipgloss.Color("63"),
		Header:    lipgloss.Color("63"),
	},
	{
		Name:      "dark",
		ArticleBg: lipgloss.Color("235"),
		ArticleFg: lipgloss.Color("252"),
		ChatBg:    lipgloss.Color("236"),
		ChatFg:    lipgloss.Color("252"),
		Border:    lipgloss.Color("240"),
		Accent:    lipgloss.Color("39"),
		Header:    lipgloss.Color("39"),
	},
	{
		Name:      "light",
		ArticleBg: lipgloss.Color("231"),
		ArticleFg: lipgloss.Color("232"),
		ChatBg:    lipgloss.Color("255"),
		ChatFg:    lipgloss.Color("232"),
		Border:    lipgloss.Color("244"),
		Accent:    lipgloss.Color("25"),
		Header:    lipgloss.Color("25"),
	},
	{
		Name:      "dracula",
		ArticleBg: lipgloss.Color("236"),
		ArticleFg: lipgloss.Color("253"),
		ChatBg:    lipgloss.Color("237"),
		ChatFg:    lipgloss.Color("253"),
		Border:    lipgloss.Color("141"),
		Accent:    lipgloss.Color("212"),
		Header:    lipgloss.Color("141"),
	},
	{
		Name:      "nord",
		ArticleBg: lipgloss.Color("236"),
		ArticleFg: lipgloss.Color("254"),
		ChatBg:    lipgloss.Color("237"),
		ChatFg:    lipgloss.Color("254"),
		Border:    lipgloss.Color("110"),
		Accent:    lipgloss.Color("110"),
		Header:    lipgloss.Color("110"),
	},
}

func getReaderTheme(name string) ReaderTheme {
	for _, t := range readerThemes {
		if t.Name == name {
			return t
		}
	}
	return readerThemes[0]
}

func nextThemeName(current string) string {
	for i, t := range readerThemes {
		if t.Name == current {
			return readerThemes[(i+1)%len(readerThemes)].Name
		}
	}
	return readerThemes[0].Name
}
