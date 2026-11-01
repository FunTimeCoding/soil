package mock_client

import "github.com/funtimecoding/soil/pkg/github/user"

func (c *Client) User() (*user.User, error) {
	return nil, c.Failure
}
