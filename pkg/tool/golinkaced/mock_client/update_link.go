package mock_client

import "github.com/funtimecoding/soil/pkg/linkace/link"

func (c *Client) UpdateLink(
	_ int,
	_ map[string]any,
) (*link.Link, error) {
	return nil, nil
}
