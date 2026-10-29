package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/module"

func (c *Client) Modules() ([]*module.Module, error) {
	return nil, c.Failure
}
