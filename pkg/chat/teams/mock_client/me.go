package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/teams/user"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) Me() (*user.User, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if c.me == nil {
		return nil, not_found.New("user", "me")
	}

	return c.me, nil
}
