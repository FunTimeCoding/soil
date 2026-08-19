package web

func optionalText(
	target **string,
	value string,
) {
	if value == "" {
		return
	}

	*target = &value
}
