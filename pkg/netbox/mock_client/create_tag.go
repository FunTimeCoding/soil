package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/tag"

func (c *Client) CreateTag(_ string) (*tag.Tag, error) {
	return nil, nil
}
