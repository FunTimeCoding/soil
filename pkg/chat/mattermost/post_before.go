package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/time/constant"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) PostBefore(
	h *model.Channel,
	t time.Time,
) (*post.Post, error) {
	anchor := ""

	for {
		page, e := c.postPage(h, anchor, true)

		if e != nil {
			return nil, e
		}

		if len(page.Order) == 0 {
			return nil, not_found.Format(
				"no post before %s in channel %s",
				t.Format(constant.DateMinute),
				h.Id,
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

		anchor = page.Order[len(page.Order)-1]
	}
}
