package model_context

func optionalText(
	target **string,
	value string,
) {
	if value == "" {
		return
	}

	*target = &value
}
