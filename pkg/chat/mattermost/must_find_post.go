package mattermost

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustFindPost(identifier string) *model.Post {
	result, e := c.FindPost(identifier)
	errors.PanicOnError(e)

	return result
}
