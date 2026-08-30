package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListVirtualMachines() (string, int) {
	result, e := c.client.ListVirtualMachines(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result), result.StatusCode
}
