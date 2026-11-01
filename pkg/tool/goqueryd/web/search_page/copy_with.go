package search_page

func copyWith(
	m map[string]string,
	key string,
	value string,
) map[string]string {
	result := map[string]string{}

	for k, v := range m {
		result[k] = v
	}

	result[key] = value

	return result
}
