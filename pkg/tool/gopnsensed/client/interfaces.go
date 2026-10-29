package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) Interfaces(
	x context.Context,
) ([]client.NetworkInterface, error) {
	r, e := c.client.ListInterfacesWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("interfaces", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
