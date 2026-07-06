package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/constant"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) PostsBefore(
	h *model.Channel,
	t time.Time,
	keep int,
) ([]*post.Post, error) {
	const maximumDelete = 1000
	var result []*post.Post
	position := 0
	pageNumber := 0

	for len(result) < maximumDelete {
		page, _, e := c.client.GetPostsForChannel(
			c.context,
			h.Id,
			pageNumber,
			constant.PerPage,
			constant.EmptyEntityTag,
			false,
			false,
		)

		if e != nil {
			return nil, e
		}

		if len(page.Order) == 0 {
			break
		}

		wrapped := post.NewSlice(post.FromList(page, false))
		f := c.Enrich(wrapped)

		if f != nil {
			return nil, f
		}

		for _, v := range wrapped {
			collect := position >= keep || v.Create.Before(t)

			if collect {
				result = append(result, v)

				if len(result) >= maximumDelete {
					return result, nil
				}
			}

			position++
		}

		pageNumber++
	}

	return result, nil
}
