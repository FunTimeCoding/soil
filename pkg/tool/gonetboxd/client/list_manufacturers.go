package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListManufacturers(
	x context.Context,
) ([]*client.Manufacturer, error) {
	r, e := c.client.ListManufacturersWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list manufacturers", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
