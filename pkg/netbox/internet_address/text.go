package internet_address

func text(
	object map[string]any,
	key string,
) string {
	result, _ := object[key].(string)

	return result
}
