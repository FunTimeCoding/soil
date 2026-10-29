package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetOversize(
	x context.Context,
	p *client.GetOversizeParams,
) (*client.OversizeReport, error) {
	r, e := c.client.GetOversizeWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("oversize", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
