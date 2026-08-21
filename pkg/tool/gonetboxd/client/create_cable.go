package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateCable(
	deviceA string,
	interfaceA string,
	deviceB string,
	interfaceB string,
) string {
	result, e := c.client.CreateCable(
		c.context,
		client.CreateCableRequest{
			DeviceA:    deviceA,
			InterfaceA: interfaceA,
			DeviceB:    deviceB,
			InterfaceB: interfaceB,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
