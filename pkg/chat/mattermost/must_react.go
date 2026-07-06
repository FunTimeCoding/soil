package mattermost

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustReact(
	p *model.Post,
	emoji string,
) {
	errors.PanicOnError(c.React(p, emoji))
}
