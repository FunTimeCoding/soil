package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) GetVirtualMachine(name string) string {
	result, e := c.client.GetVirtualMachine(c.context, name)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
