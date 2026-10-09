package mattermost

import "time"

func (c *Client) startKeepAlive() {
	if c.keepAlive <= 0 {
		return
	}

	c.stopPing = make(chan struct{})
	go c.ping(c.webSocket, c.stopPing, time.NewTicker(c.keepAlive))
}
