package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustPostBeforeMilli(
	h *model.Channel,
	beforeMilli int64,
) *post.Post {
	result, e := c.PostBeforeMilli(h, beforeMilli)
	errors.PanicOnError(e)

	return result
}
