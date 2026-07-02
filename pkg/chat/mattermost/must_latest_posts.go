package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/post"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustLatestPosts(
	h *model.Channel,
	limit int,
) []*post.Post {
	result, e := c.LatestPosts(h, limit)
	errors.PanicOnError(e)

	return result
}
