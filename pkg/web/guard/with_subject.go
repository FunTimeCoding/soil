package guard

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/guard/subject_key"
	"net/http"
)

func WithSubject(
	q *http.Request,
	subject string,
) *http.Request {
	return q.WithContext(
		context.WithValue(q.Context(), subject_key.Key{}, subject),
	)
}
