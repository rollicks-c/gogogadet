package funk

func Map[T1 comparable, T2 any](data []T1, f func(e T1) T2) []T2 {
	result := make([]T2, len(data))
	for i, v := range data {
		result[i] = f(v)
	}
	return result
}
