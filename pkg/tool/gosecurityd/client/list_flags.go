package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListFlags(
	x context.Context,
	p *client.ListFlagsParams,
) ([]*client.Flag, error) {
	r, e := c.client.ListFlagsWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("flags", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
