package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) SearchItems(
	x context.Context,
	p *client.SearchItemsParams,
) (*client.SearchResult, error) {
	r, e := c.client.SearchItemsWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("search items", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
