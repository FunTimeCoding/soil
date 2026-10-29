package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListTunnelGroups(
	x context.Context,
) ([]*client.TunnelGroup, error) {
	r, e := c.client.ListTunnelGroupsWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list tunnel groups", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
