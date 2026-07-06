package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) Channels(
	t *model.Team,
	u *model.User,
) ([]*model.Channel, error) {
	result, _, e := c.client.GetChannelsForTeamForUser(
		c.context,
		t.Id,
		u.Id,
		false,
		constant.EmptyEntityTag,
	)

	return result, e
}
