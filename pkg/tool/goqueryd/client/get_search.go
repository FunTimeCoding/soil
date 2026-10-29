package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetSearch(
	x context.Context,
	p *client.GetSearchParams,
) (*client.SearchOutcome, error) {
	r, e := c.client.GetSearchWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(constant.Search, r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
