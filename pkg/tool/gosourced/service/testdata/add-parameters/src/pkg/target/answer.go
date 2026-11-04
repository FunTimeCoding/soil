package target

func Answer(
	message string,
	code int,
) string {
	if code > 0 {
		return message
	}

	return ""
}
