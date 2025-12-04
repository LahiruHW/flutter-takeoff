package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ProgressMsg struct {
	Percent int
	Status  string
	Done    bool
}

type ProgressModel struct {
	spinner  spinner.Model
	progress int
	status   string
	done     bool
}

func NewProgress() ProgressModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(PrimaryColor)
	return ProgressModel{
		spinner: s,
	}
}

func (m ProgressModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m ProgressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ProgressMsg:
		m.progress = msg.Percent
		m.status = msg.Status
		m.done = msg.Done
		if m.done {
			return m, tea.Quit
		}
		return m, nil

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m ProgressModel) View() string {
	if m.done {
		return SuccessStyle.Render("✓ Complete!\n")
	}

	var b strings.Builder

	// Progress bar
	width := 40
	filled := int(float64(width) * float64(m.progress) / 100.0)
	bar := strings.Repeat("█", filled) + strings.Repeat("░", width-filled)

	b.WriteString(fmt.Sprintf("\n %s %s %3d%%\n\n",
		m.spinner.View(),
		lipgloss.NewStyle().Foreground(SecondaryColor).Render(bar),
		m.progress))

	// Status message
	b.WriteString(SubtleStyle.Render("  " + m.status + "\n"))

	return b.String()
}

// SimpleProgressBar renders a static progress bar
func SimpleProgressBar(percent int, width int) string {
	if width < 10 {
		width = 40
	}

	filled := int(float64(width) * float64(percent) / 100.0)
	bar := strings.Repeat("█", filled) + strings.Repeat("░", width-filled)

	return lipgloss.NewStyle().
		Foreground(SecondaryColor).
		Render(bar) +
		fmt.Sprintf(" %3d%%", percent)
}
