package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustRecentPosts(
	h *model.Channel,
	sinceMilli int64,
) []*post.Post {
	result, e := c.RecentPosts(h, sinceMilli)
	errors.PanicOnError(e)

	return result
}
