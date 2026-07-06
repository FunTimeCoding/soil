package mattermost

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustUpdateChannelHeader(
	h *model.Channel,
	prefix string,
	text string,
	separator string,
) *model.Channel {
	result, e := c.UpdateChannelHeader(h, prefix, text, separator)
	errors.PanicOnError(e)

	return result
}
