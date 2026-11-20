package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListSystems(x context.Context) ([]*client.SystemSummary, error) {
	r, e := c.client.ListSystemsWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("systems", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
