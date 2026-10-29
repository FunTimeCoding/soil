package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetDevice(
	x context.Context,
	name string,
) (client.Device, error) {
	r, e := c.client.GetDeviceWithResponse(x, name)

	if e != nil {
		return client.Device{}, e
	}

	if r.JSON200 == nil {
		return client.Device{}, refusal.New(
			"get device",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
