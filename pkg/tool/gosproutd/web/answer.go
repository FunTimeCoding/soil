package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"net/http"
)

func (s *Server) answer(
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

	kind := constant.AnswerKindChoice
	value := r.URL.Query().Get(constant.ValueParameter)

	if value == "" {
		value = r.FormValue(constant.ValueParameter)
		kind = constant.AnswerKind(r.FormValue(constant.KindParameter))
	}

	if value == "" {
		s.view.RenderFragment(
			w,
			decisionRowOpen(s.service.Decision(identifier)),
		)

		return
	}

	s.service.Answer(identifier, kind, value)
	s.view.RenderFragment(w, decisionRow(s.service.Decision(identifier)))
}
