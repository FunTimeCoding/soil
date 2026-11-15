package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetUsage(
	x context.Context,
	p *client.GetUsageParams,
) (client.Usage, error) {
	r, e := c.client.GetUsageWithResponse(x, p)

	if e != nil {
		return client.Usage{}, e
	}

	if r.JSON200 == nil {
		return client.Usage{}, refusal.New("usage", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
