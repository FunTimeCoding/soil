package mock_client

import (
	"github.com/funtimecoding/soil/pkg/iterm/screen"
	"github.com/funtimecoding/soil/pkg/iterm/session"
)

func (c *Client) AddSession(s *session.Session) {
	c.sessions = append(c.sessions, s)
	scr := screen.Stub()
	scr.Identifier = s.Identifier
	c.screens[s.Identifier] = scr
}
