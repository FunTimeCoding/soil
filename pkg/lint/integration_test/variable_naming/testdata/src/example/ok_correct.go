package example

func OkCorrect(m map[string]int) {
	v, okay := m["key"]
	_ = v
	_ = okay
}
