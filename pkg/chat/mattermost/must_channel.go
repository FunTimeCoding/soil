package mattermost

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustChannel(identifier string) *model.Channel {
	result, e := c.Channel(identifier)
	errors.PanicOnError(e)

	return result
}
