package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) postPage(
	h *model.Channel,
	anchor string,
) (*model.PostList, error) {
	if anchor == "" {
		page, r, e := c.client.GetPostsForChannel(
			c.context,
			h.Id,
			0,
			constant.MattermostMaxPerPage,
			constant.MattermostEmptyEntityTag,
			constant.MattermostCollapsedThreads,
			false,
		)

		return page, wrapError(r, e)
	}

	page, r, e := c.client.GetPostsBefore(
		c.context,
		h.Id,
		anchor,
		0,
		constant.MattermostMaxPerPage,
		constant.MattermostEmptyEntityTag,
		constant.MattermostCollapsedThreads,
		false,
	)

	return page, wrapError(r, e)
}
