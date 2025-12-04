# Extension Examples

This document shows how to extend the Flutter Installation Tool with new features.

## Example 1: Adding macOS Support

### Step 1: Create macOS Installer

Create `pkg/installer/macos.go`:

```go
package installer

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type MacOSInstaller struct {
	Config *InstallConfig
}

func NewMacOSInstaller(config *InstallConfig) *MacOSInstaller {
	return &MacOSInstaller{Config: config}
}

func (m *MacOSInstaller) CheckDependencies() []Dependency {
	return []Dependency{
		m.checkGit(),
		m.checkXcode(),
		m.checkCocoaPods(),
		m.checkFlutter(),
	}
}

func (m *MacOSInstaller) checkGit() Dependency {
	dep := Dependency{
		Name:        "Git",
		Description: "Version control system",
		Required:    true,
	}

	cmd := exec.Command("git", "--version")
	output, err := cmd.CombinedOutput()
	if err == nil {
		dep.IsInstalled = true
		dep.Version = strings.TrimSpace(string(output))
	}
	return dep
}

func (m *MacOSInstaller) checkXcode() Dependency {
	dep := Dependency{
		Name:        "Xcode",
		Description: "Required for iOS development",
		Required:    true,
	}

	// Check if Xcode is installed
	cmd := exec.Command("xcode-select", "-p")
	_, err := cmd.CombinedOutput()
	dep.IsInstalled = (err == nil)

	// Get Xcode version
	if dep.IsInstalled {
		cmd := exec.Command("xcodebuild", "-version")
		if output, err := cmd.CombinedOutput(); err == nil {
			lines := strings.Split(string(output), "\n")
			if len(lines) > 0 {
				dep.Version = strings.TrimSpace(lines[0])
			}
		}
	}
	return dep
}

func (m *MacOSInstaller) checkCocoaPods() Dependency {
	dep := Dependency{
		Name:        "CocoaPods",
		Description: "Dependency manager for iOS",
		Required:    true,
	}

	cmd := exec.Command("pod", "--version")
	output, err := cmd.CombinedOutput()
	if err == nil {
		dep.IsInstalled = true
		dep.Version = strings.TrimSpace(string(output))
	}
	return dep
}

func (m *MacOSInstaller) checkFlutter() Dependency {
	dep := Dependency{
		Name:        "Flutter SDK",
		Description: "Flutter development framework",
		Required:    false,
	}

	cmd := exec.Command("flutter", "--version")
	output, err := cmd.CombinedOutput()
	if err == nil {
		dep.IsInstalled = true
		lines := strings.Split(string(output), "\n")
		if len(lines) > 0 {
			dep.Version = strings.TrimSpace(lines[0])
		}
	}
	return dep
}

func (m *MacOSInstaller) GetDefaultFlutterPath() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, "flutter")
}

func (m *MacOSInstaller) InstallCocoaPods() error {
	cmd := exec.Command("sudo", "gem", "install", "cocoapods")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
```

### Step 2: Update Main Application

In `main.go`, add platform detection:

```go
import "runtime"

func detectPlatform() installer.Platform {
	switch runtime.GOOS {
	case "windows":
		return installer.PlatformWindows
	case "darwin":
		return installer.PlatformMacOS
	case "linux":
		return installer.PlatformLinux
	default:
		return installer.PlatformWindows
	}
}

func createInstaller(platform installer.Platform, config *installer.InstallConfig) interface{} {
	switch platform {
	case installer.PlatformWindows:
		return installer.NewWindowsInstaller(config)
	case installer.PlatformMacOS:
		return installer.NewMacOSInstaller(config)
	case installer.PlatformLinux:
		return installer.NewLinuxInstaller(config)
	default:
		return installer.NewWindowsInstaller(config)
	}
}
```

## Example 2: Adding iOS Development Support

### Update Types

In `pkg/installer/types.go`, add iOS-specific configuration:

