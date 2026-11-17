package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListPowerOutlets(
	x context.Context,
	name string,
) ([]*client.PowerOutlet, error) {
	r, e := c.client.ListPowerOutletsWithResponse(x, name)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list power outlets", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
