package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) Posts(h *model.Channel) (*model.PostList, error) {
	result, _, e := c.client.GetPostsForChannel(
		c.context,
		h.Id,
		0,
		constant.MattermostPerPage,
		constant.MattermostEmptyEntityTag,
		true,
		false,
	)

	return result, e
}
