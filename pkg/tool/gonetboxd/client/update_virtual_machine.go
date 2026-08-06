package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) UpdateVirtualMachine(
	name string,
	newName string,
	primaryAddress string,
) string {
	body := client.UpdateVirtualMachineRequest{}

	if newName != "" {
		body.Name = &newName
	}

	if primaryAddress != "" {
		body.PrimaryAddress = &primaryAddress
	}

	result, e := c.client.UpdateVirtualMachine(c.context, name, body)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
