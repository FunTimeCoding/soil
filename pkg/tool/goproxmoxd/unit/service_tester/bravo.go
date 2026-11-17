package service_tester

func (o *Tester) Bravo() *Clients {
	o.t.Helper()

	return o.Instance("bravo")
}
