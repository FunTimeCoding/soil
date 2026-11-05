package requester

func (r *Requester) WithSubject(s string) *Requester {
	r.subject = s

	return r
}
