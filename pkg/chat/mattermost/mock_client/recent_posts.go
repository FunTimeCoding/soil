package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) RecentPosts(
	h *model.Channel,
	_ int,
) ([]*post.Post, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	var raw []*model.Post

	for _, p := range c.post {
		if p.ChannelId == h.Id {
			raw = append(raw, p)
		}
	}

	post.Sort(raw)

	return post.NewSlice(raw), nil
}
