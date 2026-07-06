package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) FindPost(identifier string) (*model.Post, error) {
	result, _, e := c.client.GetPost(
		c.context,
		identifier,
		constant.EmptyEntityTag,
	)

	return result, e
}
