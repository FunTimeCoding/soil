package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/session_page"
	"maragu.dev/gomponents"
)

func (s *Server) rosterSection() (gomponents.Node, error) {
	cards, e := s.rosterCards()

	if e != nil {
		return nil, e
	}

	return session_page.Roster(cards), nil
}
