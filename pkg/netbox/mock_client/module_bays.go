package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/module_bay"

func (c *Client) ModuleBays() ([]*module_bay.Bay, error) {
	return nil, c.Failure
}
