package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListStorages(
	x context.Context,
	node string,
	p *client.ListStoragesParams,
) ([]client.Storage, error) {
	r, e := c.client.ListStoragesWithResponse(x, node, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list storages", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
