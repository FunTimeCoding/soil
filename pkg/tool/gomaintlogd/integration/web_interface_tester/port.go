package web_interface_tester

func (o *Tester) Port() int {
	return o.server.Port()
}
