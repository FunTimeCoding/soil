package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListContainers(
	x context.Context,
	p *client.ListContainersParams,
) ([]client.Container, error) {
	r, e := c.client.ListContainersWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list containers", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
