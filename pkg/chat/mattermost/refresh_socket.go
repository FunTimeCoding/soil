package mattermost

func (c *Client) RefreshSocket() error {
	if c.webSocket != nil {
		c.webSocket.Close()
	}

	s, e := newWebSocket(c.host, c.token, c.insecure)

	if e != nil {
		return e
	}

	c.webSocket = s

	return nil
}
