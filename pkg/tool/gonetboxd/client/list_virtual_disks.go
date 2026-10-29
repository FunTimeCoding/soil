package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListVirtualDisks(
	x context.Context,
	machine string,
) ([]*client.VirtualDisk, error) {
	r, e := c.client.ListVirtualDisksWithResponse(x, machine)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list virtual disks", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
