package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) DeletePost(p *model.Post) error {
	r, e := c.client.DeletePost(c.context, p.Id)
	f := wrapError(r, e)

	if not_found.Is(f) {
		return not_found.New(constant.MattermostPostField, p.Id)
	}

	return f
}
