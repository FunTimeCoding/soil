package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateCable(
	deviceA string,
	interfaceA string,
	deviceB string,
	interfaceB string,
) *response.Response {
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

	return response.New(web.ReadString(result), result.StatusCode)
}
