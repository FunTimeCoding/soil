package teams

func (c *Client) Close() {
	c.graph.Close()
}
