package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) Clear(h *model.Channel) (int, error) {
	var deleted int

	for range constant.MattermostClearBatchLimit {
		page, e := c.postPage(h, "")

		if e != nil {
			return deleted, e
		}

		posts := post.FromList(page, false)

		if len(posts) == 0 {
			return deleted, nil
		}

		for _, p := range posts {
			f := c.DeletePost(p)

			if not_found.Is(f) {
				continue
			}

			if f != nil {
				return deleted, f
			}

			deleted++
		}
	}

	return deleted, nil
}
