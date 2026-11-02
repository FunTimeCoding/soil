package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PostUser(
	x context.Context,
	b client.PostUserJSONRequestBody,
) (*client.User, error) {
	r, e := c.client.PostUserWithResponse(x, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("post user", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
