# Development Guide

## Getting Started with CLI Development in Go

This guide will help you understand how to create beautiful command-line applications using Go and the Charm libraries.

## Key Concepts

### 1. The Elm Architecture (TEA)

Bubble Tea is based on The Elm Architecture, which consists of three main components:

- **Model**: The application state
- **Update**: How the state changes in response to messages
- **View**: How to render the state to the terminal

```go
type Model struct {
    // Your application state
}

func (m Model) Init() tea.Cmd {
    // Initialize your model
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    // Handle messages and update state
}

func (m Model) View() string {
    // Render the UI
}
```

### 2. Styling with Lipgloss

Lipgloss allows you to style terminal output with colors, borders, padding, and more:

```go
import "github.com/charmbracelet/lipgloss"

titleStyle := lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#7C3AED")).
    MarginTop(1)

fmt.Println(titleStyle.Render("My Title"))
```

### 3. Common UI Components with Bubbles

Bubbles provides ready-to-use components:

- **list**: Interactive lists with filtering
- **spinner**: Loading animations
- **progress**: Progress bars
- **textinput**: Text input fields
- **table**: Data tables

## Color Scheme

The Flutter Installation Tool uses the following color palette:

| Color Name | Hex Code | Usage |
|------------|----------|-------|
| Primary | `#7C3AED` | Titles, borders, selected items |
| Secondary | `#10B981` | Success messages, checkboxes |
| Accent | `#F59E0B` | Warnings, important info |
| Error | `#EF4444` | Error messages |
| Text | `#E5E7EB` | Normal text |
| Subtle | `#9CA3AF` | Helper text, descriptions |

## Creating Custom Components

### Custom Menu Example

```go
type MenuItem struct {
    Title       string
    Description string
    Value       string
}

// Implement list.Item interface
func (i MenuItem) FilterValue() string { 
    return i.Title 
}

// Create a custom delegate for rendering
type itemDelegate struct{}

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
    i, ok := item.(MenuItem)
    if !ok {
        return
    }
    
    // Custom rendering logic
    if index == m.Index() {
        fmt.Fprint(w, selectedStyle.Render("▸ " + i.Title))
    } else {
        fmt.Fprint(w, normalStyle.Render("  " + i.Title))
    }
}
```

### Custom Progress Bar

```go
type ProgressModel struct {
    percent int
    status  string
}

func (m ProgressModel) View() string {
    width := 40
    filled := int(float64(width) * float64(m.percent) / 100.0)
    bar := strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
    
    return fmt.Sprintf("%s %3d%% - %s", bar, m.percent, m.status)
}
```

## Message Passing

Messages are how Bubble Tea applications communicate state changes:

```go
// Define custom messages
type ProgressMsg struct {
    Percent int
    Status  string
}

// Send messages
func downloadFile() tea.Msg {
    // Do work...
    return ProgressMsg{Percent: 50, Status: "Downloading..."}
}

// Handle messages in Update
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case ProgressMsg:
        m.progress = msg.Percent
        m.status = msg.Status
    }
    return m, nil
}
```

## Commands

Commands are functions that return messages:

```go
func checkDependenciesCmd() tea.Msg {
    // Perform checks
    deps := checkAllDependencies()
    return DependenciesCheckedMsg{Dependencies: deps}
}

// Use in Init or Update
func (m Model) Init() tea.Cmd {
    return checkDependenciesCmd
}
```

## Best Practices

### 1. Separate Concerns

- Keep UI code in `pkg/ui/`
- Keep business logic in `pkg/installer/`
- Keep data types in separate files

### 2. Use Consistent Styling

Define styles once and reuse them:

```go
var (
    TitleStyle   = lipgloss.NewStyle().Bold(true).Foreground(PrimaryColor)
    SuccessStyle = lipgloss.NewStyle().Bold(true).Foreground(SecondaryColor)
    ErrorStyle   = lipgloss.NewStyle().Bold(true).Foreground(ErrorColor)
)
```

### 3. Handle Errors Gracefully

Always provide helpful error messages:

```go
if err != nil {
    return ErrorStyle.Render("✗ " + err.Error())
}
```

### 4. Make It Interactive

Use keyboard controls:

```go
case tea.KeyMsg:
    switch msg.String() {
    case "q", "ctrl+c":
        return m, tea.Quit
    case "enter":
        // Handle selection
    case "up", "k":
        // Move up
    case "down", "j":
        // Move down
    }
```

## Testing

### Unit Testing UI Components

```go
func TestCheckbox(t *testing.T) {
    result := Checkbox(true, "Test")
    if !strings.Contains(result, "☑") {
        t.Error("Expected checked box")
    }
}
```

### Integration Testing

```go
func TestMenuSelection(t *testing.T) {
    items := []MenuItem{
        {Title: "Option 1", Value: "opt1"},
    }
    
    m := NewMenu("Test", items, 50, 10)
    // Simulate key presses
    m, _ = m.Update(tea.KeyMsg{Type: tea.KeyEnter})
    
    if m.Choice() != "opt1" {
        t.Error("Expected opt1 to be selected")
    }
}
```

## Debugging Tips

### 1. Log to File

Since stdout is used for the UI, log to a file:

```go
f, _ := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
log.SetOutput(f)
log.Println("Debug message")
```

### 2. Use Print Statements Carefully

Only print outside the Bubble Tea program:

```go
p := tea.NewProgram(model)
finalModel, err := p.Run()
// Now you can use fmt.Println safely
fmt.Println("Program finished")
```

### 3. Test Components Separately

Test UI components independently before integrating:

```go
// Test a style
styled := TitleStyle.Render("Test")
fmt.Println(styled)
```

## Performance Considerations

### 1. Minimize Renders

Only update the model when necessary:

```go
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case ProgressMsg:
        // Only update if changed
        if m.progress != msg.Percent {
            m.progress = msg.Percent
        }
    }
    return m, nil
}
```

### 2. Use Lazy Rendering

Don't render components that aren't visible:

```go
func (m Model) View() string {
    if m.hidden {
        return ""
    }
    // Render normally
}
```

## Resources

- [Bubble Tea Examples](https://github.com/charmbracelet/bubbletea/tree/master/examples)
- [Lipgloss Examples](https://github.com/charmbracelet/lipgloss/tree/master/examples)
- [Bubbles Components](https://github.com/charmbracelet/bubbles)
- [Charm Community](https://github.com/charmbracelet)

## Common Patterns

### Spinner + Progress

```go
type Model struct {
    spinner  spinner.Model
    progress int
}

func (m Model) View() string {
    return fmt.Sprintf("%s %s %d%%",
        m.spinner.View(),
        progressBar(m.progress),
        m.progress)
}
```

### Multi-Step Wizard

```go
type Step int

const (
    StepWelcome Step = iota
    StepDependencies
    StepInstall
    StepComplete
)

type Model struct {
    step Step
}

func (m Model) View() string {
    switch m.step {
    case StepWelcome:
        return renderWelcome()
    case StepDependencies:
        return renderDependencies()
    // ...
    }
}
```

---

Happy CLI Development! 🚀
