package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/tag"

func (c *Client) Tags() ([]*tag.Tag, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}
