package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"

	"flutter_takeoff/pkg/installer"
	"flutter_takeoff/pkg/ui"
	"flutter_takeoff/pkg/version"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Print welcome banner
	printBanner()

	// Detect platform
	config := &installer.InstallConfig{
		Platform: installer.PlatformWindows,
		Target:   installer.TargetAndroid,
	}

	// Create installer
	windowsInstaller := installer.NewWindowsInstaller(config)

	// Main menu
	for {
		choice := showMainMenu()

		switch choice {
		case "check":
			checkDependencies(windowsInstaller)
		case "install":
			runInstallation(windowsInstaller)
		case "doctor":
			runFlutterDoctor(windowsInstaller)
		case "version":
			showVersionInfo()
		case "quit":
			fmt.Println(ui.SuccessStyle.Render("\nThank you for using Flutter Takeoff! 👋\n"))
			return
		}
	}
}

func printBanner() {
	banner := `
				╔═══════════════════════════════════════════════════════╗
				║                                                       ║
				║              🚀 Flutter Takeoff 🚀                    ║
				║                                                       ║
				║        Simplify Your Flutter Development Setup        ║
				║                                                       ║
				╚═══════════════════════════════════════════════════════╝
`
	fmt.Println(ui.TitleStyle.Render(banner))

	// Show version with detected target platforms
	targetPlatforms := getTargetPlatforms()
	fmt.Printf("%s\n", ui.SubtleStyle.Render("  v"+version.FullVersion()+" | "+targetPlatforms+"\n"))

	// Detect and display platform information
	platformInfo := detectPlatform()
	fmt.Printf("%s\n\n", ui.SubtleStyle.Render("  "+platformInfo))
}

// returns available target platforms based on OS
func getTargetPlatforms() string {
	osName := runtime.GOOS

	var platforms []string

	switch osName {
	case "windows":
		platforms = []string{"Android", "Web", "Windows Desktop"}
	case "darwin":
		platforms = []string{"iOS", "Android", "Web", "macOS Desktop"}
	case "linux":
		platforms = []string{"Android", "Web", "Linux Desktop"}
	default:
		platforms = []string{"Web"}
	}

	return strings.Join(platforms, " + ")
}

// detects the current OS and architecture
func detectPlatform() string {
	osName := runtime.GOOS
	arch := runtime.GOARCH

	// Get friendly OS name
	var osDisplay string
	var osIcon string

	switch osName {
	case "windows":
		osDisplay = "Windows"
		osIcon = "🪟"
	case "darwin":
		osDisplay = "macOS"
		osIcon = "🍎"
	case "linux":
		osDisplay = "Linux"
		osIcon = "🐧"
	default:
		osDisplay = osName
		osIcon = "💻"
	}

	// Get friendly architecture name
	var archDisplay string
	switch arch {
	case "amd64":
		archDisplay = "64-bit"
	case "386":
		archDisplay = "32-bit"
	case "arm64":
		archDisplay = "ARM64"
	case "arm":
		archDisplay = "ARM"
	default:
		archDisplay = arch
	}

	// Check if running in supported environment
	var supportStatus string
	switch osName {
	case "windows":
		supportStatus = ui.SuccessStyle.Render("✓ Fully Supported")
	case "darwin", "linux":
		supportStatus = ui.WarningStyle.Render("⚠ Experimental")
	default:
		supportStatus = ui.ErrorStyle.Render("✗ Unsupported")
	}

	return fmt.Sprintf("Platform: %s \u00A0 %s %s | %s",
		osIcon, osDisplay, archDisplay, supportStatus)
}

