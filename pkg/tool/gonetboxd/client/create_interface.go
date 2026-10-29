package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateInterface(
	x context.Context,
	device string,
	name string,
	interfaceType string,
	physicalAddress string,
) (client.Interface, error) {
	body := client.CreateInterfaceRequest{Name: name, Type: interfaceType}

	if physicalAddress != "" {
		body.PhysicalAddress = &physicalAddress
	}

	r, e := c.client.CreateInterfaceWithResponse(x, device, body)

	if e != nil {
		return client.Interface{}, e
	}

	if r.JSON201 == nil {
		return client.Interface{}, refusal.New(
			"create interface",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
