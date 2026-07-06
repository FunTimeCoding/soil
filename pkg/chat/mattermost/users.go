package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) Users(
	page int,
	perPage int,
) ([]*model.User, error) {
	result, _, e := c.client.GetUsers(
		c.context,
		page,
		perPage,
		constant.EmptyEntityTag,
	)

	return result, e
}
