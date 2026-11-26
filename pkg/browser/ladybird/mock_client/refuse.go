package mock_client

func (c *Client) Refuse(
	declaration string,
	e error,
) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.refusals[declaration] = e
}
