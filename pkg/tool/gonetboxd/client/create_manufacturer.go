package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateManufacturer(
	x context.Context,
	name string,
) (client.Manufacturer, error) {
	r, e := c.client.CreateManufacturerWithResponse(
		x,
		client.CreateNameRequest{Name: name},
	)

	if e != nil {
		return client.Manufacturer{}, e
	}

	if r.JSON201 == nil {
		return client.Manufacturer{}, refusal.New(
			"create manufacturer",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
