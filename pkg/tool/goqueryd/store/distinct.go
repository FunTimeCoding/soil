package store

func distinct(values []string) []string {
	seen := map[string]bool{}
	var result []string

	for _, value := range values {
		if seen[value] {
			continue
		}

		seen[value] = true
		result = append(result, value)
	}

	return result
}
