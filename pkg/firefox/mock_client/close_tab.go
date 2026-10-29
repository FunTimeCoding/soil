package mock_client

func (c *Client) CloseTab(identifier int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if e := c.failures.Caller(); e != nil {
		return e
	}

	i, _, f := c.findTab(identifier)

	if f != nil {
		return f
	}

	c.tabs = append(c.tabs[:i], c.tabs[i+1:]...)

	return nil
}
