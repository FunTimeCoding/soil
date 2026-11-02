package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetGroup(x context.Context) ([]*client.Group, error) {
	r, e := c.client.GetGroupWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get group", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
