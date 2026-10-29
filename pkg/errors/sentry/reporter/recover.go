package reporter

func (r *Reporter) Recover(v any) string {
	if r.echo {
		r.print(v, "", nil)
	}

	if r.hub == nil {
		return ""
	}

	identifier := r.hub.Recover(v)

	if identifier == nil {
		return ""
	}

	return string(*identifier)
}
