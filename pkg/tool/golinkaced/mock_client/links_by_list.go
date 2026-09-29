package mock_client

import "github.com/funtimecoding/soil/pkg/linkace/link"

func (c *Client) LinksByList(_ int) ([]*link.Link, error) {
	return nil, nil
}
