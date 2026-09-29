package mock_client

import "github.com/funtimecoding/soil/pkg/linkace/tag"

func (c *Client) CreateTag(_ string) (*tag.Tag, error) {
	return nil, nil
}
