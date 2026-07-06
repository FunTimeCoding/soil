package mattermost

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustPostHasReaction(
	p *model.Post,
	reaction string,
) bool {
	result, e := c.PostHasReaction(p, reaction)
	errors.PanicOnError(e)

	return result
}
