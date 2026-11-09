package roster_card

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"

func New(s *session.Session) *Card {
	return &Card{Session: s}
}
