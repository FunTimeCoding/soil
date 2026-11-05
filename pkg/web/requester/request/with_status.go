package request

func (q *Request) WithStatus(
	status int,
	class string,
) *Request {
	if q.Statuses == nil {
		q.Statuses = map[int]string{}
	}

	q.Statuses[status] = class

	return q
}
