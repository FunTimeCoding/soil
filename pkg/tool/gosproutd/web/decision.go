package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"net/http"
)

func (s *Server) decision(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier, okay := decisionIdentifier(w, r)

	if !okay {
		return
	}

	found := s.service.Decision(identifier)

	if r.URL.Query().Get(constant.OpenParameter) != "" {
		s.view.RenderFragment(w, decisionRowOpen(found))

		return
	}

	s.view.RenderFragment(w, decisionRow(found))
}
