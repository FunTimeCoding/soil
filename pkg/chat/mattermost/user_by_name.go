package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) UserByName(name string) (*model.User, error) {
	result, r, e := c.client.GetUserByUsername(
		c.context,
		name,
		constant.MattermostEmptyEntityTag,
	)

	return result, wrapError(r, e)
}
