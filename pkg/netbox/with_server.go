package netbox

import "github.com/netbox-community/go-netbox/v4"

func WithServer(locator string) Option {
	return func(c *Client) {
		c.client.GetConfig().Servers = netbox.ServerConfigurations{
			{URL: locator},
		}
	}
}
