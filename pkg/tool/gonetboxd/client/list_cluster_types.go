package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListClusterTypes(
	x context.Context,
) ([]*client.ClusterType, error) {
	r, e := c.client.ListClusterTypesWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list cluster types", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
