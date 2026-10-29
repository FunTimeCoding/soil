package model_context_tester

func (o *Tester) Captured() int {
	return len(o.server.Reporter.Events())
}
