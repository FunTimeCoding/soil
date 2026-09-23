package unit

func buildContains(
	built []string,
	entry string,
) bool {
	for _, candidate := range built {
		if candidate == entry {
			return true
		}
	}

	return false
}
