package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/export_template"

func (c *Client) ExportTemplates() ([]*export_template.Template, error) {
	return nil, c.Failure
}
