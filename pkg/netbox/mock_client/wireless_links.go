package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/wireless_link"

func (c *Client) WirelessLinks() ([]*wireless_link.Link, error) {
	return nil, c.Failure
}
