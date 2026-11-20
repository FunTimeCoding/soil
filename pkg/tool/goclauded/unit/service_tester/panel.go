package service_tester

func (o *Tester) Panel(identifier string) *Tester {
	o.t.Helper()
	o.Store.EnsureSession(identifier)
	o.WriteContextLoadFile(identifier)
	o.Transcript.EnrichSession(identifier)

	return o
}
