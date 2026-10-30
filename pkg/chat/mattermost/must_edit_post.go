package mattermost

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustEditPost(
	p *model.Post,
	message string,
) *model.Post {
	result, e := c.EditPost(p, message)
	errors.PanicOnError(e)

	return result
}
