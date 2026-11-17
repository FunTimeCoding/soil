package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreatePowerPort(
	x context.Context,
	device string,
	name string,
) (client.PowerPort, error) {
	r, e := c.client.CreatePowerPortWithResponse(
		x,
		device,
		client.CreatePowerPortRequest{Name: name},
	)

	if e != nil {
		return client.PowerPort{}, e
	}

	if r.JSON201 == nil {
		return client.PowerPort{}, refusal.New(
			"create power port",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
