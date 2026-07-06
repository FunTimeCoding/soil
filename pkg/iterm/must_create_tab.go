package iterm

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/iterm/session"
)

func (c *Client) MustCreateTab() *session.Session {
	result, e := c.CreateTab()
	errors.PanicOnError(e)

	return result
}
