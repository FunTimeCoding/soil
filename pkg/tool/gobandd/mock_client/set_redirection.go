package mock_client

func (c *Client) SetRedirection(
	listener bool,
	screenService int,
) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.listener = listener
	c.screenService = screenService
}
