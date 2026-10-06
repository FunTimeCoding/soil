package requester

func (r *Requester) WithHeader(
	k string,
	v string,
) *Requester {
	r.header[k] = v

	return r
}
