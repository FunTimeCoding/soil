package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) AddDeviceTag(
	x context.Context,
	device string,
	tag string,
) (client.Device, error) {
	r, e := c.client.AddDeviceTagWithResponse(x, device, tag)

	if e != nil {
		return client.Device{}, e
	}

	if r.JSON200 == nil {
		return client.Device{}, refusal.New(
			"add device tag",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
