package example

func OkRename(m map[string]int) {
	v, found := m["key"]
	_ = v
	_ = found
}
