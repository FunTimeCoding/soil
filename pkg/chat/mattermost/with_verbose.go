package mattermost

func WithVerbose(v bool) Option {
	return func(c *Client) {
		c.verbose = v
	}
}
