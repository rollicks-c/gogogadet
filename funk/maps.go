package funk

func PickOne[T1 comparable, T2 any](data map[T1]T2) (T2, bool) {
	for _, v := range data {
		return v, true
	}
	var empty T2
	return empty, false
}

func GetKeys[T1 comparable, T2 any](data map[T1]T2) []T1 {
	keys := make([]T1, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	return keys
}
