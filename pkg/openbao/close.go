package openbao

func (c *Client) Close() {
	c.pool.Close()
}
