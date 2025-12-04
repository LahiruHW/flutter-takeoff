# Quick Reference: Creating Beautiful CLIs with Go

A practical guide to the libraries and patterns used in this project.

## Table of Contents
1. [Bubble Tea Basics](#bubble-tea-basics)
2. [Lipgloss Styling](#lipgloss-styling)
3. [Bubbles Components](#bubbles-components)
4. [Common Patterns](#common-patterns)
5. [Keyboard Shortcuts](#keyboard-shortcuts)

---

## Bubble Tea Basics

### The Three Core Methods

Every Bubble Tea program implements these three methods:

```go
type Model struct {
    // Your state here
}

// Init: Initialize your model
func (m Model) Init() tea.Cmd {
    return nil  // or return a command to execute
}

// Update: Handle messages and update state
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    return m, nil
}

// View: Render your UI
func (m Model) View() string {
    return "Hello, World!"
}
```

### Running a Program

```go
func main() {
    p := tea.NewProgram(Model{})
    if _, err := p.Run(); err != nil {
        fmt.Printf("Error: %v", err)
        os.Exit(1)
    }
}
```

### Message Types

```go
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        // Handle keyboard input
        switch msg.String() {
        case "q", "ctrl+c":
            return m, tea.Quit
        case "enter":
            // Handle enter key
        }
    
    case tea.WindowSizeMsg:
        // Handle window resize
        m.width = msg.Width
        m.height = msg.Height
    
    case MyCustomMsg:
        // Handle custom messages
    }
    return m, nil
}
```

---

## Lipgloss Styling

### Basic Styles

```go
import "github.com/charmbracelet/lipgloss"

// Create a style
titleStyle := lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#7C3AED")).
    Background(lipgloss.Color("#1F2937")).
    Padding(1, 2).
    Margin(1, 0)

// Use the style
fmt.Println(titleStyle.Render("My Title"))
```

### Colors

```go
// Hex colors
color := lipgloss.Color("#FF5733")

// ANSI colors (0-255)
color := lipgloss.Color("196")

// Named colors
lipgloss.AdaptiveColor{
    Light: "#000000",  // For light terminals
    Dark:  "#FFFFFF",  // For dark terminals
}
```

### Layout Properties

```go
style := lipgloss.NewStyle().
    Width(50).              // Width
    Height(10).             // Height
    Align(lipgloss.Center). // Alignment: Left, Center, Right
    Padding(1, 2, 1, 2).   // Top, Right, Bottom, Left
    Margin(1, 0).          // Top/Bottom, Left/Right
    Border(lipgloss.RoundedBorder()).
    BorderForeground(lipgloss.Color("#7C3AED"))
```

### Text Formatting

```go
style := lipgloss.NewStyle().
    Bold(true).
    Italic(true).
    Underline(true).
    Strikethrough(true).
    Faint(true).
    Blink(true)
```

### Borders

```go
// Border styles
lipgloss.NormalBorder()
lipgloss.RoundedBorder()
lipgloss.BlockBorder()
lipgloss.OuterHalfBlockBorder()
lipgloss.InnerHalfBlockBorder()
lipgloss.ThickBorder()
lipgloss.DoubleBorder()

// Custom border
customBorder := lipgloss.Border{
    Top:         "─",
    Bottom:      "─",
    Left:        "│",
    Right:       "│",
    TopLeft:     "╭",
    TopRight:    "╮",
    BottomLeft:  "╰",
    BottomRight: "╯",
}

style := lipgloss.NewStyle().Border(customBorder)
```

### Composing Styles

```go
// Join horizontally
lipgloss.JoinHorizontal(lipgloss.Top, "Hello", "World")

// Join vertically
lipgloss.JoinVertical(lipgloss.Left, "Line 1", "Line 2")

// Place in position
lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, "Text")
```

---

## Bubbles Components

### List Component

```go
import "github.com/charmbracelet/bubbles/list"

// Define list item
type item struct {
    title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

// Create list
items := []list.Item{
    item{title: "Option 1", desc: "First option"},
    item{title: "Option 2", desc: "Second option"},
}

l := list.New(items, list.NewDefaultDelegate(), 80, 14)
l.Title = "Choose an option"

// In Update method
var cmd tea.Cmd
m.list, cmd = m.list.Update(msg)

// In View method
return m.list.View()
```

### Spinner Component

```go
import "github.com/charmbracelet/bubbles/spinner"

// Create spinner
s := spinner.New()
s.Spinner = spinner.Dot       // Dot, Line, Globe, etc.
s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

// In Init
func (m Model) Init() tea.Cmd {
    return m.spinner.Tick
}

// In Update
case spinner.TickMsg:
    var cmd tea.Cmd
    m.spinner, cmd = m.spinner.Update(msg)
    return m, cmd

// In View
return m.spinner.View() + " Loading..."
```

### Progress Bar

```go
import "github.com/charmbracelet/bubbles/progress"

// Create progress bar
prog := progress.New(progress.WithDefaultGradient())

// Set progress (0.0 to 1.0)
prog.SetPercent(0.5)

// In View
return prog.View()
```

### Text Input

```go
import "github.com/charmbracelet/bubbles/textinput"

// Create text input
ti := textinput.New()
ti.Placeholder = "Enter your name..."
ti.Focus()

// In Init
func (m Model) Init() tea.Cmd {
    return textinput.Blink
}

// In Update
var cmd tea.Cmd
m.textInput, cmd = m.textInput.Update(msg)

// Get value
value := m.textInput.Value()

// In View
return m.textInput.View()
```

---

## Common Patterns

### Loading State

```go
type Model struct {
    loading  bool
    spinner  spinner.Model
    err      error
}

func (m Model) View() string {
    if m.loading {
        return m.spinner.View() + " Loading..."
    }
    if m.err != nil {
        return "Error: " + m.err.Error()
    }
    return "Content loaded!"
}
```

### Multi-Step Form

```go
type step int

const (
    stepName step = iota
    stepEmail
    stepConfirm
)

type Model struct {
    step       step
    nameInput  textinput.Model
    emailInput textinput.Model
}

func (m Model) View() string {
    switch m.step {
    case stepName:
        return "Name: " + m.nameInput.View()
    case stepEmail:
        return "Email: " + m.emailInput.View()
    case stepConfirm:
        return "Confirm your details..."
    }
}
```

### Tabs/Navigation

```go
type tab int

const (
    homeTab tab = iota
    settingsTab
    aboutTab
)

type Model struct {
    activeTab tab
}

func (m Model) View() string {
    tabs := []string{"Home", "Settings", "About"}
    var renderedTabs []string
    
    for i, t := range tabs {
        style := normalTabStyle
        if tab(i) == m.activeTab {
            style = activeTabStyle
        }
        renderedTabs = append(renderedTabs, style.Render(t))
    }
    
    header := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
    return header + "\n\n" + m.renderContent()
}
```

### Command/Long-Running Tasks

```go
// Define a message type
type resultMsg struct {
    result string
    err    error
}

// Create a command
func doWork() tea.Msg {
    // Simulate work
    time.Sleep(2 * time.Second)
    return resultMsg{result: "Done!"}
}

// In Update
case tea.KeyMsg:
    if msg.String() == "enter" {
        return m, doWork  // Return command
    }

case resultMsg:
    m.result = msg.result
    m.err = msg.err
```

---

## Keyboard Shortcuts

### Default Keys

```go
case tea.KeyMsg:
    switch msg.String() {
    case "ctrl+c", "q":
        return m, tea.Quit
    
    case "enter", "return":
        // Submit/Select
    
    case "esc":
        // Cancel/Back
    
    case "up", "k":
        // Move up
    
    case "down", "j":
        // Move down
    
    case "left", "h":
        // Move left
    
    case "right", "l":
        // Move right
    
    case " ", "space":
        // Toggle/Select
    
    case "tab":
        // Next field
    
    case "shift+tab":
        // Previous field
    }
```

### Key Types

```go
case tea.KeyMsg:
    switch msg.Type {
    case tea.KeyCtrlC:
        return m, tea.Quit
    
    case tea.KeyEnter:
        // Handle enter
    
    case tea.KeySpace:
        // Handle space
    
    case tea.KeyUp, tea.KeyDown:
        // Handle arrows
    
    case tea.KeyRunes:
        // Handle character input
        char := msg.Runes[0]
    }
```

---

## Quick Examples

### Simple Menu

```go
package main

import (
    "fmt"
    "os"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

var itemStyle = lipgloss.NewStyle().PaddingLeft(4)
var selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))

type model struct {
    choices  []string
    cursor   int
    selected map[int]struct{}
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            return m, tea.Quit
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }
        case "down", "j":
            if m.cursor < len(m.choices)-1 {
                m.cursor++
            }
        case " ":
            _, ok := m.selected[m.cursor]
            if ok {
                delete(m.selected, m.cursor)
            } else {
                m.selected[m.cursor] = struct{}{}
            }
        case "enter":
            return m, tea.Quit
        }
    }
    return m, nil
}

func (m model) View() string {
    s := "What should we buy?\n\n"

    for i, choice := range m.choices {
        cursor := " "
        if m.cursor == i {
            cursor = ">"
        }

        checked := " "
        if _, ok := m.selected[i]; ok {
            checked = "x"
        }

        s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
    }

    s += "\nPress q to quit.\n"
    return s
}

func main() {
    p := tea.NewProgram(model{
        choices:  []string{"Apples", "Oranges", "Bananas"},
        selected: make(map[int]struct{}),
    })

    if _, err := p.Run(); err != nil {
        fmt.Printf("Error: %v", err)
        os.Exit(1)
    }
}
```

### Progress Indicator

```go
type model struct {
    progress float64
}

func (m model) Init() tea.Cmd {
    return tickCmd()
}

func tickCmd() tea.Cmd {
    return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
        return tickMsg(t)
    })
}

type tickMsg time.Time

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg.(type) {
    case tickMsg:
        m.progress += 0.01
        if m.progress >= 1.0 {
            return m, tea.Quit
        }
        return m, tickCmd()
    case tea.KeyMsg:
        return m, tea.Quit
    }
    return m, nil
}

func (m model) View() string {
    prog := int(m.progress * 50)
    return fmt.Sprintf("\n  %s%s %3.0f%%\n\n",
        strings.Repeat("█", prog),
        strings.Repeat("░", 50-prog),
        m.progress*100,
    )
}
```

---

## Resources

- [Bubble Tea Docs](https://github.com/charmbracelet/bubbletea)
- [Lipgloss Docs](https://github.com/charmbracelet/lipgloss)
- [Bubbles Docs](https://github.com/charmbracelet/bubbles)
- [Examples](https://github.com/charmbracelet/bubbletea/tree/master/examples)
- [Charm.sh](https://charm.sh/)

---

**Happy CLI Building!** ✨
