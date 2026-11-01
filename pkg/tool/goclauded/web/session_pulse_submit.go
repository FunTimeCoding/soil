package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/session_page"
	"net/http"
)

func (s *Server) sessionPulseSubmit(
	w http.ResponseWriter,
	q *http.Request,
) {
	identifier := q.PathValue(constant.Identifier)

	if body := q.FormValue(constant.Body); body != "" {
		d := constant.DeliveryQueued

		if q.FormValue(constant.Immediate) != "" {
			d = constant.DeliveryImmediate
		}

		if _, e := s.service.SendPulse(identifier, "", body, d); e != nil {
			s.view.Answer(w, q, s.reporter, e)

			return
		}
	}

	http.Redirect(w, q, session_page.Path(identifier), http.StatusSeeOther)
}
