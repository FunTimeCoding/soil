package netbox

import "github.com/funtimecoding/soil/pkg/netbox/source"

func (c *Client) Sources() ([]*source.Source, error) {
	result, _, e := c.client.CoreAPI.CoreDataSourcesList(
		c.context,
	).Execute()

	if e != nil {
		return nil, e
	}

	return source.NewSlice(result.Results), nil
}
