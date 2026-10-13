package jira

func (c *Client) Close() {
	c.basic.Close()
	c.pool.Close()
}
