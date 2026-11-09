package mock_client

import "github.com/funtimecoding/soil/pkg/linkace/list"

func (c *Client) List(_ string) (*list.List, error) {
	return nil, c.Failure
}
