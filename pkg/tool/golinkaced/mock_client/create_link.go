package mock_client

import "github.com/funtimecoding/soil/pkg/linkace/link"

func (c *Client) CreateLink(
	_ string,
	_ string,
	_ int,
	_ []string,
) (*link.Link, error) {
	return nil, nil
}
