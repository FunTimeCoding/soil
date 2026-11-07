package permission

func mostCommon(counts map[string]int) string {
	var result string

	for reason, n := range counts {
		if n > counts[result] || n == counts[result] && reason < result {
			result = reason
		}
	}

	return result
}
