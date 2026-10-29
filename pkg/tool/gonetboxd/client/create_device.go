package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateDevice(
	x context.Context,
	name string,
	role string,
	deviceType string,
	site string,
	tenant *string,
) (client.Device, error) {
	r, e := c.client.CreateDeviceWithResponse(
		x,
		client.CreateDeviceRequest{
			Name:   name,
			Role:   role,
			Type:   deviceType,
			Site:   site,
			Tenant: tenant,
		},
	)

	if e != nil {
		return client.Device{}, e
	}

	if r.JSON201 == nil {
		return client.Device{}, refusal.New(
			"create device",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
