package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateInterface(
	device string,
	name string,
	interfaceType string,
	physicalAddress string,
) *response.Response {
	body := client.CreateInterfaceRequest{Name: name, Type: interfaceType}

	if physicalAddress != "" {
		body.PhysicalAddress = &physicalAddress
	}

	result, e := c.client.CreateInterface(c.context, device, body)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
