package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListPhysicalAddresses(
	x context.Context,
) ([]client.PhysicalAddressOwner, error) {
	r, e := c.client.ListPhysicalAddressesWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(
			"list physical addresses",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
