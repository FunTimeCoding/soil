package mock_client

func (c *Client) Triggered() []string {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return append([]string{}, c.triggered...)
}
