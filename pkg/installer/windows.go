package installer

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// WindowsInstaller handles Flutter installation on Windows
type WindowsInstaller struct {
	Config *InstallConfig
}

// NewWindowsInstaller creates a new Windows installer
func NewWindowsInstaller(config *InstallConfig) *WindowsInstaller {
	return &WindowsInstaller{Config: config}
}

// CheckDependencies checks if required dependencies are installed
func (w *WindowsInstaller) CheckDependencies() []Dependency {
	deps := []Dependency{
		w.checkGit(),
		w.checkJava(),
		w.checkAndroidSDK(),
		w.checkFlutter(),
	}
	return deps
}

func (w *WindowsInstaller) checkGit() Dependency {
	dep := Dependency{
		Name:        "Git",
		Description: "Version control system (required for Flutter)",
		Required:    true,
	}

	cmd := exec.Command("git", "--version")
	output, err := cmd.CombinedOutput()
	if err == nil {
		dep.IsInstalled = true
		dep.Version = strings.TrimSpace(string(output))
	}

	// Try to find git path
	gitPath, err := exec.LookPath("git")
	if err == nil {
		w.Config.GitPath = gitPath
	}

	return dep
}

func (w *WindowsInstaller) checkJava() Dependency {
	dep := Dependency{
		Name:        "Java JDK",
		Description: "Java Development Kit 17+ (required for Android development)",
		Required:    true,
	}

	cmd := exec.Command("java", "-version")
	output, err := cmd.CombinedOutput()
	if err == nil {
		dep.IsInstalled = true
		dep.Version = strings.TrimSpace(string(output))
	}

	// Try to find JAVA_HOME
	javaHome := os.Getenv("JAVA_HOME")
	if javaHome != "" {
		w.Config.JavaPath = javaHome
	}

	return dep
}

func (w *WindowsInstaller) checkAndroidSDK() Dependency {
	dep := Dependency{
		Name:        "Android SDK",
		Description: "Android command-line tools (required for Android development)",
		Required:    true,
	}

	// Check common locations
	possiblePaths := []string{
		os.Getenv("ANDROID_HOME"),
		os.Getenv("ANDROID_SDK_ROOT"),
		filepath.Join(os.Getenv("LOCALAPPDATA"), "Android", "Sdk"),
		filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Local", "Android", "Sdk"),
	}

	for _, path := range possiblePaths {
		if path != "" {
			if _, err := os.Stat(filepath.Join(path, "platform-tools")); err == nil {
				dep.IsInstalled = true
				w.Config.AndroidSDKPath = path

				// Try to get version
				adbPath := filepath.Join(path, "platform-tools", "adb.exe")
				if cmd := exec.Command(adbPath, "version"); cmd != nil {
					if output, err := cmd.CombinedOutput(); err == nil {
						dep.Version = strings.Split(string(output), "\n")[0]
					}
				}
				break
			}
		}
	}

	return dep
}

func (w *WindowsInstaller) checkFlutter() Dependency {
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

	// Try to find flutter path
	flutterPath, err := exec.LookPath("flutter")
	if err == nil {
		w.Config.FlutterPath = filepath.Dir(filepath.Dir(flutterPath))
	}

	return dep
}

// GetDefaultFlutterPath returns the default Flutter installation path
func (w *WindowsInstaller) GetDefaultFlutterPath() string {
	userProfile := os.Getenv("USERPROFILE")
	return filepath.Join(userProfile, "flutter")
}

// GetDefaultAndroidSDKPath returns the default Android SDK path
func (w *WindowsInstaller) GetDefaultAndroidSDKPath() string {
	localAppData := os.Getenv("LOCALAPPDATA")
	return filepath.Join(localAppData, "Android", "Sdk")
}

// DownloadFlutter downloads and extracts Flutter SDK
func (w *WindowsInstaller) DownloadFlutter(progressCallback func(percent int, status string)) error {
	// This is a placeholder - actual implementation would download from
	// https://storage.googleapis.com/flutter_infra_release/releases/stable/windows/flutter_windows_3.x.x-stable.zip
	progressCallback(0, "Preparing to download Flutter SDK...")

	// Check if path exists
	if _, err := os.Stat(w.Config.FlutterPath); os.IsNotExist(err) {
		if err := os.MkdirAll(w.Config.FlutterPath, 0755); err != nil {
			return fmt.Errorf("failed to create Flutter directory: %w", err)
		}
	}

	progressCallback(100, "Flutter download complete!")
	return nil
}

// SetupEnvironmentPath adds Flutter to PATH
func (w *WindowsInstaller) SetupEnvironmentPath() error {
	// This would use PowerShell to permanently add to PATH
	// For now, just a placeholder
	binPath := filepath.Join(w.Config.FlutterPath, "bin")
	fmt.Printf("Add to PATH: %s\n", binPath)
	return nil
}

// AcceptAndroidLicenses runs flutter doctor --android-licenses
func (w *WindowsInstaller) AcceptAndroidLicenses() error {
	cmd := exec.Command("flutter", "doctor", "--android-licenses")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// RunFlutterDoctor runs flutter doctor to verify installation
func (w *WindowsInstaller) RunFlutterDoctor() (string, error) {
	cmd := exec.Command("flutter", "doctor", "-v")
	output, err := cmd.CombinedOutput()
	return string(output), err
}
