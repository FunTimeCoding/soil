package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetUser(x context.Context) ([]*client.User, error) {
	r, e := c.client.GetUserWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get user", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
