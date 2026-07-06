package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) Team(name string) (*model.Team, error) {
	result, _, e := c.client.GetTeamByName(
		c.context,
		name,
		constant.EmptyEntityTag,
	)

	return result, e
}
