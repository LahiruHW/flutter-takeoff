package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Color scheme
	PrimaryColor   = lipgloss.Color("#7C3AED") // Purple
	SecondaryColor = lipgloss.Color("#10B981") // Green
	AccentColor    = lipgloss.Color("#F59E0B") // Amber
	ErrorColor     = lipgloss.Color("#EF4444") // Red
	TextColor      = lipgloss.Color("#E5E7EB") // Light gray
	SubtleColor    = lipgloss.Color("#9CA3AF") // Gray

	// Title styles
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(PrimaryColor).
			MarginTop(1).
			MarginBottom(1)

	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(SecondaryColor).
			MarginTop(1)

	// Text styles
	NormalStyle = lipgloss.NewStyle().
			Foreground(TextColor)

	SubtleStyle = lipgloss.NewStyle().
			Foreground(SubtleColor)

	SuccessStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(SecondaryColor)

	ErrorStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ErrorColor)

	WarningStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(AccentColor)

	// Component styles
	SelectedItemStyle = lipgloss.NewStyle().
				Foreground(PrimaryColor).
				Bold(true).
				PaddingLeft(2)

	UnselectedItemStyle = lipgloss.NewStyle().
				Foreground(TextColor).
				PaddingLeft(2)

	CheckboxStyle = lipgloss.NewStyle().
			Foreground(SecondaryColor).
			Bold(true)

	// Container styles
	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(PrimaryColor).
			Padding(1, 2)

	HelpStyle = lipgloss.NewStyle().
			Foreground(SubtleColor).
			MarginTop(1)
)

// Checkbox renders a checkbox with label
func Checkbox(checked bool, label string) string {
	box := "☐"
	style := UnselectedItemStyle
	if checked {
		box = "☑"
		style = CheckboxStyle
	}
	return style.Render(box + " " + label)
}

// StatusIndicator renders a status with icon
func StatusIndicator(status string, message string) string {
	var icon string
	var style lipgloss.Style

	switch status {
	case "success":
		icon = "✓"
		style = SuccessStyle
	case "error":
		icon = "✗"
		style = ErrorStyle
	case "warning":
		icon = "⚠"
		style = WarningStyle
	case "info":
		icon = `ℹ`
		style = NormalStyle
	default:
		icon = "•"
		style = SubtleStyle
	}

	return style.Render(icon) + " " + message
}
