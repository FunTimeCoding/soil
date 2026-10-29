package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/custom_link"

func (c *Client) CustomLinks() ([]*custom_link.Link, error) {
	return nil, c.Failure
}
