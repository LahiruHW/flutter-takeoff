package installer

// Platform represents the target operating system
type Platform string

const (
	PlatformWindows Platform = "windows"
	PlatformMacOS   Platform = "macos"
	PlatformLinux   Platform = "linux"
)

// TargetPlatform represents the mobile/desktop platform to develop for
type TargetPlatform string

const (
	TargetAndroid TargetPlatform = "android"
	TargetIOS     TargetPlatform = "ios"
	TargetWeb     TargetPlatform = "web"
	TargetDesktop TargetPlatform = "desktop"
)

// Dependency represents a required software dependency
type Dependency struct {
	Name        string
	Description string
	IsInstalled bool
	Version     string
	Required    bool
}

// InstallConfig holds configuration for the installation
type InstallConfig struct {
	FlutterPath    string
	AndroidSDKPath string
	GitPath        string
	JavaPath       string
	Platform       Platform
	Target         TargetPlatform
}
