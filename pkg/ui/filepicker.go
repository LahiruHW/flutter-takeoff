package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// FilePickerModel represents a simple directory browser
type FilePickerModel struct {
	currentPath string
	items       []fileItem
	cursor      int
	selected    string
	done        bool
	showHidden  bool
	height      int
	startIndex  int // For scrolling
}

type fileItem struct {
	name  string
	path  string
	isDir bool
}

// NewFilePicker creates a new file picker starting at the given path
func NewFilePicker(startPath string, height int) FilePickerModel {
	if startPath == "" {
		startPath = os.Getenv("USERPROFILE") // Default to user home on Windows
	}

	absPath, _ := filepath.Abs(startPath)

	return FilePickerModel{
		currentPath: absPath,
		height:      height,
		showHidden:  false,
	}
}

// Init initializes the file picker
func (m FilePickerModel) Init() tea.Cmd {
	return nil
}

// Update handles user input
func (m FilePickerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.done = true
			m.selected = "" // Cancelled
			return m, tea.Quit

		case "enter":
			if len(m.items) == 0 {
				return m, nil
			}

			selectedItem := m.items[m.cursor]

			if selectedItem.isDir {
				// Navigate into directory
				m.currentPath = selectedItem.path
				m.cursor = 0
				m.startIndex = 0
				m.loadDirectory()
			} else {
				// This shouldn't happen in directory picker, but handle it
				m.selected = m.currentPath
				m.done = true
				return m, tea.Quit
			}

		case "s":
			// Select current directory
			m.selected = m.currentPath
			m.done = true
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
				// Scroll up if needed
				if m.cursor < m.startIndex {
					m.startIndex = m.cursor
				}
			}

		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
				// Scroll down if needed
				visibleItems := m.height - 6 // Account for header/footer
				if m.cursor >= m.startIndex+visibleItems {
					m.startIndex = m.cursor - visibleItems + 1
				}
			}

		case "h":
			// Toggle hidden files
			m.showHidden = !m.showHidden
			m.cursor = 0
			m.startIndex = 0
			m.loadDirectory()

		case "home":
			m.cursor = 0
			m.startIndex = 0

		case "end":
			m.cursor = len(m.items) - 1
			visibleItems := m.height - 6
			if len(m.items) > visibleItems {
				m.startIndex = len(m.items) - visibleItems
			}

		case "pgup":
			visibleItems := m.height - 6
			m.cursor -= visibleItems
			if m.cursor < 0 {
				m.cursor = 0
			}
			m.startIndex = m.cursor

		case "pgdown":
			visibleItems := m.height - 6
			m.cursor += visibleItems
			if m.cursor >= len(m.items) {
				m.cursor = len(m.items) - 1
			}
			if m.cursor >= m.startIndex+visibleItems {
				m.startIndex = m.cursor - visibleItems + 1
			}
		}
	}

	return m, nil
}

// View renders the file picker
func (m FilePickerModel) View() string {
	if m.done {
		return ""
	}

	// Lazy load directory if items not loaded
	if m.items == nil {
		m.loadDirectory()
	}

	var b strings.Builder

	// Header
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("39")).
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		Width(80)

	b.WriteString(headerStyle.Render(fmt.Sprintf("📁 Select Directory: %s", m.currentPath)))
	b.WriteString("\n\n")

	// Items
	visibleItems := m.height - 6
	endIndex := m.startIndex + visibleItems
	if endIndex > len(m.items) {
		endIndex = len(m.items)
	}

	if len(m.items) == 0 {
		b.WriteString(SubtleStyle.Render("  (empty directory)\n"))
	} else {
		for i := m.startIndex; i < endIndex; i++ {
			item := m.items[i]
			cursor := "  "
			if m.cursor == i {
				cursor = "→ "
			}

			icon := "📄"
			if item.isDir {
				if item.name == ".." {
					icon = "⬆️ "
				} else {
					icon = "📁"
				}
			}

			nameStyle := NormalStyle
			if m.cursor == i {
				nameStyle = lipgloss.NewStyle().
					Foreground(lipgloss.Color("39")).
					Bold(true)
			}

			b.WriteString(fmt.Sprintf("%s%s %s\n", cursor, icon, nameStyle.Render(item.name)))
		}
	}

	// Scrolling indicator
	if len(m.items) > visibleItems {
		b.WriteString(SubtleStyle.Render(fmt.Sprintf("\n  (showing %d-%d of %d)",
			m.startIndex+1, endIndex, len(m.items))))
	}

	// Footer
	b.WriteString("\n\n")
	b.WriteString(SubtleStyle.Render("  ↑/↓: Navigate  Enter: Open folder  S: Select this directory"))
	b.WriteString("\n")
	b.WriteString(SubtleStyle.Render("  H: Toggle hidden files  Esc: Cancel"))

	return b.String()
}

// loadDirectory loads the contents of the current directory
func (m *FilePickerModel) loadDirectory() {
	m.items = []fileItem{}

	// Add parent directory option if not at root
	if m.currentPath != filepath.VolumeName(m.currentPath)+string(filepath.Separator) {
		parentPath := filepath.Dir(m.currentPath)
		m.items = append(m.items, fileItem{
			name:  "..",
			path:  parentPath,
			isDir: true,
		})
	}

	// Read directory contents
	entries, err := os.ReadDir(m.currentPath)
	if err != nil {
		// If we can't read the directory, just show error item
		return
	}

	// Separate directories and files
	var dirs []fileItem

	for _, entry := range entries {
		// Skip hidden files if not showing them
		if !m.showHidden && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		if entry.IsDir() {
			dirs = append(dirs, fileItem{
				name:  entry.Name(),
				path:  filepath.Join(m.currentPath, entry.Name()),
				isDir: true,
			})
		}
	}

	// Add directories (sorted)
	m.items = append(m.items, dirs...)
}

// Selected returns the selected path
func (m FilePickerModel) Selected() string {
	return m.selected
}

// IsDone returns whether selection is complete
func (m FilePickerModel) IsDone() bool {
	return m.done
}
