package guard

import (
	"github.com/funtimecoding/soil/pkg/web/guard/subject_key"
	"net/http"
)

func Subject(q *http.Request) string {
	subject, _ := q.Context().Value(subject_key.Key{}).(string)

	return subject
}
