package funk

// Diff returns slice containing A - B
func Diff[T comparable](a, b []T) []T {

	// create lookup for b
	elements := make(map[T]struct{})
	for _, item := range b {
		elements[item] = struct{}{}
	}

	// build diff
	diff := []T{}
	for _, item := range a {
		if _, found := elements[item]; !found {
			diff = append(diff, item)
		}
	}
	return diff
}
