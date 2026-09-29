package mock_client

import "github.com/funtimecoding/soil/pkg/linkace/list"

func (c *Client) CreateList(
	_ string,
	_ string,
) (*list.List, error) {
	return nil, nil
}
