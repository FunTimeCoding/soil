package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/manufacturer"

func (c *Client) Manufacturers() ([]*manufacturer.Manufacturer, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}
