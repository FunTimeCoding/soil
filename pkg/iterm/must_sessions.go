package iterm

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/iterm/session"
)

func (c *Client) MustSessions() []*session.Session {
	result, e := c.Sessions()
	errors.PanicOnError(e)

	return result
}
