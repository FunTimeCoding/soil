package mattermost

import (
	"fmt"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) PostDefault(text string) (*model.Post, error) {
	if c.channel == nil {
		return nil, fmt.Errorf(
			"no default channel configured: %w",
			ErrorNotConfigured,
		)
	}

	return c.PostSimple(c.channel, text)
}
