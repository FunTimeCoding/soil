package raid_parser

func (c *Client) Close() {
	c.pool.Close()
}
