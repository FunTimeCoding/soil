package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"net/http"
)

func (s *Server) activityPage(
	w http.ResponseWriter,
	q *http.Request,
) {
	n, e := s.activity(q.Context(), q.URL.Query()[constant.Kind])

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	s.view.RenderFragment(w, n)
}
