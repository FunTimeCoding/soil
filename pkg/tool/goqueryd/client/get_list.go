package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetList(
	x context.Context,
	p *client.GetListParams,
) (*client.ListOutcome, error) {
	r, e := c.client.GetListWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list documents", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
