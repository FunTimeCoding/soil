package mock_client

func (c *Client) Navigate(
	context string,
	locator string,
) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	v, e := c.find(context)

	if e != nil {
		return e
	}

	v.Locator = locator

	return nil
}
