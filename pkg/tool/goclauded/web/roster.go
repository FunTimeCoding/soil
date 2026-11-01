package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/session_page"
	"maragu.dev/gomponents"
)

func (s *Server) roster() (gomponents.Node, error) {
	sessions, e := s.service.ListSessions()

	if e != nil {
		return nil, e
	}

	identifiers := make([]string, 0, len(sessions))

	for _, v := range sessions {
		identifiers = append(identifiers, v.Identifier)
	}

	labels, f := s.service.LabelsBySessions(identifiers)

	if f != nil {
		return nil, f
	}

	pulses, g := s.service.LatestPulsesBySessions(identifiers)

	if g != nil {
		return nil, g
	}

	return session_page.Roster(sessions, labels, pulses), nil
}
