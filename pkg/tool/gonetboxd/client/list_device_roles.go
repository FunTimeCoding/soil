package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListDeviceRoles(
	x context.Context,
) ([]*client.DeviceRole, error) {
	r, e := c.client.ListDeviceRolesWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list device roles", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
