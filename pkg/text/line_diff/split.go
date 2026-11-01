package line_diff

func split(s string) []string {
	var result []string
	start := 0

	for i, c := range s {
		if c == '\n' {
			result = append(result, s[start:i])
			start = i + 1
		}
	}

	if start < len(s) {
		result = append(result, s[start:])
	}

	return result
}
