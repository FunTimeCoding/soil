package service_tester

func (t *Tester) DynamicCalls() int {
	return len(t.Dynamic.Actions())
}
