package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
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
		constant.MattermostEmptyEntityTag,
	)

	return result, wrapError(e)
}
