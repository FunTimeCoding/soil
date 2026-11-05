package requester

func (r *Requester) WithStatus(
	status int,
	class string,
) *Requester {
	r.table = r.table.With(status, class)

	return r
}
