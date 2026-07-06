package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) MustPostBefore(
	h *model.Channel,
	t time.Time,
) *post.Post {
	result, e := c.PostBefore(h, t)
	errors.PanicOnError(e)

	return result
}
