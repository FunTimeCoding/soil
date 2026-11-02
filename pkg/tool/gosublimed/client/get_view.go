package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetView(
	x context.Context,
	identifier int,
) (*client.View, error) {
	r, e := c.client.GetViewWithResponse(x, identifier)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get view", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
