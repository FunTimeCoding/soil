package proxmox

func (c *Client) Close() {
	c.pool.Close()
}
