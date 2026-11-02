package mattermost

import "context"

func (c *Client) RefreshSocket(x context.Context) error {
	c.socketMutex.Lock()
	defer c.socketMutex.Unlock()

	if c.closed {
		return errorClosed()
	}

	if c.webSocket != nil {
		c.stopKeepAlive()
		c.webSocket.Close()
	}

	s, e := newWebSocket(x, c.host, c.token, c.insecure)

	if e != nil {
		return e
	}

	c.webSocket = s
	c.startKeepAlive()

	return nil
}
