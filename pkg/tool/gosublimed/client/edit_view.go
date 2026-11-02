package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) EditView(
	x context.Context,
	identifier int,
	b client.EditViewJSONRequestBody,
) (*client.View, error) {
	r, e := c.client.EditViewWithResponse(x, identifier, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("edit view", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
