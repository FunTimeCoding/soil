package example

func OkNumbered(m map[string]int) {
	v, ok := m["a"]
	w, ok2 := m["b"]
	_ = v
	_ = w
	_ = ok
	_ = ok2
}
