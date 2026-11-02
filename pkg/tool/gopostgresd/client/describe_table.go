package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) DescribeTable(
	x context.Context,
	table string,
	params *client.DescribeTableParams,
) (*client.QueryResult, error) {
	r, e := c.client.DescribeTableWithResponse(x, table, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("describe table", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
