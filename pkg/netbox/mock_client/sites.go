package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/site"

func (c *Client) Sites() ([]*site.Site, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}
