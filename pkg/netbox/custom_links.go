package netbox

import "github.com/funtimecoding/soil/pkg/netbox/custom_link"

func (c *Client) CustomLinks() ([]*custom_link.Link, error) {
	result, _, e := c.client.ExtrasAPI.ExtrasCustomLinksList(
		c.context,
	).Execute()

	if e != nil {
		return nil, e
	}

	return custom_link.NewSlice(result.Results), nil
}
