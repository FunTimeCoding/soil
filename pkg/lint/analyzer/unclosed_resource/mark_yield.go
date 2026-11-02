package unclosed_resource

func markYield(
	m map[string]map[int]map[string]bool,
	function string,
	argument int,
	field string,
) bool {
	if m[function][argument][field] {
		return false
	}

	if m[function] == nil {
		m[function] = make(map[int]map[string]bool)
	}

	if m[function][argument] == nil {
		m[function][argument] = make(map[string]bool)
	}

	m[function][argument][field] = true

	return true
}
