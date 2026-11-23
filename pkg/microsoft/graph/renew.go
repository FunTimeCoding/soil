package graph

func (c *Client) Renew() error {
	_, e := c.current(true)

	return e
}
