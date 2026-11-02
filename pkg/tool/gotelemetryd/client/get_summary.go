package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetSummary(
	x context.Context,
	p *client.GetSummaryParams,
) ([]client.SummaryEntry, error) {
	r, e := c.client.GetSummaryWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get summary", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
