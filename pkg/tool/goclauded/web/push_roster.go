package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/session_page"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"net/http"
)

func (s *Server) pushRoster(w http.ResponseWriter) {
	cards, e := s.rosterCards()

	if e != nil {
		s.reporter.CaptureException(e)

		return
	}

	layout.PushEvent(w, constant.Roster, session_page.Roster(cards))

	for _, c := range cards {
		identifier := c.Session.Identifier
		layout.PushEvent(
			w,
			session_page.Event(identifier, constant.RosterActivityPart),
			session_page.ActivityDot(c),
		)
		layout.PushEvent(
			w,
			session_page.Event(identifier, constant.RosterStatePart),
			session_page.State(c),
		)
	}
}
