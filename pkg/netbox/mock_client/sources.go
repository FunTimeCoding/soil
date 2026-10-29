package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/source"

func (c *Client) Sources() ([]*source.Source, error) {
	return nil, c.Failure
}
