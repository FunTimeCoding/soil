package sentry

func (c *Client) Close() {
	c.basic.Close()
}
