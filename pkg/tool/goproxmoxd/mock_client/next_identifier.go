package mock_client

func (c *Client) NextIdentifier() (int, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	result := c.nextIdentifier
	c.nextIdentifier++

	return result, nil
}
