package target

func Check(name string) string {
	if IsValid(name) {
		return name
	}

	return "unknown"
}
