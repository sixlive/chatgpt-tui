package settings

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tearingItUp786/golang-tui/util"
)

// var subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}

type Model struct {
	terminalWidth int
	isFocused     bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

func listItem(heading string, value string) string {
	headingEl := lipgloss.NewStyle().
		PaddingLeft(2).
		Foreground(lipgloss.Color("#FFC0CB")).
		Render
	spanEl := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#fff")).
		Render

	return headingEl(" "+heading, spanEl(value))
}

func (m Model) View() string {
	borderColor := lipgloss.Color("#bbb")
	if m.isFocused {
		borderColor = lipgloss.Color("#d70073")
	}
	list := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), true).
		BorderForeground(borderColor).
		Height(8).
		Width(m.terminalWidth/3 - 5)

	listHeader := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		MarginLeft(2).
		Render

	return list.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			listHeader("Settings"),
			listItem("model", "gpt-4.5"),
			listItem("frequency", "something"),
			listItem("max_tokens", fmt.Sprint(300)),
		),
	)
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case util.FocusEvent:
		m.isFocused = msg.IsFocused
		return m, nil
	case tea.WindowSizeMsg:
		m.terminalWidth = msg.Width
		return m, nil
	}
	return m, nil
}

func New() Model {
	return Model{
		terminalWidth: 20,
	}
}
