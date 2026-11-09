package mattermost

import (
	"github.com/mattermost/mattermost/server/public/model"
	"strings"
)

func (c *Client) WipePrefix(
	h *model.Channel,
	prefix string,
) {
	me := c.MustMe()

	for _, p := range c.MustRecentPosts(h, 0) {
		if !strings.HasPrefix(p.Raw.Message, prefix) || p.Raw.UserId != me.Id {
			continue
		}

		c.MustDeletePost(p.Raw)
	}
}
