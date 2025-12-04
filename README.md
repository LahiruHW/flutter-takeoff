<div align="center">
  <img src="assets/logo.svg" alt="Flutter Takeoff Logo" width="200"/>
  
  # Flutter Takeoff
  
  A beautiful, interactive CLI tool built with Go that simplifies Flutter SDK installation
</div>

Currently only supported on Windows with Android support.

> [!CAUTION]
> ### **The stable version of this project is still in an experimental phase - please use it at your own risk**


## ✨ Features

- **Beautiful Terminal UI** - Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lipgloss](https://github.com/charmbracelet/lipgloss) for an elegant user experience
- **Dependency Checking** - Automatically detects installed prerequisites (Git, Java JDK, Android SDK)
- **Interactive Menus** - Dropdown selections minimize manual typing
- **Custom Installation Paths** - Choose where to install Flutter SDK
- **Progress Indicators** - Visual feedback during installation
- **Extensible Architecture** - Designed to support additional platforms (macOS, Linux, iOS) in the future

## 🎯 Current Scope

- **Platform**: Windows only
- **Target**: Android development
- **Prerequisites Detection**: Git, Java JDK, Android SDK, Flutter SDK

## 📋 Prerequisites

Before running the Flutter Installation Tool, ensure you have:

1. **Git for Windows** - [Download here](https://git-scm.com/download/win)
2. **Java JDK 17+** - [Download Temurin](https://adoptium.net/)
3. **Android SDK** - Either install Android Studio or use command-line tools

## 🚀 Quick Start

### Download & Run

1. Download the latest release from [Releases](https://github.com/LahiruHW/flutter-takeoff/releases)

   **OR** build from source:

   ```powershell
   # Using build script (recommended - injects version info)
   .\build.ps1

   # Or simple build
   go build -o flutter-installer.exe .
   ```

2. Run the installer:

   ```powershell
   .\flutter-installer.exe
   ```

3. Follow the interactive prompts!

## 📖 Usage

The tool provides an interactive menu with the following options:

### 1. Check Dependencies

Scans your system for required software and displays installation status:

- ✓ Git
- ✓ Java JDK
- ✓ Android SDK
- ✓ Flutter SDK (if already installed)

### 2. Install Flutter SDK

Guides you through Flutter installation:

- Choose custom installation path or use default (`%USERPROFILE%\flutter`)
- Downloads Flutter SDK
- Extracts files
- Adds Flutter to PATH
- Provides next steps for Android license acceptance

### 3. Run Flutter Doctor

Executes `flutter doctor -v` to diagnose your Flutter installation and identify any issues.

### 4. Version Info

Displays detailed version and build information:

- Semantic version number
- Build date and time
- Git commit hash
- Git branch name
- Links to repository and issue tracker

### 5. Exit

Safely exits the application.

## 🏗️ Project Structure

```
flutter_takeoff/
├── main.go                     # Main application entry point
├── go.mod                      # Go module definition
├── go.sum                      # Dependency checksums
├── build.ps1                   # PowerShell build script
├── build.sh                    # Bash build script
├── CHANGELOG.md                # Version history
├── pkg/
│   ├── installer/
│   │   ├── types.go           # Core data types and interfaces
│   │   └── windows.go         # Windows-specific installation logic
│   ├── ui/
│   │   ├── styles.go          # Color schemes and styling
│   │   ├── menu.go            # Interactive menu components
│   │   └── progress.go        # Progress bars and spinners
│   └── version/
│       └── version.go         # Version management
└── README.md                   # This file
```

## 🎨 UI Components

The tool uses several beautiful UI components:

- **Styled Text**: Color-coded status messages (success, error, warning, info)
- **Interactive Menus**: Keyboard-navigable dropdown lists
- **Progress Bars**: Visual installation progress
- **Spinners**: Animated loading indicators
- **Checkboxes**: Task completion visualization

## 🔧 Architecture Highlights

### Extensibility

The project is designed with extensibility in mind:

```go
// Easy to add new platforms
type Platform string
const (
    PlatformWindows Platform = "windows"
    PlatformMacOS   Platform = "macos"    // Ready for future
    PlatformLinux   Platform = "linux"    // Ready for future
)

// Support multiple development targets
type TargetPlatform string
const (
    TargetAndroid TargetPlatform = "android"
    TargetIOS     TargetPlatform = "ios"      // Ready for future
    TargetWeb     TargetPlatform = "web"      // Ready for future
    TargetDesktop TargetPlatform = "desktop"  // Ready for future
)
```

### Key Interfaces

- `WindowsInstaller` - Handles Windows-specific operations
- `InstallConfig` - Centralized configuration management
- `Dependency` - Represents checkable prerequisites

## 📚 Learning Resources

This project demonstrates several Go best practices and CLI development techniques:

### Libraries Used

- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)** - Terminal UI framework based on The Elm Architecture
- **[Lipgloss](https://github.com/charmbracelet/lipgloss)** - Style definitions for terminal output
- **[Bubbles](https://github.com/charmbracelet/bubbles)** - Common UI components (lists, spinners, progress bars)

### Patterns Demonstrated

1. **MVC-like Architecture** - Separation of UI, business logic, and data
2. **Command Pattern** - Menu-driven actions
3. **Strategy Pattern** - Platform-specific installers
4. **Builder Pattern** - Configuration construction

## 🛠️ Development

### Building from Source

```powershell
# Clone the repository
git clone <your-repo-url>
cd flutter_install_tool

# Install dependencies
go mod download

# Build with version information (recommended)
.\build.ps1

# Or on Linux/macOS
chmod +x build.sh
./build.sh

# Simple build (without version info)
go build -o flutter-installer.exe .

# Run
.\flutter-installer.exe
```

### Version Management

The project uses semantic versioning. To release a new version:

1. Update version in `pkg/version/version.go`:

   ```go
   const (
       Major      = 1
       Minor      = 1
       Patch      = 0
       PreRelease = "" // or "beta", "rc.1", etc.
   )
   ```

2. Update `CHANGELOG.md` with changes

3. Build release:

   ```powershell
   .\build.ps1 -OutputName flutter-installer-v1.1.0
   ```

4. Tag and push:
   ```bash
   git tag -a v1.1.0 -m "Release v1.1.0"
   git push origin v1.1.0
   ```

### Adding New Platforms

To add support for macOS or Linux:

1. Create a new installer file (e.g., `pkg/installer/macos.go`)
2. Implement platform-specific dependency checking
3. Add platform detection logic in `main.go`
4. Update the configuration to support the new platform

Example:

```go
type MacOSInstaller struct {
    Config *InstallConfig
}

func (m *MacOSInstaller) CheckDependencies() []Dependency {
    // macOS-specific implementation
}
```

### Adding New Target Platforms

To add iOS, web, or desktop support:

1. Update dependency checks for the new target
2. Add target-specific installation steps
3. Modify the UI to allow target selection

## 🎯 Future Enhancements

- [ ] macOS support
- [ ] Linux support
- [ ] iOS development setup (for macOS)
- [ ] Web development setup
- [ ] Desktop development setup (Windows/Linux/macOS)
- [ ] Automatic Git installation
- [ ] Automatic Java JDK installation
- [ ] Automatic Android SDK installation
- [ ] Flutter version selection (stable/beta/dev)
- [ ] VS Code extension installation
- [ ] Android Studio plugin installation
- [ ] Uninstallation feature
- [ ] Update Flutter feature
- [ ] Configuration file support

## 🤝 Contributing

Contributions are welcome! Here are some ways you can help:

- Report bugs
- Suggest new features
- Improve documentation
- Submit pull requests

## 📝 License

This project is open source and available under the MIT License.

## 🙏 Acknowledgments

- [Flutter Team](https://flutter.dev/) for the amazing framework
- [Charm](https://charm.sh/) for the beautiful CLI libraries
- The Go community for excellent tooling and support

## 📞 Support

If you encounter any issues or have questions:

1. Check the [Flutter documentation](https://docs.flutter.dev/)
2. Review the dependency installation guides linked in the tool
3. Run `flutter doctor -v` for detailed diagnostics

## 🎓 Learning More

This tool was built as a learning project to demonstrate:

- Modern CLI application development in Go
- Beautiful terminal user interfaces
- System integration and automation
- Software installation best practices

### Recommended Reading

- [Bubble Tea Tutorial](https://github.com/charmbracelet/bubbletea/tree/master/tutorials)
- [Flutter Installation Guide](https://docs.flutter.dev/get-started/install)
- [Android Development Setup](https://docs.flutter.dev/platform-integration/android/setup)

---

**Made with ❤️ and Go by @LahiruHW** | Happy Flutter Development! 🎯
