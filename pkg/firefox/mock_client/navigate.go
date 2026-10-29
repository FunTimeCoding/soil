package mock_client

func (c *Client) Navigate(
	identifier int,
	l string,
) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if e := c.failures.Caller(); e != nil {
		return e
	}

	_, t, f := c.findTab(identifier)

	if f != nil {
		return f
	}

	t.Locator = l

	return nil
}
