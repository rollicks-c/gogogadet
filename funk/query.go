package funk

import (
	"golang.org/x/exp/constraints"
	_ "golang.org/x/exp/constraints"
	"strings"
)

func Contains[T comparable](pool []T, item T) bool {
	for _, p := range pool {
		if p == item {
			return true
		}
	}
	return false
}

func Max[T constraints.Ordered](pool ...T) T {
	var maxValue T
	for i, p := range pool {
		if i == 0 || p > maxValue {
			maxValue = p
		}
	}
	return maxValue
}

func Min[T constraints.Ordered](pool ...T) T {
	var minValue T
	for i, p := range pool {
		if i == 0 || p < minValue {
			minValue = p
		}
	}
	return minValue
}

func FuzzySearch(exp string, pool ...string) []string {
	sel := make([]string, 0)
	for _, p := range pool {
		if strings.HasPrefix(p, exp) {
			sel = append(sel, p)
		}
	}
	return sel
}
