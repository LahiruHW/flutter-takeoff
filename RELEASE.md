# Release Guide

This document describes how to create and publish new releases of Flutter Takeoff.

## Version Numbering

We follow [Semantic Versioning 2.0.0](https://semver.org/):

```
MAJOR.MINOR.PATCH[-PRERELEASE]
```

- **MAJOR**: Incompatible API/behavior changes
- **MINOR**: New features (backwards compatible)
- **PATCH**: Bug fixes (backwards compatible)
- **PRERELEASE**: Pre-release versions (e.g., `beta`, `rc.1`)

### Examples

- `1.0.0` - Initial stable release
- `1.1.0` - Added macOS support
- `1.1.1` - Fixed dependency detection bug
- `2.0.0` - Complete rewrite (breaking changes)
- `1.2.0-beta` - Beta version of 1.2.0
- `1.2.0-rc.1` - Release candidate 1 for 1.2.0

## Release Process

### 1. Prepare the Release

#### Update Version Number

Edit `pkg/version/version.go`:

```go
const (
    Major      = 1
    Minor      = 1  // Increment for new features
    Patch      = 0  // Increment for bug fixes
    PreRelease = "" // Leave empty for stable, or "beta", "rc.1", etc.
)
```

#### Update CHANGELOG.md

Add a new section for the version:

```markdown
## [1.1.0] - 2025-12-05

### Added
- macOS support
- iOS development setup
- Configuration file support

### Fixed
- Dependency detection on some systems
- Progress bar display issues

### Changed
- Improved error messages
- Updated UI colors
```

#### Update Documentation

- Update README.md if features changed
- Update relevant guide files
- Update screenshots if UI changed

### 2. Test the Release

```powershell
# Build the release
.\build.ps1 -OutputName flutter-installer-test

# Test all features
.\flutter-installer-test.exe

# Test on different Windows versions if possible
# Windows 10, Windows 11, Windows Server
```

### 3. Build Release Binaries

```powershell
# Windows build
.\build.ps1 -OutputName flutter-installer-v1.1.0-windows-amd64

# If supporting multiple platforms:
# macOS build
$env:GOOS="darwin"; $env:GOARCH="amd64"; go build -o flutter-installer-v1.1.0-macos-amd64 .

# Linux build
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o flutter-installer-v1.1.0-linux-amd64 .
```

### 4. Create Git Tag

```bash
# Create annotated tag
git tag -a v1.1.0 -m "Release v1.1.0: Added macOS support"

# Verify tag
git tag -l -n1 v1.1.0

# Push tag to remote
git push origin v1.1.0
```

### 5. Create GitHub Release

1. Go to: https://github.com/yourusername/flutter-install-tool/releases
2. Click "Draft a new release"
3. Select the tag: `v1.1.0`
4. Release title: `v1.1.0 - macOS Support`
5. Description: Copy from CHANGELOG.md
6. Attach binaries:
   - `flutter-installer-v1.1.0-windows-amd64.exe`
   - `flutter-installer-v1.1.0-macos-amd64` (if available)
   - `flutter-installer-v1.1.0-linux-amd64` (if available)
7. Check "Set as the latest release" for stable versions
8. Check "This is a pre-release" for beta/rc versions
9. Click "Publish release"

### 6. Post-Release

1. Announce the release:
   - Update project homepage
   - Post on social media
   - Notify users/community

2. Monitor for issues:
   - Watch GitHub issues
   - Check user feedback
   - Prepare hotfix if critical bugs found

## Hotfix Releases

For critical bugs in production:

1. Create hotfix branch from tag:
   ```bash
   git checkout -b hotfix/v1.1.1 v1.1.0
   ```

2. Fix the bug and update:
   - Version: `1.1.0` → `1.1.1`
   - CHANGELOG.md

3. Test thoroughly

4. Merge to main and tag:
   ```bash
   git checkout main
   git merge hotfix/v1.1.1
   git tag -a v1.1.1 -m "Hotfix: Critical dependency detection bug"
   git push origin main v1.1.1
   ```

5. Create GitHub release

## Pre-Release Process

For beta/RC releases:

1. Set PreRelease in version.go:
   ```go
   PreRelease = "beta" // or "rc.1", "rc.2", etc.
   ```

2. Build and tag:
   ```bash
   .\build.ps1 -OutputName flutter-installer-v1.2.0-beta
   git tag -a v1.2.0-beta -m "Beta release v1.2.0"
   git push origin v1.2.0-beta
   ```

3. Mark as pre-release on GitHub
4. Gather feedback
5. Fix issues and release RC versions
6. Final stable release when ready

## Build Script Options

### PowerShell (Windows)

```powershell
# Basic build
.\build.ps1

# Custom output name
.\build.ps1 -OutputName my-flutter-tool

# The script automatically:
# - Gets current date/time
# - Gets git commit hash (if in git repo)
# - Gets git branch name
# - Injects into binary via ldflags
```

### Bash (Linux/macOS)

```bash
# Make executable
chmod +x build.sh

# Basic build
./build.sh

# Custom output name
./build.sh my-flutter-tool
```

## Checklist

Before releasing, ensure:

- [ ] Version updated in `pkg/version/version.go`
- [ ] CHANGELOG.md updated
- [ ] README.md updated (if needed)
- [ ] All tests pass
- [ ] Built and tested binary
- [ ] Git committed all changes
- [ ] Git tag created and pushed
- [ ] GitHub release created
- [ ] Binaries attached to release
- [ ] Release notes written
- [ ] Announcement prepared

## Rollback Process

If a release has critical issues:

1. Mark the release as a pre-release on GitHub
2. Add warning notice to release notes
3. Prepare hotfix or new version
4. Release fixed version
5. Update "latest" release pointer

## Version History Template

```markdown
## [X.Y.Z] - YYYY-MM-DD

### Added
- New feature 1
- New feature 2

### Changed
- Changed behavior 1
- Updated dependency X to version Y

### Deprecated
- Feature X will be removed in version Z

### Removed
- Removed deprecated feature Y

### Fixed
- Fixed bug #123
- Fixed issue with...

### Security
- Fixed security vulnerability in...
```

## Automation Ideas (Future)

- GitHub Actions for automated builds
- Automated changelog generation from commits
- Automated version bump PRs
- Release note generation from CHANGELOG
- Binary signing for security
- Checksum generation for downloads

---

**Remember**: Always test thoroughly before releasing!
