package mock_client

func (c *Client) UngroupTab(identifier int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if e := c.failures.Caller(); e != nil {
		return e
	}

	_, t, f := c.findTab(identifier)

	if f != nil {
		return f
	}

	t.GroupIdentifier = 0

	return nil
}
