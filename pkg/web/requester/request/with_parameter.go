package request

func (q *Request) WithParameter(
	k string,
	v string,
) *Request {
	q.Parameters.Set(k, v)

	return q
}
