package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateDeviceType(
	x context.Context,
	model string,
	manufacturer string,
) (client.DeviceType, error) {
	r, e := c.client.CreateDeviceTypeWithResponse(
		x,
		client.CreateDeviceTypeRequest{
			Model:        model,
			Manufacturer: manufacturer,
		},
	)

	if e != nil {
		return client.DeviceType{}, e
	}

	if r.JSON201 == nil {
		return client.DeviceType{}, refusal.New(
			"create device type",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
