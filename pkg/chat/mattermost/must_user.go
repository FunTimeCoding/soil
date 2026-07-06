package mattermost

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustUser(identifier string) *model.User {
	result, e := c.User(identifier)
	errors.PanicOnError(e)

	return result
}
