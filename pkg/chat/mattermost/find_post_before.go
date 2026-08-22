package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) FindPostBefore(
	h *model.Channel,
	t time.Time,
) (*post.Post, bool, error) {
	result, e := c.PostBefore(h, t)

	if e != nil {
		if not_found.Is(e) {
			return nil, false, nil
		}

		return nil, false, e
	}

	return result, true, nil
}
