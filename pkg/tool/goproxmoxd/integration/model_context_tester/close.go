package model_context_tester

func (o *Tester) Close() {
	o.Client.Close()
	o.server.Close()
}
