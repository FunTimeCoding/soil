package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/site"

func (c *Client) Sites() ([]*site.Site, error) {
	return nil, nil
}
