package mattermost

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustUsers(
	page int,
	perPage int,
) []*model.User {
	result, e := c.Users(page, perPage)
	errors.PanicOnError(e)

	return result
}
