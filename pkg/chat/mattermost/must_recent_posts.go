package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustRecentPosts(
	h *model.Channel,
	limit int,
) []*post.Post {
	result, e := c.RecentPosts(h, limit)
	errors.PanicOnError(e)

	return result
}
