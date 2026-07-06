package mattermost

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustPostSimple(
	h *model.Channel,
	text string,
) *model.Post {
	result, e := c.PostSimple(h, text)
	errors.PanicOnError(e)

	return result
}
