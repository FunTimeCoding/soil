package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListSchemas(
	x context.Context,
	params *client.ListSchemasParams,
) (*client.QueryResult, error) {
	r, e := c.client.ListSchemasWithResponse(x, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list schemas", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
