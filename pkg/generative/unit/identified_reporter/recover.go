package identified_reporter

func (r *Reporter) Recover(v any) string {
	r.Memory.Recover(v)

	return r.identifier
}
