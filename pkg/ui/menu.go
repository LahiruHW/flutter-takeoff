package ui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MenuItem struct {
	Title       string
	Description string
	Value       string
}

func (i MenuItem) FilterValue() string { return i.Title }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(MenuItem)
	if !ok {
		return
	}

	str := i.Title

	if index == m.Index() {
		fmt.Fprint(w, SelectedItemStyle.Render("▸ "+str))
	} else {
		fmt.Fprint(w, UnselectedItemStyle.Render("  "+str))
	}
}

type MenuModel struct {
	list     list.Model
	choice   string
	quitting bool
}

func NewMenu(title string, items []MenuItem, width, height int) MenuModel {
	listItems := make([]list.Item, len(items))
	for i, item := range items {
		listItems[i] = item
	}

	l := list.New(listItems, itemDelegate{}, width, height)
	l.Title = title
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = TitleStyle
	l.Styles.PaginationStyle = SubtleStyle
	l.Styles.HelpStyle = HelpStyle

	return MenuModel{list: l}
}

func (m MenuModel) Init() tea.Cmd {
	return nil
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(MenuItem)
			if ok {
				m.choice = i.Value
			}
			m.quitting = true
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m MenuModel) View() string {
	if m.quitting {
		return ""
	}
	return "\n" + m.list.View()
}

func (m MenuModel) Choice() string {
	return m.choice
}

// SimpleMenu creates a simple text-based menu without full TUI
func SimpleMenu(title string, options []string) string {
	var b strings.Builder

	b.WriteString(TitleStyle.Render(title) + "\n\n")

	for i, opt := range options {
		b.WriteString(fmt.Sprintf("%s %d. %s\n",
			SubtleStyle.Render("→"),
			i+1,
			NormalStyle.Render(opt)))
	}

	return b.String()
}

// ConfirmPrompt creates a yes/no confirmation prompt
func ConfirmPrompt(message string) string {
	return WarningStyle.Render("?") + " " + message + " " +
		SubtleStyle.Render("(y/n)")
}

// Header creates a styled header
func Header(text string) string {
	return "\n" + lipgloss.NewStyle().
		Bold(true).
		Foreground(PrimaryColor).
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderForeground(PrimaryColor).
		Padding(0, 1).
		Render(text) + "\n"
}
