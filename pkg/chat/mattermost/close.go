package mattermost

func (c *Client) Close() {
	c.socketMutex.Lock()
	c.closed = true
	c.stopKeepAlive()

	if c.webSocket != nil {
		c.webSocket.Close()
		c.webSocket = nil
	}

	c.socketMutex.Unlock()
	c.pool.Close()
}
