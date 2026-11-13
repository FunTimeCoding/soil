package service_tester

func (o *Tester) WriteContextLoadFile(identifier string) {
	o.t.Helper()
	o.WriteTranscript(identifier, "context-loads.jsonl")
}
