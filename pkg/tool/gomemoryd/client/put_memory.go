package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PutMemory(
	x context.Context,
	identifier int64,
	b client.PutMemoryJSONRequestBody,
) (*client.MemoryResponse, error) {
	r, e := c.client.PutMemoryWithResponse(x, identifier, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("put memory", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
