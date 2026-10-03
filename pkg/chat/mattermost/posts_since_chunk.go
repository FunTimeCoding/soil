package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) postsSinceChunk(
	h *model.Channel,
	since time.Time,
) ([]*model.Post, error) {
	list, _, e := c.client.GetPostsSince(
		c.context,
		h.Id,
		since.UnixMilli(),
		constant.MattermostCollapsedThreads,
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return post.FromList(list, true), nil
}
