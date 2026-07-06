package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) ChannelByName(
	t *model.Team,
	name string,
) (*model.Channel, error) {
	result, _, e := c.client.GetChannelByName(
		c.context,
		name,
		t.Id,
		constant.EmptyEntityTag,
	)

	return result, e
}
