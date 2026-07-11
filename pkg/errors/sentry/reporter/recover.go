package reporter

func (r *Reporter) Recover(v any) string {
	if r.hub == nil {
		return ""
	}

	identifier := r.hub.Recover(v)

	if identifier == nil {
		return ""
	}

	return string(*identifier)
}
