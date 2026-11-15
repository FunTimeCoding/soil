package request

func (q *Request) WithRefusal(f func(int, []byte) error) *Request {
	q.Refusal = f

	return q
}
