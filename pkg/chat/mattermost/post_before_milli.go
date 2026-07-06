package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) PostBeforeMilli(
	h *model.Channel,
	beforeMilli int64,
) (*post.Post, error) {
	return c.PostBefore(h, time.UnixMilli(beforeMilli))
}
