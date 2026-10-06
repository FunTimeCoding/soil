package requester

func (r *Requester) WithRefusal(f func(int, []byte) error) *Requester {
	r.refusal = f

	return r
}
