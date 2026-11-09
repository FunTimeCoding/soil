package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) ClearOld(
	h *model.Channel,
	cutoff time.Time,
	keep int,
) (int, int, error) {
	if c.Failure != nil {
		return 0, 0, c.Failure
	}

	var channel []*model.Post

	for _, p := range c.post {
		if p.ChannelId == h.Id {
			channel = append(channel, p)
		}
	}

	post.Sort(channel)
	var deleted int

	for i, p := range channel {
		newest := len(channel)-i <= keep

		if newest && !time.UnixMilli(p.CreateAt).Before(cutoff) {
			continue
		}

		delete(c.post, p.Id)
		deleted++
	}

	return deleted, 0, nil
}
