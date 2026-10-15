package ansible_inventory

func selected(
	fields map[string]any,
	field string,
) []string {
	var result []string

	switch v := fields[field].(type) {
	case []string:
		result = v
	case []any:
		for _, e := range v {
			if s, okay := e.(string); okay {
				result = append(result, s)
			}
		}
	}

	return result
}
