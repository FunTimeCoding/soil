package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/configuration_template"

func (c *Client) ConfigurationTemplates() ([]*configuration_template.Template, error) {
	return nil, c.Failure
}
