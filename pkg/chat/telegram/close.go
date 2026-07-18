package telegram

func (c *Client) Close() {
	if c.store != nil {
		c.store.Close()
	}
}
