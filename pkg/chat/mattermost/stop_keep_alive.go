package mattermost

func (c *Client) stopKeepAlive() {
	if c.stopPing == nil {
		return
	}

	close(c.stopPing)
	c.stopPing = nil
}
