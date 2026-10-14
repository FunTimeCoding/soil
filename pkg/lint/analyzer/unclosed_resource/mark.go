package unclosed_resource

func mark(
	m map[string]map[int]bool,
	function string,
	argument int,
) bool {
	if m[function][argument] {
		return false
	}

	if m[function] == nil {
		m[function] = make(map[int]bool)
	}

	m[function][argument] = true

	return true
}
