package mattermost

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustReactions(p *model.Post) []*model.Reaction {
	result, e := c.Reactions(p)
	errors.PanicOnError(e)

	return result
}
