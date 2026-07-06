package mattermost

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustChannelByName(
	t *model.Team,
	name string,
) *model.Channel {
	result, e := c.ChannelByName(t, name)
	errors.PanicOnError(e)

	return result
}
