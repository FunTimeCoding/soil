package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListNodes(
	x context.Context,
	p *client.ListNodesParams,
) ([]client.Node, error) {
	r, e := c.client.ListNodesWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list nodes", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
