package opnsense

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDeleteHost(identifier string) {
	errors.PanicOnError(c.DeleteHost(identifier))
}
