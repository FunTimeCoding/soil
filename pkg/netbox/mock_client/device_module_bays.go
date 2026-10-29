package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/module_bay"

func (c *Client) DeviceModuleBays(_ string) ([]*module_bay.Bay, error) {
	return nil, c.Failure
}
