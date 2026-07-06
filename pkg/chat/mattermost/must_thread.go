package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustThread(p *model.Post) []*post.Post {
	result, e := c.Thread(p)
	errors.PanicOnError(e)

	return result
}
