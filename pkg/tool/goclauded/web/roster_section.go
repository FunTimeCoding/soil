package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) rosterSection() gomponents.Node {
	sessions, e := s.service.ListSessions()
	errors.PanicOnError(e)

	if len(sessions) == 0 {
		return html.P(gomponents.Text("No active sessions."))
	}

	var cards []gomponents.Node

	for i := range sessions {
		labels, f := s.service.LabelsBySession(sessions[i].Identifier)
		errors.PanicOnError(f)
		latest, g := s.service.LatestPulse(sessions[i].Identifier)
		errors.PanicOnError(g)
		cards = append(
			cards,
			sessionCard(&sessions[i], sessions[i].Lines, labels, latest),
		)
	}

	return html.Div(
		html.Class("roster-grid"),
		gomponents.Group(cards),
	)
}
