package semver

import (
	"regexp"
	"strconv"
)

var Re = regexp.MustCompile(`^v?(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

// MaybeParse parses semver string into struct.
func MaybeParse(value string) (Semver, bool) {
	var m = Re.FindAllStringSubmatch(value, -1)
	if m == nil {
		return Empty, false
	}
	var major, _ = strconv.Atoi(m[0][1])
	var minor, _ = strconv.Atoi(m[0][2])
	var patch, _ = strconv.Atoi(m[0][3])
	var prerelease = m[0][4]
	var build = m[0][5]
	return Semver{major, minor, patch, prerelease, build}, true
}

// Parse semver string into struct.
// Returns [Empty] on failure.
func Parse(value string) Semver {
	if semver, ok := MaybeParse(value); ok {
		return semver
	} else {
		return Empty
	}
}
