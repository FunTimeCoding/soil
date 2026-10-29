package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListVirtualInterfaces(
	x context.Context,
	name string,
) ([]*client.Interface, error) {
	r, e := c.client.ListVirtualInterfacesWithResponse(x, name)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(
			"list virtual interfaces",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
