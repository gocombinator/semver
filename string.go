package semver

import "fmt"

// String returns normalised semver string.
func (semver Semver) String() string {
	if semver.Prerelease != "" && semver.Build != "" {
		return fmt.Sprintf("v%d.%d.%d-%s+%s", semver.Major, semver.Minor, semver.Patch, semver.Prerelease, semver.Build)
	}
	if semver.Prerelease != "" {
		return fmt.Sprintf("v%d.%d.%d-%s", semver.Major, semver.Minor, semver.Patch, semver.Prerelease)
	}
	return fmt.Sprintf("v%d.%d.%d", semver.Major, semver.Minor, semver.Patch)
}
