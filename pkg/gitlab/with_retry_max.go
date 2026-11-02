package gitlab

func WithRetryMax(retries int) Option {
	return func(c *Client) {
		c.retryMax = &retries
	}
}
