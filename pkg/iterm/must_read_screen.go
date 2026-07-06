package iterm

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/iterm/screen"
)

func (c *Client) MustReadScreen(identifier string) *screen.Screen {
	result, e := c.ReadScreen(identifier)
	errors.PanicOnError(e)

	return result
}
