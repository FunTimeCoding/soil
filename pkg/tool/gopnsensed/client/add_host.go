package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) AddHost(
	x context.Context,
	h client.HostRequest,
	apply bool,
) (*client.Identifier, error) {
	r, e := c.client.AddHostWithResponse(
		x,
		&client.AddHostParams{Apply: &apply},
		h,
	)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("add host", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
