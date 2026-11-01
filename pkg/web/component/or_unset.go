package component

func OrUnset(value string) string {
	if value == "" {
		return "—"
	}

	return value
}
