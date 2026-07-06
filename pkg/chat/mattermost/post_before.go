package mattermost

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/constant"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	library "github.com/funtimecoding/soil/pkg/time"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) PostBefore(
	h *model.Channel,
	t time.Time,
) (*post.Post, error) {
	pageNumber := 0

	for {
		page, _, e := c.client.GetPostsForChannel(
			c.context,
			h.Id,
			pageNumber,
			constant.PerPage,
			constant.EmptyEntityTag,
			true,
			false,
		)

		if e != nil {
			return nil, e
		}

		if len(page.Order) == 0 {
			return nil, fmt.Errorf(
				"no post before %s in channel %s: %w",
				t.Format(library.DateMinute),
				h.Id,
				constant.ErrorNotFound,
			)
		}

		wrapped := post.NewSlice(post.FromList(page, false))
		f := c.Enrich(wrapped)

		if f != nil {
			return nil, f
		}

		for _, v := range wrapped {
			if v.Create.Before(t) {
				return v, nil
			}
		}

		pageNumber++
	}
}
