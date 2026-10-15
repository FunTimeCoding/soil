package netbox

import "github.com/funtimecoding/soil/pkg/web/constant"

func WithPublicHost(host string) Option {
	return func(c *Client) {
		c.client.GetConfig().AddDefaultHeader(constant.ForwardedHost, host)
		c.client.GetConfig().AddDefaultHeader(
			constant.ForwardedProtocol,
			constant.Secure,
		)
	}
}
