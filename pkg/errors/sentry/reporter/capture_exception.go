package reporter

func (r *Reporter) CaptureException(e error) string {
	if r.echo {
		r.print(e, "", nil)
	}

	return r.send(r.hub, e)
}
