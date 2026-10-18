package request

import "context"

func (q *Request) WithContext(x context.Context) *Request {
	q.Context = x

	return q
}
