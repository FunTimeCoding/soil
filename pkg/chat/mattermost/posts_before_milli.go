package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) PostsBeforeMilli(
	h *model.Channel,
	beforeMilli int64,
	keep int,
) ([]*post.Post, error) {
	return c.PostsBefore(h, time.UnixMilli(beforeMilli), keep)
}
