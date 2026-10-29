package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) UpdateDevice(
	x context.Context,
	name string,
	body client.UpdateDeviceRequest,
) (client.Device, error) {
	r, e := c.client.UpdateDeviceWithResponse(x, name, body)

	if e != nil {
		return client.Device{}, e
	}

	if r.JSON200 == nil {
		return client.Device{}, refusal.New(
			"update device",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
