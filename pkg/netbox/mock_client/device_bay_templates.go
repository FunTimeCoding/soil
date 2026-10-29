package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device_bay_template"

func (c *Client) DeviceBayTemplates() ([]*device_bay_template.Template, error) {
	return nil, c.Failure
}
