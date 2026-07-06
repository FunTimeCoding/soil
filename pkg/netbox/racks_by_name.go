package netbox

import "github.com/funtimecoding/soil/pkg/netbox/rack"

func (c *Client) RacksByName(n string) ([]*rack.Rack, error) {
	result, _, e := c.client.DcimAPI.DcimRacksList(
		c.context,
	).Name([]string{n}).Execute()

	if e != nil {
		return nil, e
	}

	return rack.NewSlice(result.Results), nil
}
