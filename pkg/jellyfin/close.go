package jellyfin

func (c *Client) Close() {
	c.basic.Close()
}
