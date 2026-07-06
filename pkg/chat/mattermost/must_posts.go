package mattermost

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustPosts(h *model.Channel) *model.PostList {
	result, e := c.Posts(h)
	errors.PanicOnError(e)

	return result
}
