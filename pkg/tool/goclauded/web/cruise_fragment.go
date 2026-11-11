package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/session_page"
	"net/http"
)

func (s *Server) cruiseFragment(
	w http.ResponseWriter,
	q *http.Request,
	identifier string,
	refused error,
) {
	setting, e := s.cruise.Setting(identifier)

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	alert := ""

	if refused != nil {
		alert = refused.Error()
	}

	s.view.RenderFragment(
		w,
		session_page.RosterCruise(
			identifier,
			setting,
			s.cruise.Grace(setting),
			alert,
		),
	)
}
