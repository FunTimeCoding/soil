package stamp

func firstSet(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}

	return ""
}
