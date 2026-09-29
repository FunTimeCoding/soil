package linkace

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDeleteList(identifier int) {
	errors.PanicOnError(c.DeleteList(identifier))
}
