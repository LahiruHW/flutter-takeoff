# Changelog

All notable changes to Flutter Takeoff will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Planned
- macOS support
- Linux support
- iOS development setup
- Automatic dependency installation
- Configuration file support
- Update checker and auto-update feature

## [1.0.0] - 2025-12-04

### Added
- Initial release
- Beautiful terminal UI using Bubble Tea and Lipgloss
- Interactive menus with keyboard navigation
- Dependency checking for Windows + Android development
  - Git detection
  - Java JDK detection
  - Android SDK detection
  - Flutter SDK detection
- Guided Flutter SDK installation
- Custom installation path selection
- Progress indicators and visual feedback
- Flutter Doctor integration
- Version information display
- Build system with version injection
- Comprehensive documentation
  - README with features and usage
  - DEVELOPMENT guide for CLI best practices
  - FLUTTER_GUIDE for manual installation
  - EXAMPLES for extending the tool
  - QUICKREF for Charm libraries reference

### Features
- Windows platform support
- Android target platform support
- Extensible architecture for future platforms
- Color-coded status messages
- Minimal user typing required
- Interactive dropdown selections

### Technical
- Go module structure
- Package-based organization
- Platform-specific installers
- UI component library
- Version management system
- Build scripts for PowerShell and Bash

## Version Format

Versions follow Semantic Versioning: `MAJOR.MINOR.PATCH[-PRERELEASE]`

- **MAJOR**: Incompatible API changes
- **MINOR**: New functionality (backwards compatible)
- **PATCH**: Bug fixes (backwards compatible)
- **PRERELEASE**: Optional pre-release identifier (e.g., `beta`, `rc.1`)

## Release Process

1. Update version in `pkg/version/version.go`
2. Update CHANGELOG.md with changes
3. Build release binaries:
   ```powershell
   .\build.ps1 -OutputName flutter-installer-v1.0.0
   ```
4. Tag the release:
   ```bash
   git tag -a v1.0.0 -m "Release v1.0.0"
   git push origin v1.0.0
   ```
5. Create GitHub release with binaries

## Links

- [Repository](https://github.com/LahiruHW/flutter-takeoff)
- [Issues](https://github.com/LahiruHW/flutter-takeoff/issues)
- [Releases](https://github.com/LahiruHW/flutter-takeoff/releases)
