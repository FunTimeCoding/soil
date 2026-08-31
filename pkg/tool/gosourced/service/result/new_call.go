package result

func NewCall(name string, count int) *Call {
	return &Call{Name: name, Count: count}
}
