package choice

func FromValues(values []string) []*Choice {
	result := make([]*Choice, 0, len(values))

	for _, v := range values {
		result = append(result, New(v, v))
	}

	return result
}
