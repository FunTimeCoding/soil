package store

func FirstValue(values []string) string {
	if len(values) == 0 {
		return ""
	}

	return values[0]
}
