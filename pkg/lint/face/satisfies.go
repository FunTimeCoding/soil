package face

func satisfies(
	methods map[string]string,
	required map[string]string,
) bool {
	for name, signature := range required {
		if methods[name] != signature {
			return false
		}
	}

	return true
}
