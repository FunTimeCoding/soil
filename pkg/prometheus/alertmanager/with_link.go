package alertmanager

import "github.com/funtimecoding/soil/pkg/web/locator"

func WithLink(
	host string,
	port int,
	secure bool,
) Option {
	return func(c *Client) {
		c.link = locator.NewEndpoint(host, port, secure)
	}
}
