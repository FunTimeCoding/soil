package mattermost

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustMe() *model.User {
	result, e := c.Me()
	errors.PanicOnError(e)

	return result
}
