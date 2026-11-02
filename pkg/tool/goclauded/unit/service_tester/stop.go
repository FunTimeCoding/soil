package service_tester

func (t *Tester) Stop() {
	t.stop.Close()
}
