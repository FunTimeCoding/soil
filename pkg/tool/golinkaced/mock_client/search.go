package mock_client

import "github.com/funtimecoding/soil/pkg/linkace/link"

func (c *Client) Search(_ string) ([]*link.Link, error) {
	return nil, nil
}
