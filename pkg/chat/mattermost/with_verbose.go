package mattermost

import "github.com/funtimecoding/soil/pkg/web/verbose_transport"

func WithVerbose(v bool) Option {
	return func(c *Client) {
		if v {
			c.client.HTTPClient = verbose_transport.New(true)
		}
	}
}
