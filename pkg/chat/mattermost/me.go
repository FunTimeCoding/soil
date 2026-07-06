package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) Me() (*model.User, error) {
	if c.meCache != nil {
		return c.meCache, nil
	}

	result, _, e := c.client.GetMe(c.context, constant.EmptyEntityTag)

	if e != nil {
		return nil, e
	}

	c.meCache = result

	return c.meCache, nil
}