```go
type InstallConfig struct {
	FlutterPath    string
	AndroidSDKPath string
	XcodePath      string
	CocoaPodsPath  string
	GitPath        string
	JavaPath       string
	Platform       Platform
	Target         TargetPlatform
	EnableIOS      bool  // New field
}

// Add iOS-specific dependencies
func (w *WindowsInstaller) CheckIOSDependencies() []Dependency {
	// On Windows, iOS development is not supported
	return []Dependency{
		{
			Name:        "iOS Development",
			Description: "iOS development is only available on macOS",
			IsInstalled: false,
			Required:    false,
		},
	}
}
```

## Example 3: Adding Configuration File Support

### Create Config Type

Create `pkg/config/config.go`:

```go
package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	FlutterPath       string   `json:"flutter_path"`
	AndroidSDKPath    string   `json:"android_sdk_path"`
	PreferredChannel  string   `json:"preferred_channel"`
	AutoUpdate        bool     `json:"auto_update"`
	DefaultTargets    []string `json:"default_targets"`
}

func LoadConfig() (*Config, error) {
	homeDir, _ := os.UserHomeDir()
	configPath := filepath.Join(homeDir, ".flutter-installer", "config.json")

	data, err := os.ReadFile(configPath)
	if err != nil {
		return DefaultConfig(), nil
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (c *Config) Save() error {
	homeDir, _ := os.UserHomeDir()
	configDir := filepath.Join(homeDir, ".flutter-installer")
	os.MkdirAll(configDir, 0755)

	configPath := filepath.Join(configDir, "config.json")
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

func DefaultConfig() *Config {
	homeDir, _ := os.UserHomeDir()
	return &Config{
		FlutterPath:      filepath.Join(homeDir, "flutter"),
		PreferredChannel: "stable",
		AutoUpdate:       false,
		DefaultTargets:   []string{"android"},
	}
}
```

## Example 4: Adding Download Progress

### Create Downloader

Create `pkg/downloader/downloader.go`:

```go
package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type ProgressCallback func(downloaded, total int64, percent float64)

func DownloadFile(url, filepath string, callback ProgressCallback) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create progress reader
	total := resp.ContentLength
	downloaded := int64(0)

	buffer := make([]byte, 32*1024) // 32KB buffer
	for {
		n, err := resp.Body.Read(buffer)
		if n > 0 {
			_, writeErr := out.Write(buffer[:n])
			if writeErr != nil {
				return writeErr
			}

			downloaded += int64(n)
			percent := float64(downloaded) / float64(total) * 100

			if callback != nil {
				callback(downloaded, total, percent)
			}
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}

	return nil
}
```

### Usage in Main

```go
import "flutter_install_tool/pkg/downloader"

func downloadFlutter() error {
	url := "https://storage.googleapis.com/flutter_infra_release/releases/stable/windows/flutter_windows_3.x.x-stable.zip"
	destination := "flutter.zip"

	callback := func(downloaded, total int64, percent float64) {
		fmt.Printf("\rDownloading: %.2f%% (%d/%d bytes)", 
			percent, downloaded, total)
	}

	return downloader.DownloadFile(url, destination, callback)
}
```

## Example 5: Adding Multi-Language Support

### Create Localization

Create `pkg/i18n/i18n.go`:

```go
package i18n

type Language string

const (
	English    Language = "en"
	Spanish    Language = "es"
	French     Language = "fr"
	German     Language = "de"
	Japanese   Language = "ja"
	Chinese    Language = "zh"
)

var translations = map[Language]map[string]string{
	English: {
		"welcome":            "Welcome to Flutter Installation Tool",
		"check_deps":         "Check Dependencies",
		"install_flutter":    "Install Flutter SDK",
		"quit":              "Exit",
	},
	Spanish: {
		"welcome":            "Bienvenido a la Herramienta de Instalación de Flutter",
		"check_deps":         "Verificar Dependencias",
		"install_flutter":    "Instalar Flutter SDK",
		"quit":              "Salir",
	},
	// Add more languages...
}

type Translator struct {
	lang Language
}

func NewTranslator(lang Language) *Translator {
	return &Translator{lang: lang}
}

func (t *Translator) T(key string) string {
	if langMap, ok := translations[t.lang]; ok {
		if text, ok := langMap[key]; ok {
			return text
		}
	}
	// Fallback to English
	if text, ok := translations[English][key]; ok {
		return text
	}
	return key
}
```

