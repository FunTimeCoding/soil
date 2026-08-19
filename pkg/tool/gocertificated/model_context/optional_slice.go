package model_context

func optionalSlice(
	target **[]string,
	value []string,
) {
	if len(value) == 0 {
		return
	}

	*target = &value
}
