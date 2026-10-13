package web_client

func (c *Client) Close() {
	c.client.CloseIdleConnections()
}
