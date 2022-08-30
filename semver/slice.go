package semver

type Slice []Semver

func (x Slice) Len() int {
	return len(x)
}

func (x Slice) Less(i, j int) bool {
	return Lt(x[i], x[j])
}

func (x Slice) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
