package mock_client

func (c *Client) Lookups() []string {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return append([]string{}, c.lookups...)
}
