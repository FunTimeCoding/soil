package mock_client

func (c *Client) SetHostCommandStatus(status string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.hostCommandStatus = status
}
