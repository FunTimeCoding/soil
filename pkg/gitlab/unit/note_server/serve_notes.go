package note_server

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"strconv"
)

func (s *Server) serveNotes(
	w http.ResponseWriter,
	q *http.Request,
) {
	s.lock.Lock()
	s.queries = append(s.queries, q.URL.Query())
	s.lock.Unlock()
	page, e := strconv.Atoi(q.URL.Query().Get("page"))
	errors.PanicOnError(e)
	size, e := strconv.Atoi(q.URL.Query().Get("per_page"))
	errors.PanicOnError(e)
	list := []map[string]any{}

	for i := s.total - (page-1)*size; i > max(s.total-page*size, 0); i-- {
		list = append(
			list,
			map[string]any{
				"id":         i,
				"body":       join.Empty("note ", strconv.Itoa(i)),
				"author":     map[string]any{"username": "delta"},
				"system":     false,
				"created_at": "2026-11-25T10:00:00Z",
			},
		)
	}

	web.Encode(w, list)
}
