package nextcloud

func (c *Client) Close() {
	c.author.Close()
	c.basic.Close()
}
