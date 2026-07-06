package mattermost

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustAllChannels() []*model.ChannelWithTeamData {
	result, e := c.AllChannels()
	errors.PanicOnError(e)

	return result
}
