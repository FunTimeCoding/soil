package search_page

func copyWithout(
	m map[string]string,
	key string,
) map[string]string {
	result := map[string]string{}

	for k, v := range m {
		if k != key {
			result[k] = v
		}
	}

	return result
}
