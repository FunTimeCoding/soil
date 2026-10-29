package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) Me() (*model.User, error) {
	if c.meCache != nil {
		return c.meCache, nil
	}

	result, r, e := c.client.GetMe(c.context, constant.MattermostEmptyEntityTag)

	if e != nil {
		return nil, wrapError(r, e)
	}

	c.meCache = result

	return c.meCache, nil
}
