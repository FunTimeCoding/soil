package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"net/http"
)

func (s *Server) reply(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier, okay := decisionIdentifier(w, r)

	if !okay {
		return
	}

	if e := r.ParseForm(); e != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)

		return
	}

	if content := r.FormValue(constant.ContentParameter); content != "" {
		if _, f := s.service.Reply(identifier, content); f != nil {
			http.Error(w, f.Error(), http.StatusBadRequest)

			return
		}
	}

	s.view.RenderFragment(w, decisionRowOpen(s.service.Decision(identifier)))
}
