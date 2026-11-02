package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateView(
	x context.Context,
	b client.CreateViewJSONRequestBody,
) (*client.View, error) {
	r, e := c.client.CreateViewWithResponse(x, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("create view", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
