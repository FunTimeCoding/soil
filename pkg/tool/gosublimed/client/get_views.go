package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetViews(x context.Context) ([]client.View, error) {
	r, e := c.client.GetViewsWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get views", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
