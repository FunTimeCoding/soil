package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListDevices(
	x context.Context,
	query *string,
) ([]*client.Device, error) {
	r, e := c.client.ListDevicesWithResponse(
		x,
		&client.ListDevicesParams{Query: query},
	)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list devices", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
