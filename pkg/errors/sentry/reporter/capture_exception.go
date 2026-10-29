package reporter

func (r *Reporter) CaptureException(e error) string {
	if r.echo {
		r.print(e, "", nil)
	}

	if r.hub == nil {
		return ""
	}

	identifier := r.hub.CaptureException(e)

	if identifier == nil {
		return ""
	}

	return string(*identifier)
}
