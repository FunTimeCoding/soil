package model_context

func optionalNumber(
	target **int,
	value float64,
) {
	if value <= 0 {
		return
	}

	converted := int(value)
	*target = &converted
}
