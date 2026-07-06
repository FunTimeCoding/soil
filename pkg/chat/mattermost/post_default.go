package mattermost

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) PostDefault(text string) (*model.Post, error) {
	if c.channel == nil {
		return nil, fmt.Errorf(
			"no default channel configured: %w",
			constant.ErrorNotConfigured,
		)
	}

	return c.PostSimple(c.channel, text)
}
