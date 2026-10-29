package mock_client

import "github.com/funtimecoding/soil/pkg/linkace/link"

func (c *Client) Links() ([]*link.Link, error) {
	return nil, c.Failure
}
