package example

func OkPerScope(m map[string]int) {
	if v, ok := m["a"]; ok {
		_ = v
	}

	if w, ok := m["b"]; ok {
		_ = w
	}
}
