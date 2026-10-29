package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListVirtualLabels(
	x context.Context,
	machine string,
) ([]*client.Label, error) {
	r, e := c.client.ListVirtualLabelsWithResponse(x, machine)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list virtual labels", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
