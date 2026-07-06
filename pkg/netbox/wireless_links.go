package netbox

import "github.com/funtimecoding/soil/pkg/netbox/wireless_link"

func (c *Client) WirelessLinks() ([]*wireless_link.Link, error) {
	result, _, e := c.client.WirelessAPI.WirelessWirelessLinksList(
		c.context,
	).Execute()

	if e != nil {
		return nil, e
	}

	return wireless_link.NewSlice(result.Results), nil
}
