package client

import "time"

func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		c.timeout = d
	}
}
