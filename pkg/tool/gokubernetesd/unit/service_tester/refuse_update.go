package service_tester

func (t *Tester) RefuseUpdate(e error) {
	t.refuse("update", e)
}
