package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web/decision_page"
	"net/http"
)

func (s *Server) decision(
	w http.ResponseWriter,
	q *http.Request,
) {
	d, okay := s.decisionIdentifier(w, q)

	if !okay {
		return
	}

	if q.URL.Query().Get(constant.OpenParameter) != "" {
		s.view.RenderFragment(w, decision_page.Hero(d, "", ""))

		return
	}

	s.view.RenderFragment(w, decision_page.Card(d))
}
