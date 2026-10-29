package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) FindPost(identifier string) (*model.Post, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	result, okay := c.post[identifier]

	if !okay {
		return nil, not_found.New(constant.MattermostPostField, identifier)
	}

	return result, nil
}
