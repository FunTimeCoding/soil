package gmail

func (c *Client) Close() {
	c.callback.Close()
}
