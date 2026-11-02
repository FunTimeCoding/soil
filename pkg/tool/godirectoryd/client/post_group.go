package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PostGroup(
	x context.Context,
	b client.PostGroupJSONRequestBody,
) (*client.Group, error) {
	r, e := c.client.PostGroupWithResponse(x, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("post group", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