## Example 6: Adding Plugin System

### Plugin Interface

Create `pkg/plugins/plugin.go`:

```go
package plugins

type Plugin interface {
	Name() string
	Description() string
	PreInstall() error
	PostInstall() error
	Validate() error
}

type PluginManager struct {
	plugins []Plugin
}

func NewPluginManager() *PluginManager {
	return &PluginManager{
		plugins: make([]Plugin, 0),
	}
}

func (pm *PluginManager) Register(p Plugin) {
	pm.plugins = append(pm.plugins, p)
}

func (pm *PluginManager) RunPreInstall() error {
	for _, p := range pm.plugins {
		if err := p.PreInstall(); err != nil {
			return fmt.Errorf("plugin %s pre-install failed: %w", p.Name(), err)
		}
	}
	return nil
}

func (pm *PluginManager) RunPostInstall() error {
	for _, p := range pm.plugins {
		if err := p.PostInstall(); err != nil {
			return fmt.Errorf("plugin %s post-install failed: %w", p.Name(), err)
		}
	}
	return nil
}
```

### Example Plugin

```go
type VSCodePlugin struct{}

func (p *VSCodePlugin) Name() string {
	return "VS Code Flutter Extension"
}

func (p *VSCodePlugin) Description() string {
	return "Installs Flutter extension for VS Code"
}

func (p *VSCodePlugin) PreInstall() error {
	// Check if VS Code is installed
	return nil
}

func (p *VSCodePlugin) PostInstall() error {
	// Install Flutter extension
	cmd := exec.Command("code", "--install-extension", "Dart-Code.flutter")
	return cmd.Run()
}

func (p *VSCodePlugin) Validate() error {
	// Verify extension is installed
	return nil
}
```

## Example 7: Adding Update Checker

Create `pkg/updater/updater.go`:

```go
package updater

import (
	"encoding/json"
	"net/http"
	"time"
)

type ReleaseInfo struct {
	Version     string    `json:"version"`
	ReleaseDate time.Time `json:"release_date"`
	DownloadURL string    `json:"download_url"`
	Changelog   string    `json:"changelog"`
}

func CheckForUpdates(currentVersion string) (*ReleaseInfo, bool, error) {
	resp, err := http.Get("https://api.example.com/flutter-installer/latest")
	if err != nil {
		return nil, false, err
	}
	defer resp.Body.Close()

	var release ReleaseInfo
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, false, err
	}

	hasUpdate := release.Version != currentVersion
	return &release, hasUpdate, nil
}

func DownloadUpdate(url, destination string) error {
	// Use the downloader from Example 4
	return nil
}
```

## Compile for Different Platforms

### Build Script

Create `build.sh`:

```bash
#!/bin/bash

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o bin/flutter-installer-windows.exe .

# Build for macOS
GOOS=darwin GOARCH=amd64 go build -o bin/flutter-installer-macos .

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o bin/flutter-installer-linux .

echo "Build complete!"
```

### PowerShell Build Script

Create `build.ps1`:

```powershell
# Build for Windows
$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o bin/flutter-installer-windows.exe .

# Build for macOS
$env:GOOS="darwin"; $env:GOARCH="amd64"; go build -o bin/flutter-installer-macos .

# Build for Linux
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o bin/flutter-installer-linux .

Write-Host "Build complete!"
```

---

These examples demonstrate the extensibility of the Flutter Installation Tool. Feel free to implement any of these features or create your own! 🚀
