package request

func (q *Request) WithHeader(
	k string,
	v string,
) *Request {
	q.Header[k] = v

	return q
}
