package funk

import "strings"

func Contains[T comparable](pool []T, item T) bool {
	for _, p := range pool {
		if p == item {
			return true
		}
	}
	return false
}

func FuzzySearch(exp string, pool []string) []string {
	sel := make([]string, 0)
	for _, p := range pool {
		if strings.HasPrefix(p, exp) {
			sel = append(sel, p)
		}
	}
	return sel
}
