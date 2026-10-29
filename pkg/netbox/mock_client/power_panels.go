package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/power_panel"

func (c *Client) PowerPanels() ([]*power_panel.Panel, error) {
	return nil, c.Failure
}
