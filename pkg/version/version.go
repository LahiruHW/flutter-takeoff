package version

import (
	"fmt"
	"time"
)

// Version information
const (
	Major      = 1
	Minor      = 0
	Patch      = 1
	PreRelease = "" // e.g., "beta", "rc.1", empty for stable
)

var (
	// BuildDate is set at build time
	BuildDate = "unknown"
	// GitCommit is set at build time
	GitCommit = "unknown"
	// GitBranch is set at build time
	GitBranch = "unknown"
)

// Version returns the semantic version string
func Version() string {
	v := fmt.Sprintf("%d.%d.%d", Major, Minor, Patch)
	if PreRelease != "" {
		v += "-" + PreRelease
	}
	return v
}

// FullVersion returns the version with build metadata
func FullVersion() string {
	v := Version()
	if GitCommit != "unknown" {
		v += fmt.Sprintf(" (commit: %.7s)", GitCommit)
	}
	if BuildDate != "unknown" {
		v += fmt.Sprintf(" built on %s", BuildDate)
	}
	return v
}

// BuildInfo returns detailed build information
func BuildInfo() map[string]string {
	return map[string]string{
		"version":    Version(),
		"build_date": BuildDate,
		"git_commit": GitCommit,
		"git_branch": GitBranch,
	}
}

// GetBuildDate returns the build date as time.Time if parseable
func GetBuildDate() (time.Time, error) {
	if BuildDate == "unknown" {
		return time.Time{}, fmt.Errorf("build date unknown")
	}
	return time.Parse(time.RFC3339, BuildDate)
}

// IsPreRelease returns true if this is a pre-release version
func IsPreRelease() bool {
	return PreRelease != ""
}

// Compare compares this version with another version string
// Returns: -1 if this < other, 0 if equal, 1 if this > other
func Compare(other string) int {
	current := fmt.Sprintf("%d.%d.%d", Major, Minor, Patch)
	if current < other {
		return -1
	}
	if current > other {
		return 1
	}
	return 0
}
