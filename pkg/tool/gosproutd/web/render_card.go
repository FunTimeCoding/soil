package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web/decision_page"
	"net/http"
)

func (s *Server) renderCard(
	w http.ResponseWriter,
	q *http.Request,
	identifier uint,
) {
	d, e := s.service.Decision(identifier)

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	s.view.RenderFragment(w, decision_page.Card(d))
}
