package netbox

import webConstant "github.com/funtimecoding/soil/pkg/web/constant"

func WithPublicHost(host string) Option {
	return func(c *Client) {
		c.client.GetConfig().AddDefaultHeader(webConstant.ForwardedHost, host)
		c.client.GetConfig().AddDefaultHeader(
			webConstant.ForwardedProtocol,
			webConstant.Secure,
		)
	}
}
