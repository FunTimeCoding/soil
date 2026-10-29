package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) Channels(
	t *model.Team,
	u *model.User,
) ([]*model.Channel, error) {
	result, r, e := c.client.GetChannelsForTeamForUser(
		c.context,
		t.Id,
		u.Id,
		false,
		constant.MattermostEmptyEntityTag,
	)

	return result, wrapError(r, e)
}
