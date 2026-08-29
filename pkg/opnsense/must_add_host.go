package opnsense

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/opnsense/request"
)

func (c *Client) MustAddHost(h *request.Host) string {
	result, e := c.AddHost(h)
	errors.PanicOnError(e)

	return result
}
