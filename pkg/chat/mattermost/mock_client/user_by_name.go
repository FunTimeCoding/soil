package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) UserByName(name string) (*model.User, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	for _, u := range c.user {
		if u.Username == name {
			return u, nil
		}
	}

	return nil, not_found.New("user", name)
}
