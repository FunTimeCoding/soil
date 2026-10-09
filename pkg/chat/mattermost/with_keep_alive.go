package mattermost

import "time"

func WithKeepAlive(d time.Duration) Option {
	return func(c *Client) {
		c.keepAlive = d
	}
}
