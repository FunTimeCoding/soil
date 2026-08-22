package mattermost

import (
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) PostDefault(text string) (*model.Post, error) {
	if c.channel == nil {
		return nil, not_configured.Format("no default channel configured")
	}

	return c.PostSimple(c.channel, text)
}
