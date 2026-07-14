package mattermost

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustAllChannels(
	limit int,
	offset int,
) []*model.ChannelWithTeamData {
	result, e := c.AllChannels(limit, offset)
	errors.PanicOnError(e)

	return result
}