// shows the main menu and returns the selected choice
func showMainMenu() string {
	items := []ui.MenuItem{
		{Title: "Check Dependencies", Description: "Verify installed prerequisites", Value: "check"},
		{Title: "Install Flutter SDK", Description: "Download and set up Flutter", Value: "install"},
		{Title: "Run Flutter Doctor", Description: "Diagnose Flutter installation", Value: "doctor"},
		{Title: "Version Info", Description: "Show version and build information", Value: "version"},
		{Title: "Exit", Description: "Quit the installer", Value: "quit"},
	}

	m := ui.NewMenu("What would you like to do?", items, 50, 15)
	p := tea.NewProgram(m)

	finalModel, err := p.Run()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if menuModel, ok := finalModel.(ui.MenuModel); ok {
		return menuModel.Choice()
	}

	return "quit"
}

func checkDependencies(inst *installer.WindowsInstaller) {
	fmt.Println(ui.Header("Checking Dependencies"))

	deps := inst.CheckDependencies()

	fmt.Println(ui.HeaderStyle.Render("Required Dependencies:\n"))

	for _, dep := range deps {
		status := "error"
		statusText := "Not installed"

		if dep.IsInstalled {
			status = "success"
			statusText = "Installed"
			if dep.Version != "" {
				statusText += " (" + dep.Version + ")"
			}
		}

		fmt.Printf("%s %s\n",
			ui.StatusIndicator(status, dep.Name+":"),
			ui.SubtleStyle.Render(statusText))

		if !dep.IsInstalled && dep.Required {
			fmt.Println(ui.SubtleStyle.Render("  → " + dep.Description))
		}
	}

	fmt.Println()

	// Check if all required dependencies are installed
	allInstalled := true
	for _, dep := range deps {
		if dep.Required && !dep.IsInstalled {
			allInstalled = false
			break
		}
	}

	if allInstalled {
		fmt.Println(ui.SuccessStyle.Render("✓ All required dependencies are installed!\n"))
	} else {
		fmt.Println(ui.WarningStyle.Render("⚠ Some dependencies are missing\n"))
		fmt.Println(ui.SubtleStyle.Render("Installation Guide:"))
		fmt.Println(ui.SubtleStyle.Render("  • Git: https://git-scm.com/download/win"))
		fmt.Println(ui.SubtleStyle.Render("  • Java JDK: https://adoptium.net/"))
		fmt.Println(ui.SubtleStyle.Render("  • Android SDK: Install Android Studio or use command-line tools\n"))
	}

	waitForEnter()
}

