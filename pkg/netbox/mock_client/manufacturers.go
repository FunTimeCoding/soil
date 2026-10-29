package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/manufacturer"

func (c *Client) Manufacturers() ([]*manufacturer.Manufacturer, error) {
	return nil, nil
}
