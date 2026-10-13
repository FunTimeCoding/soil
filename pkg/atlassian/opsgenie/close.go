package opsgenie

func (c *Client) Close() {
	c.userClient.Close()

	if c.teamClient != nil {
		c.teamClient.Close()
	}
}
