package semver

// NextMajor returns major-bumped semver.
func (semver Semver) NextMajor() Semver {
	return Semver{semver.Major + 1, 0, 0, "", ""}
}

// NextMinor returns minor-bumped semver.
func (semver Semver) NextMinor() Semver {
	return Semver{semver.Major, semver.Minor + 1, 0, "", ""}
}

// NextPatch returns patch-bumped semver.
func (semver Semver) NextPatch() Semver {
	return Semver{semver.Major, semver.Minor, semver.Patch + 1, "", ""}
}
