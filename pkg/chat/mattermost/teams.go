package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) Teams(userIdentifier string) ([]*model.Team, error) {
	result, _, e := c.client.GetTeamsForUser(
		c.context,
		userIdentifier,
		constant.MattermostEmptyEntityTag,
	)

	return result, e
}
