package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) oldPosts(
	h *model.Channel,
	cutoff time.Time,
	keep int,
) ([]*post.Post, error) {
	var result []*post.Post
	position := 0
	anchor := ""

	for len(result) < constant.MattermostClearBatchSize {
		page, e := c.postPage(h, anchor)

		if e != nil {
			return nil, e
		}

		if len(page.Order) == 0 {
			break
		}

		for _, v := range post.NewSlice(post.FromList(page, false)) {
			if position >= keep || v.Create.Before(cutoff) {
				result = append(result, v)

				if len(result) >= constant.MattermostClearBatchSize {
					return result, nil
				}
			}

			position++
		}

		anchor = page.Order[len(page.Order)-1]
	}

	return result, nil
}
