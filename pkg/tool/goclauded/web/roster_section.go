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

	identifiers := make([]string, 0, len(sessions))

	for _, i := range sessions {
		identifiers = append(identifiers, i.Identifier)
	}

	labels, f := s.service.LabelsBySessions(identifiers)
	errors.PanicOnError(f)
	pulses, g := s.service.LatestPulsesBySessions(identifiers)
	errors.PanicOnError(g)
	var cards []gomponents.Node

	for i := range sessions {
		cards = append(
			cards,
			sessionCard(
				&sessions[i],
				sessions[i].Lines,
				labels[sessions[i].Identifier],
				pulses[sessions[i].Identifier],
			),
		)
	}

	return html.Div(html.Class("roster-grid"), gomponents.Group(cards))
}
