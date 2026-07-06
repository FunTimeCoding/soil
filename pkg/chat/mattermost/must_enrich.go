package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustEnrich(v []*post.Post) {
	errors.PanicOnError(c.Enrich(v))
}
