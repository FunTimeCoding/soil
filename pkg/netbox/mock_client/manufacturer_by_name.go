package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/manufacturer"

func (c *Client) ManufacturerByName(_ string) (*manufacturer.Manufacturer, error) {
	return nil, nil
}
