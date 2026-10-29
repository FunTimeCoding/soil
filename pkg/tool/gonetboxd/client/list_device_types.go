package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListDeviceTypes(
	x context.Context,
) ([]*client.DeviceType, error) {
	r, e := c.client.ListDeviceTypesWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list device types", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
