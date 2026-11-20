package request

import "time"

func (q *Request) WithSince(t time.Time) *Request {
	q.Since = t

	return q
}
