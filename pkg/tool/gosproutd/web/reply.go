package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web/decision_page"
	"net/http"
)

func (s *Server) reply(
	w http.ResponseWriter,
	q *http.Request,
) {
	d, okay := s.decisionIdentifier(w, q)

	if !okay {
		return
	}

	alert := ""

	if content := q.FormValue(constant.ContentParameter); content != "" {
		if _, e := s.service.Reply(d.Identifier, content); e != nil {
			alert = e.Error()
		}
	}

	updated, f := s.service.Decision(d.Identifier)

	if f != nil {
		s.view.Answer(w, q, s.reporter, f)

		return
	}

	s.view.RenderFragment(w, decision_page.Hero(updated, "", alert))
}
