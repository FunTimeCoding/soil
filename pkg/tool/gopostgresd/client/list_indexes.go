package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListIndexes(
	x context.Context,
	params *client.ListIndexesParams,
) (*client.QueryResult, error) {
	r, e := c.client.ListIndexesWithResponse(x, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list indexes", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
