package closing_source

func (c *Client) Closed() int {
	return c.closed
}
