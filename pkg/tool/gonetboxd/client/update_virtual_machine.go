package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) UpdateVirtualMachine(
	name string,
	body client.UpdateVirtualMachineRequest,
) (string, int) {
	result, e := c.client.UpdateVirtualMachine(c.context, name, body)
	errors.PanicOnError(e)

	return web.ReadString(result), result.StatusCode
}
