package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/constant"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
)

// LatestPosts fetches the newest posts for display, returned oldest
// first. Not to be confused with PostsBefore, which skips the newest
// posts and collects older ones for deletion workflows.
func (c *Client) LatestPosts(
	h *model.Channel,
	limit int,
) ([]*post.Post, error) {
	if limit <= 0 || limit > constant.PerPage {
		limit = constant.PerPage
	}

	page, _, e := c.client.GetPostsForChannel(
		c.context,
		h.Id,
		0,
		limit,
		constant.EmptyEntityTag,
		false,
		false,
	)

	if e != nil {
		return nil, e
	}

	result := post.NewSlice(post.FromList(page, true))
	f := c.Enrich(result)

	if f != nil {
		return nil, f
	}

	return result, nil
}
