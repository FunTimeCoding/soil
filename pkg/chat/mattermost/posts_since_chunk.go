package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) postsSinceChunk(
	h *model.Channel,
	since time.Time,
) ([]*model.Post, error) {
	// collapsedThreads must stay false: the collapsed view returns
	// only thread roots, so since-scoped reads would miss every reply.
	list, _, e := c.client.GetPostsSince(
		c.context,
		h.Id,
		since.UnixMilli(),
		false,
	)

	if e != nil {
		return nil, e
	}

	return post.FromList(list, true), nil
}
