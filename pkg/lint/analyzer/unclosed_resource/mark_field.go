package unclosed_resource

func markField(
	m map[string]map[string]bool,
	holder string,
	field string,
) bool {
	if m[holder][field] {
		return false
	}

	if m[holder] == nil {
		m[holder] = make(map[string]bool)
	}

	m[holder][field] = true

	return true
}
