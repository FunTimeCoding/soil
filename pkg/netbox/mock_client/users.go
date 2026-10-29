package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/user"

func (c *Client) Users() ([]*user.User, error) {
	return nil, c.Failure
}
