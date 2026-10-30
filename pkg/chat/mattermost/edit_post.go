package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) EditPost(
	p *model.Post,
	message string,
) (*model.Post, error) {
	result, r, e := c.client.PatchPost(
		c.context,
		p.Id,
		&model.PostPatch{Message: &message},
	)
	f := wrapError(r, e)

	if not_found.Is(f) {
		return nil, not_found.New(constant.MattermostPostField, p.Id)
	}

	return result, f
}
