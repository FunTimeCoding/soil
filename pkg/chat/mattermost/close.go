package mattermost

func (c *Client) Close() {
	c.pool.Close()
}
