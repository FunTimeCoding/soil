package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetStats(x context.Context) (*client.Stats, error) {
	r, e := c.client.GetStatsWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get stats", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
