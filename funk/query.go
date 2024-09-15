package funk

func Contains[T comparable](pool []T, item T) bool {
	for _, p := range pool {
		if p == item {
			return true
		}
	}
	return false
}
