package service_tester

func (t *Tester) RefuseAll(e error) {
	t.refuse("*", e)
}
