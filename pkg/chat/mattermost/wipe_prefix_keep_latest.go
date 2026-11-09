package mattermost

import (
	"github.com/mattermost/mattermost/server/public/model"
	"strings"
)

func (c *Client) WipePrefixKeepLatest(
	h *model.Channel,
	prefix string,
) {
	me := c.MustMe()
	var candidates []*model.Post

	for _, p := range c.MustRecentPosts(h, 0) {
		if strings.HasPrefix(p.Raw.Message, prefix) && p.Raw.UserId == me.Id {
			candidates = append(candidates, p.Raw)
		}
	}

	for i := 0; i < len(candidates)-1; i++ {
		c.MustDeletePost(candidates[i])
	}
}
