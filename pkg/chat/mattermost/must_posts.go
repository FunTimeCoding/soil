package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) MustPosts(
	h *model.Channel,
	since time.Time,
) []*post.Post {
	result, e := c.Posts(h, since)
	errors.PanicOnError(e)

	return result
}
