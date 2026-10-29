package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/service_template"

func (c *Client) ServiceTemplates() ([]*service_template.Template, error) {
	return nil, c.Failure
}
