package client

func WithLong() Option {
	return func(c *Client) {
		c.long = true
	}
}
