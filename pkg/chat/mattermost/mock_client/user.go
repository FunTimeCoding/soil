package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) User(identifier string) (*model.User, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	result, okay := c.user[identifier]

	if !okay {
		return nil, not_found.New("user", identifier)
	}

	return result, nil
}
