package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListCables(x context.Context) ([]*client.Cable, error) {
	r, e := c.client.ListCablesWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list cables", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
