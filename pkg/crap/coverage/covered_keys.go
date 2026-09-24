package coverage

func CoveredKeys(functions map[string]float64) []string {
	var result []string

	for key, percent := range functions {
		if percent > 0 {
			result = append(result, key)
		}
	}

	return result
}
