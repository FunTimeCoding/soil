package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) Posts(h *model.Channel) (*model.PostList, error) {
	result, _, e := c.client.GetPostsForChannel(
		c.context,
		h.Id,
		0,
		constant.PerPage,
		constant.EmptyEntityTag,
		true,
		false,
	)

	return result, e
}
