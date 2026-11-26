package mock_client

func (c *Client) Typed() []string {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return append([]string{}, c.typed...)
}
