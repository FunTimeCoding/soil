package guard

import (
	"context"
	"net/http"
)

func WithSubject(
	q *http.Request,
	subject string,
) *http.Request {
	return q.WithContext(context.WithValue(q.Context(), subjectKey{}, subject))
}
