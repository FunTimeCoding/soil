package suppress

func isKey(field string) bool {
	if field == "" {
		return false
	}

	for _, r := range field {
		if (r < 'a' || r > 'z') && r != '_' {
			return false
		}
	}

	return true
}
