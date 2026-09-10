package refusal

func (r *Refusal) Error() string {
	return r.reason
}
