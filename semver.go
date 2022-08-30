package semver

type Semver struct {
	Major      int
	Minor      int
	Patch      int
	Prerelease string
	Build      string
}

var Empty = Semver{}

// Lt returns true if a < b, false otherwise.
func Lt(a, b Semver) bool {
	if a.Major != b.Major {
		return a.Major < b.Major
	}
	if a.Minor != b.Minor {
		return a.Minor < b.Minor
	}
	if a.Patch != b.Patch {
		return a.Patch < b.Patch
	}
	if a.Prerelease != b.Prerelease {
		return a.Prerelease < b.Prerelease
	}
	return a.Build < b.Build
}
