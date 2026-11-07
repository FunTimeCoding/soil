package guard

import "net/http"

func Subject(q *http.Request) string {
	subject, _ := q.Context().Value(subjectKey{}).(string)

	return subject
}
