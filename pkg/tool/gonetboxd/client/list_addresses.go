package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListAddresses(
	x context.Context,
	name string,
) ([]*client.Address, error) {
	r, e := c.client.ListAddressesWithResponse(x, name)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list addresses", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
