package web

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"net/http"
)

func (s *Server) dismiss(
	w http.ResponseWriter,
	q *http.Request,
) {
	d, okay := s.decisionIdentifier(w, q)

	if !okay {
		return
	}

	state := constant.State(q.FormValue(constant.StateParameter))

	switch state {
	case constant.StateDeclined, constant.StateIrrelevant,
		constant.StatePostponed:
		s.service.Dismiss(d.Identifier, state)
	default:
		s.view.Answer(
			w,
			q,
			s.reporter,
			validation.New("invalid state: %s", state),
		)

		return
	}

	s.renderCard(w, q, d.Identifier)
}
