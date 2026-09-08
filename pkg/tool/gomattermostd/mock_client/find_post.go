package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) FindPost(identifier string) (*model.Post, error) {
	result, okay := c.post[identifier]

	if !okay {
		return nil, not_found.New("post", identifier)
	}

	return result, nil
}
