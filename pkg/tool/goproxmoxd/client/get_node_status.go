package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetNodeStatus(
	x context.Context,
	node string,
	p *client.GetNodeStatusParams,
) (*client.NodeStatus, error) {
	r, e := c.client.GetNodeStatusWithResponse(x, node, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get node status", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
