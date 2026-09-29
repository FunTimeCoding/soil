package linkace

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDeleteLink(identifier int) {
	errors.PanicOnError(c.DeleteLink(identifier))
}
