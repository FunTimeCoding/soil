package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetThread(
	x context.Context,
	identifier string,
) ([]*client.Post, error) {
	r, e := c.client.GetThreadWithResponse(x, identifier)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("thread", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
