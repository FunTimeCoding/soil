package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListPrefixAddresses(
	x context.Context,
	prefix string,
) ([]*client.Address, error) {
	r, e := c.client.ListPrefixAddressesWithResponse(
		x,
		&client.ListPrefixAddressesParams{Prefix: prefix},
	)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list prefix addresses", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
