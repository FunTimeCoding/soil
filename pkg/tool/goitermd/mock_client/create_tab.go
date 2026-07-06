package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/iterm/screen"
	"github.com/funtimecoding/soil/pkg/iterm/session"
)

func (c *Client) CreateTab() (*session.Session, error) {
	identifier := fmt.Sprintf("session-%d", c.nextID)
	c.nextID++
	s := session.Stub()
	s.Identifier = identifier
	s.TabIdentifier = fmt.Sprintf("tab-%d", c.nextID-1)
	c.sessions = append(c.sessions, s)
	scr := screen.Stub()
	scr.Identifier = identifier
	c.screens[identifier] = scr

	return s, nil
}
