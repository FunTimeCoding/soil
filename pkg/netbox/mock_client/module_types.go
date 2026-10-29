package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/module_type"

func (c *Client) ModuleTypes() ([]*module_type.Type, error) {
	return nil, c.Failure
}
