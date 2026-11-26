package discussion_server

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web"
	"io"
	"net/http"
	"strings"
)

func (s *Server) serve(
	w http.ResponseWriter,
	q *http.Request,
) {
	body, e := io.ReadAll(q.Body)
	errors.PanicOnError(e)
	s.lock.Lock()
	s.requests = append(
		s.requests,
		join.Space(
			q.Method,
			q.URL.RequestURI(),
			strings.TrimSpace(string(body)),
		),
	)
	s.lock.Unlock()

	if q.Method == http.MethodPost {
		web.Encode(
			w,
			map[string]any{
				"id":     11,
				"body":   "answered",
				"author": map[string]any{"username": "delta"},
			},
		)

		return
	}

	web.Encode(w, map[string]any{"id": "thread", "notes": []any{}})
}
