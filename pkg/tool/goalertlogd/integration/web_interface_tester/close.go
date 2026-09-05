package web_interface_tester

func (o *Tester) Close() {
	o.server.Close()
}
