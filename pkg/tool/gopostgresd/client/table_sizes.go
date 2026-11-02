package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) TableSizes(
	x context.Context,
	params *client.TableSizesParams,
) (*client.QueryResult, error) {
	r, e := c.client.TableSizesWithResponse(x, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("table sizes", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
