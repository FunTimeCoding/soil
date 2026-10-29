package ollama

func WithEmbedModel(s string) Option {
	return func(c *Client) {
		c.model = s
	}
}
