package request

import "net/http"

func (q *Request) Unchanged(status int) bool {
	return !q.Since.IsZero() && status == http.StatusNotModified
}
