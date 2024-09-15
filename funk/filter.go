package funk

type NullableTypes interface {
	int | string
}

func NonEmpty[T NullableTypes](pool ...T) []T {

	result := make([]T, 0, len(pool))
	var empty T
	for _, item := range pool {
		if item == empty {
			continue
		}
		result = append(result, item)
	}
	return result
}

func First[T any](pool []T) (T, bool) {
	var empty T
	if len(pool) == 0 {
		return empty, false
	}
	return pool[0], true
}

func FirstNonEmpty[T NullableTypes](pool ...T) (T, bool) {
	return First(NonEmpty(pool...))
}

func Default[T NullableTypes](exp, fallback T) T {
	var empty T
	if exp == empty {
		return fallback
	}
	return exp
}
