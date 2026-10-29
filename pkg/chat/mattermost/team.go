package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) Team(name string) (*model.Team, error) {
	result, r, e := c.client.GetTeamByName(
		c.context,
		name,
		constant.MattermostEmptyEntityTag,
	)

	return result, wrapError(r, e)
}
