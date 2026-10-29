package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"
)

func (c *Client) GetChannelPosts(
	_ context.Context,
	name string,
	p *client.GetChannelPostsParams,
) ([]*client.Post, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	posts, okay := c.Posts[name]

	if !okay {
		return nil, not_found.New("channel", name)
	}

	var result []*client.Post

	for _, v := range posts {
		if !v.CreatedAt.Before(p.Since) {
			result = append(result, v)
		}
	}

	return result, nil
}
