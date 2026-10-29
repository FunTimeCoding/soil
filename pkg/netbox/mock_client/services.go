package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/service"

func (c *Client) Services() ([]*service.Service, error) {
	return nil, c.Failure
}
