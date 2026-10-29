package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetChannelPosts(
	x context.Context,
	name string,
	p *client.GetChannelPostsParams,
) ([]*client.Post, error) {
	r, e := c.client.GetChannelPostsWithResponse(x, name, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("channel posts", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
