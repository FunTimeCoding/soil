package service_tester

func (o *Tester) Alfa() *Clients {
	o.t.Helper()

	return o.Instance("alfa")
}
