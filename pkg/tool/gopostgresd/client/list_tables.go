package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListTables(
	x context.Context,
	params *client.ListTablesParams,
) (*client.QueryResult, error) {
	r, e := c.client.ListTablesWithResponse(x, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list tables", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
