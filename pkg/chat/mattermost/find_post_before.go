package mattermost

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) FindPostBefore(
	h *model.Channel,
	t time.Time,
) (*post.Post, bool, error) {
	result, e := c.PostBefore(h, t)

	if e != nil {
		if errors.Is(e, constant.ErrorMattermostNotFound) {
			return nil, false, nil
		}

		return nil, false, e
	}

	return result, true, nil
}
