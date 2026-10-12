package mock_client

func (c *Client) SetPower(
	state int,
	available ...int,
) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.state = state
	c.available = available
}
