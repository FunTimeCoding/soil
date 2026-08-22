package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) postPage(
	h *model.Channel,
	anchor string,
	collapsed bool,
) (*model.PostList, error) {
	if anchor == "" {
		page, _, e := c.client.GetPostsForChannel(
			c.context,
			h.Id,
			0,
			constant.MattermostMaxPerPage,
			constant.MattermostEmptyEntityTag,
			collapsed,
			false,
		)

		return page, wrapError(e)
	}

	page, _, e := c.client.GetPostsBefore(
		c.context,
		h.Id,
		anchor,
		0,
		constant.MattermostMaxPerPage,
		constant.MattermostEmptyEntityTag,
		collapsed,
		false,
	)

	return page, wrapError(e)
}
