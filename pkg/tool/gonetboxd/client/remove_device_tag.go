package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) RemoveDeviceTag(
	x context.Context,
	device string,
	tag string,
) (client.Device, error) {
	r, e := c.client.RemoveDeviceTagWithResponse(x, device, tag)

	if e != nil {
		return client.Device{}, e
	}

	if r.JSON200 == nil {
		return client.Device{}, refusal.New(
			"remove device tag",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
