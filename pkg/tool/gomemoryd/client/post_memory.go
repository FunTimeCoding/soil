package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PostMemory(
	x context.Context,
	b client.PostMemoryJSONRequestBody,
) (*client.MemoryResponse, error) {
	r, e := c.client.PostMemoryWithResponse(x, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("post memory", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
