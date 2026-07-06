package netbox

import "github.com/funtimecoding/soil/pkg/netbox/system_number"

func (c *Client) SystemNumbers() ([]*system_number.Number, error) {
	result, _, e := c.client.IpamAPI.IpamAsnsList(
		c.context,
	).Execute()

	if e != nil {
		return nil, e
	}

	return system_number.NewSlice(result.Results), nil
}
