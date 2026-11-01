package maps

func Join[K comparable, V any](parts ...map[K]V) map[K]V {
	result := make(map[K]V)

	for _, part := range parts {
		for k, v := range part {
			result[k] = v
		}
	}

	return result
}
