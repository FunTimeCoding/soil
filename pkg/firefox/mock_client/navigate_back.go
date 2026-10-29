package mock_client

func (c *Client) NavigateBack(identifier int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if e := c.failures.Caller(); e != nil {
		return e
	}

	_, _, f := c.findTab(identifier)

	return f
}
