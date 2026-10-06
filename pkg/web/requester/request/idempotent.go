package request

import "net/http"

func (q *Request) Idempotent() bool {
	return q.Method == http.MethodGet || q.Method == http.MethodHead
}