func runInstallation(inst *installer.WindowsInstaller) {
	fmt.Println(ui.Header("Flutter SDK Installation"))

	// Check if Flutter is already installed
	deps := inst.CheckDependencies()
	var flutterDep installer.Dependency
	for _, dep := range deps {
		if dep.Name == "Flutter SDK" {
			flutterDep = dep
			break
		}
	}

	if flutterDep.IsInstalled {
		fmt.Println(ui.WarningStyle.Render("⚠ Flutter is already installed!"))
		fmt.Println(ui.SubtleStyle.Render("  Version: " + flutterDep.Version))
		fmt.Printf("\n%s ", ui.ConfirmPrompt("Do you want to reinstall?"))

		reader := bufio.NewReader(os.Stdin)
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(strings.ToLower(response))

		if response != "y" && response != "yes" {
			fmt.Println(ui.SubtleStyle.Render("\nInstallation cancelled.\n"))
			return
		}
	}

	// Get installation path
	defaultPath := inst.GetDefaultFlutterPath()
	fmt.Printf("\n%s\n", ui.NormalStyle.Render("Flutter installation path:"))
	fmt.Printf("%s %s\n", ui.SubtleStyle.Render("Default:"), defaultPath)
	fmt.Printf("%s ", ui.SubtleStyle.Render("Press Enter to use default, or type a custom path:"))

	reader := bufio.NewReader(os.Stdin)
	customPath, _ := reader.ReadString('\n')
	customPath = strings.TrimSpace(customPath)

	if customPath != "" {
		inst.Config.FlutterPath = customPath
	} else {
		inst.Config.FlutterPath = defaultPath
	}

	fmt.Printf("\n%s %s\n\n",
		ui.SuccessStyle.Render("✓ Installation path set to:"),
		inst.Config.FlutterPath)

	// Installation steps menu
	fmt.Println(ui.HeaderStyle.Render("Installation Steps:\n"))
	fmt.Println(ui.Checkbox(true, "Download Flutter SDK"))
	fmt.Println(ui.Checkbox(true, "Extract to installation path"))
	fmt.Println(ui.Checkbox(true, "Add Flutter to PATH"))
	fmt.Println(ui.Checkbox(true, "Accept Android licenses"))
	fmt.Println(ui.Checkbox(true, "Run flutter doctor"))

	fmt.Printf("\n%s ", ui.ConfirmPrompt("Continue with installation?"))
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(strings.ToLower(response))

	if response != "y" && response != "yes" {
		fmt.Println(ui.SubtleStyle.Render("\nInstallation cancelled.\n"))
		return
	}

	// Simulate installation process
	fmt.Println(ui.Header("Installing Flutter SDK"))

	steps := []struct {
		name    string
		percent int
	}{
		{"Preparing installation...", 10},
		{"Downloading Flutter SDK (this may take a few minutes)...", 30},
		{"Extracting files...", 60},
		{"Setting up environment...", 80},
		{"Configuring PATH...", 90},
		{"Finalizing installation...", 100},
	}

	for _, step := range steps {
		fmt.Printf("\r%s %s",
			ui.SimpleProgressBar(step.percent, 40),
			ui.SubtleStyle.Render(step.name))
		// In real implementation, this would actually download and install
	}

	fmt.Println()
	fmt.Println(ui.SuccessStyle.Render("✓ Flutter SDK installation complete!\n"))

	fmt.Println(ui.WarningStyle.Render("⚠ Important Next Steps:\n"))
	fmt.Println(ui.SubtleStyle.Render("1. Restart your terminal/command prompt"))
	fmt.Println(ui.SubtleStyle.Render("2. Run 'flutter doctor' to verify installation"))
	fmt.Println(ui.SubtleStyle.Render("3. Accept Android licenses with 'flutter doctor --android-licenses'\n"))

	waitForEnter()
}

func runFlutterDoctor(inst *installer.WindowsInstaller) {
	fmt.Println(ui.Header("Running Flutter Doctor"))

	output, err := inst.RunFlutterDoctor()
	if err != nil {
		fmt.Println(ui.ErrorStyle.Render("✗ Failed to run flutter doctor"))
		fmt.Println(ui.SubtleStyle.Render("  Error: " + err.Error()))
		fmt.Println(ui.SubtleStyle.Render("\n  Make sure Flutter is installed and added to PATH\n"))
	} else {
		fmt.Println(output)
	}

	waitForEnter()
}

func showVersionInfo() {
	fmt.Println(ui.Header("Version Information"))

	info := version.BuildInfo()

	fmt.Printf("%s %s\n",
		ui.SuccessStyle.Render("Version:"),
		info["version"])

	fmt.Printf("%s %s\n",
		ui.NormalStyle.Render("Build Date:"),
		ui.SubtleStyle.Render(info["build_date"]))

	fmt.Printf("%s %s\n",
		ui.NormalStyle.Render("Git Commit:"),
		ui.SubtleStyle.Render(info["git_commit"]))

	fmt.Printf("%s %s\n",
		ui.NormalStyle.Render("Git Branch:"),
		ui.SubtleStyle.Render(info["git_branch"]))

	if version.IsPreRelease() {
		fmt.Printf("\n%s\n",
			ui.WarningStyle.Render("⚠ This is a pre-release version"))
	}

	fmt.Println()
	fmt.Println(ui.SubtleStyle.Render("GitHub: https://github.com/LahiruHW/flutter-takeoff"))
	fmt.Println(ui.SubtleStyle.Render("Report issues: https://github.com/LahiruHW/flutter-takeoff/issues"))
	fmt.Println()

	waitForEnter()
}

func waitForEnter() {
	fmt.Println(ui.SubtleStyle.Render("Press Enter to continue..."))
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
