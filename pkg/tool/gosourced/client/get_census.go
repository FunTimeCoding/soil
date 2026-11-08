package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetCensus(
	x context.Context,
	p *client.GetCensusParams,
) (client.Census, error) {
	r, e := c.client.GetCensusWithResponse(x, p)

	if e != nil {
		return client.Census{}, e
	}

	if r.JSON200 == nil {
		return client.Census{}, refusal.New("census", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
