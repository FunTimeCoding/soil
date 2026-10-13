package loki

func (c *Client) Close() {
	c.basic.Close()
}
