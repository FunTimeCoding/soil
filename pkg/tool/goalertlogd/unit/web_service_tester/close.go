package web_service_tester

func (o *Tester) Close() {
	o.server.Close()
}
