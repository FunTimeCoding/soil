package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) Log(
	x context.Context,
	limit int,
) ([]client.LogEntry, error) {
	params := &client.FirewallLogParams{}

	if limit > 0 {
		params.Limit = &limit
	}

	r, e := c.client.FirewallLogWithResponse(x, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("firewall log", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
