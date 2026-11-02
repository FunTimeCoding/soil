package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/constant"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) Search(
	x context.Context,
	p *client.SearchParams,
) (*client.SearchResult, error) {
	r, e := c.client.SearchWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(constant.Search, r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
