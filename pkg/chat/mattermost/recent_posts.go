package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) RecentPosts(
	h *model.Channel,
	limit int,
) ([]*post.Post, error) {
	if limit <= 0 || limit > constant.MattermostMaxPerPage {
		limit = constant.MattermostMaxPerPage
	}

	page, r, e := c.client.GetPostsForChannel(
		c.context,
		h.Id,
		0,
		limit,
		constant.MattermostEmptyEntityTag,
		constant.MattermostCollapsedThreads,
		false,
	)

	if e != nil {
		return nil, wrapError(r, e)
	}

	result := post.NewSlice(post.FromList(page, true))

	if f := c.Enrich(result); f != nil {
		return nil, f
	}

	return result, nil
}
