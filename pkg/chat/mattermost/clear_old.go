package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) ClearOld(
	h *model.Channel,
	cutoff time.Time,
	keep int,
) (int, int, error) {
	var deleted int
	var skipped int

	for range constant.MattermostClearBatchLimit {
		posts, e := c.oldPosts(h, cutoff, keep)

		if e != nil {
			return deleted, skipped, e
		}

		if len(posts) == 0 {
			return deleted, skipped, nil
		}

		for _, p := range posts {
			f := c.DeletePost(p.Raw)

			if not_found.Is(f) {
				skipped++

				continue
			}

			if f != nil {
				return deleted, skipped, f
			}

			deleted++
		}
	}

	return deleted, skipped, nil
}
