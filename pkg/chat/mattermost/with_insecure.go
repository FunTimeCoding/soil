package mattermost

func WithInsecure() Option {
	return func(c *Client) {
		c.insecure = true
	}
}
