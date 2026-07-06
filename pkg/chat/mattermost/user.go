package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) User(identifier string) (*model.User, error) {
	if c.user.HasIdentifier(identifier) {
		return c.user.ByIdentifier(identifier), nil
	}

	result, _, e := c.client.GetUser(
		c.context,
		identifier,
		constant.EmptyEntityTag,
	)

	if e != nil {
		return nil, e
	}

	c.user.Add(result)

	return result, nil
}
