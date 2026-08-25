package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/iterm/screen"
	"github.com/funtimecoding/soil/pkg/iterm/session"
)

func (c *Client) CreateTab() (*session.Session, error) {
	identifier := fmt.Sprintf("session-%d", c.nextIdentifier)
	c.nextIdentifier++
	s := session.Stub()
	s.Identifier = identifier
	s.TabIdentifier = fmt.Sprintf("tab-%d", c.nextIdentifier-1)
	c.sessions = append(c.sessions, s)
	scr := screen.Stub()
	scr.Identifier = identifier
	c.screens[identifier] = scr

	return s, nil
}
