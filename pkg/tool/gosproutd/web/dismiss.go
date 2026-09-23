package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"net/http"
)

func (s *Server) dismiss(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier, okay := decisionIdentifier(w, r)

	if !okay {
		return
	}

	state := constant.State(r.URL.Query().Get(constant.StateParameter))

	switch state {
	case constant.StateDeclined, constant.StateIrrelevant,
		constant.StatePostponed:
		s.service.Dismiss(identifier, state)
	default:
		http.Error(w, "invalid state", http.StatusBadRequest)

		return
	}

	s.view.RenderFragment(w, decisionRow(s.service.Decision(identifier)))
}
