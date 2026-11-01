package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) DeleteMemory(
	x context.Context,
	identifier int64,
	p *client.DeleteMemoryParams,
) (*client.MemoryResponse, error) {
	r, e := c.client.DeleteMemoryWithResponse(x, identifier, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("delete memory", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
