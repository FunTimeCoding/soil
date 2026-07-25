package mattermost

func (c *Client) RefreshSocket() {
	if c.webSocket != nil {
		c.webSocket.Close()
	}

	c.webSocket = newWebSocket(c.host, c.token, c.insecure)
}
