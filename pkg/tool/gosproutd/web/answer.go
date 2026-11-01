package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web/decision_page"
	"net/http"
)

func (s *Server) answer(
	w http.ResponseWriter,
	q *http.Request,
) {
	d, okay := s.decisionIdentifier(w, q)

	if !okay {
		return
	}

	value := q.FormValue(constant.ValueParameter)

	if value == "" {
		s.view.RenderFragment(
			w,
			decision_page.Hero(d, "Write an answer first.", ""),
		)

		return
	}

	s.service.Answer(
		d.Identifier,
		constant.AnswerKind(q.FormValue(constant.KindParameter)),
		value,
	)
	s.renderCard(w, q, d.Identifier)
}
