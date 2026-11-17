package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreatePowerOutlet(
	x context.Context,
	device string,
	name string,
) (client.PowerOutlet, error) {
	r, e := c.client.CreatePowerOutletWithResponse(
		x,
		device,
		client.CreatePowerOutletRequest{Name: name},
	)

	if e != nil {
		return client.PowerOutlet{}, e
	}

	if r.JSON201 == nil {
		return client.PowerOutlet{}, refusal.New(
			"create power outlet",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
