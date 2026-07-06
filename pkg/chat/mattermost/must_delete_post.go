package mattermost

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustDeletePost(p *model.Post) {
	errors.PanicOnError(c.DeletePost(p))
}
