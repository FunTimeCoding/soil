package netbox

func (c *Client) Close() {
	c.pool.Close()
}
