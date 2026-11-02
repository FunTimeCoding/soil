package counting_mattermost

func (c *Client) Mes() int {
	return int(c.me.Load())
}
