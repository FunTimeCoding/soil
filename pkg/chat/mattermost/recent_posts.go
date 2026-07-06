package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) RecentPosts(
	h *model.Channel,
	sinceMilli int64,
) ([]*post.Post, error) {
	list, _, e := c.client.GetPostsSince(
		c.context,
		h.Id,
		sinceMilli,
		true,
	)

	if e != nil {
		return nil, e
	}

	t := time.UnixMilli(sinceMilli)
	var posts []*model.Post

	for _, v := range post.FromList(list, true) {
		if v.Message == "" {
			continue
		}

		if time.UnixMilli(v.CreateAt).Before(t) {
			continue
		}

		posts = append(posts, v)
	}

	result := post.NewSlice(posts)
	f := c.Enrich(result)

	if f != nil {
		return nil, f
	}

	return result, nil
}
