package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) DeriveHardwareAddress(
	x context.Context,
	p *client.DeriveHardwareAddressParams,
) (*client.DerivedAddress, error) {
	r, e := c.client.DeriveHardwareAddressWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(
			"derive hardware address",
			r.HTTPResponse,
			r.Body,
		)
	}

	return r.JSON200, nil
}
