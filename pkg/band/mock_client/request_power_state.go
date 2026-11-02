package mock_client

func (c *Client) RequestPowerState(state int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.requested = append(c.requested, state)

	return nil
}
