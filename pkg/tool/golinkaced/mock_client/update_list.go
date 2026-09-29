package mock_client

import "github.com/funtimecoding/soil/pkg/linkace/list"

func (c *Client) UpdateList(
	_ int,
	_ map[string]any,
) (*list.List, error) {
	return nil, nil
}
