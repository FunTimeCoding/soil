package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListMachines(
	x context.Context,
	p *client.ListMachinesParams,
) ([]client.Machine, error) {
	r, e := c.client.ListMachinesWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list machines", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
