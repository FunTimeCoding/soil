package ssh

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDial() *Client {
	errors.PanicOnError(c.Dial())

	return c
}
