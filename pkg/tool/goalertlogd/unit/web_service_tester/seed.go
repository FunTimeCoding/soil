package web_service_tester

func (o *Tester) Seed(count int) {
	o.server.Seed(count)
}
